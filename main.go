package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/netspeed/internal/colorutil"
	"github.com/showwin/speedtest-go/speedtest"
	"github.com/showwin/speedtest-go/speedtest/transport"
)

func newSpinner(suffix string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = suffix
	err := s.Color("cyan")
	if err != nil {
		return nil
	}
	return s
}

func main() {
	fmt.Println()
	colorutil.Bold.Println("⏱  Testing internet speed...")
	fmt.Println()

	// Create speedtest client
	client := speedtest.New()

	// Find server
	sp := newSpinner("  Finding best server...")
	sp.Start()

	serverList, err := client.FetchServers()
	if err != nil {
		sp.Stop()
		colorutil.Red.Fprintf(os.Stderr, "Error fetching servers: %v\n", err)
		os.Exit(1)
	}

	targets, err := serverList.FindServer([]int{})
	if err != nil || len(targets) == 0 {
		sp.Stop()
		colorutil.Red.Fprintf(os.Stderr, "Error finding server: %v\n", err)
		os.Exit(1)
	}

	server := targets[0]
	sp.Stop()

	colorutil.Cyan.Printf("   Server:    ")
	colorutil.DimWhite.Printf("%s (%s) [%.2f km]\n", server.Name, server.Sponsor, server.Distance)

	// Ping test
	sp = newSpinner("  Testing ping...")
	sp.Start()

	err = server.PingTest(nil)
	if err != nil {
		sp.Stop()
		colorutil.Red.Fprintf(os.Stderr, "Error during ping test: %v\n", err)
		os.Exit(1)
	}
	sp.Stop()

	latencyMs := float64(server.Latency) / float64(time.Millisecond)
	jitterMs := float64(server.Jitter) / float64(time.Millisecond)

	colorutil.Cyan.Printf("   Ping:      ")
	colorutil.ByPing(latencyMs).Printf("%.2f ms\n", latencyMs)

	colorutil.Cyan.Printf("   Jitter:    ")
	colorutil.ByJitter(jitterMs).Printf("%.2f ms\n", jitterMs)

	// Download test
	sp = newSpinner("  Testing download speed...")
	sp.Start()

	err = server.DownloadTest()
	if err != nil {
		sp.Stop()
		colorutil.Red.Fprintf(os.Stderr, "Error during download test: %v\n", err)
		os.Exit(1)
	}
	sp.Stop()

	dlMbps := server.DLSpeed.Mbps()

	colorutil.Cyan.Printf("   Download:  ")
	colorutil.ByDownload(dlMbps).Printf("%.2f Mbps\n", dlMbps)

	// Upload test
	sp = newSpinner("  Testing upload speed...")
	sp.Start()

	err = server.UploadTest()
	if err != nil {
		sp.Stop()
		colorutil.Red.Fprintf(os.Stderr, "Error during upload test: %v\n", err)
		os.Exit(1)
	}
	sp.Stop()

	ulMbps := server.ULSpeed.Mbps()

	colorutil.Cyan.Printf("   Upload:    ")
	colorutil.ByUpload(ulMbps).Printf("%.2f Mbps\n", ulMbps)

	// Packet loss test
	sp = newSpinner("  Testing packet loss...")
	sp.Start()

	analyzer := speedtest.NewPacketLossAnalyzer(nil)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	pktLoss := -1.0
	err = analyzer.RunWithContext(ctx, server.Host, func(pl *transport.PLoss) {
		pktLoss = pl.LossPercent()
	})
	sp.Stop()

	colorutil.Cyan.Printf("   Pkt Loss:  ")
	if err != nil || pktLoss < 0 {
		colorutil.Yellow.Printf("N/A\n")
	} else {
		colorutil.ByPacketLoss(pktLoss).Printf("%.2f%%\n", pktLoss)
	}

	fmt.Println()
}
