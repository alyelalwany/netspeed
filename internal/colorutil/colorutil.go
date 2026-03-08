package colorutil

import "github.com/fatih/color"

// Pre-configured color styles for consistent terminal output.
var (
	Cyan     = color.New(color.FgCyan, color.Bold)
	Green    = color.New(color.FgGreen)
	Yellow   = color.New(color.FgYellow)
	Red      = color.New(color.FgRed)
	Bold     = color.New(color.Bold)
	DimWhite = color.New(color.FgWhite)
)

// ByDownload returns a color based on download speed thresholds.
// Green: >50 Mbps, Yellow: >10 Mbps, Red: <=10 Mbps.
func ByDownload(mbps float64) *color.Color {
	switch {
	case mbps > 50:
		return Green
	case mbps > 10:
		return Yellow
	default:
		return Red
	}
}

// ByUpload returns a color based on upload speed thresholds.
// Green: >25 Mbps, Yellow: >5 Mbps, Red: <=5 Mbps.
func ByUpload(mbps float64) *color.Color {
	switch {
	case mbps > 25:
		return Green
	case mbps > 5:
		return Yellow
	default:
		return Red
	}
}

// ByPing returns a color based on ping latency thresholds.
// Green: <30 ms, Yellow: <100 ms, Red: >=100 ms.
func ByPing(ms float64) *color.Color {
	switch {
	case ms < 30:
		return Green
	case ms < 100:
		return Yellow
	default:
		return Red
	}
}

// ByJitter returns a color based on jitter thresholds.
// Green: <5 ms, Yellow: <20 ms, Red: >=20 ms.
func ByJitter(ms float64) *color.Color {
	switch {
	case ms < 5:
		return Green
	case ms < 20:
		return Yellow
	default:
		return Red
	}
}

// ByPacketLoss returns a color based on packet loss thresholds.
// Green: <1%, Yellow: <5%, Red: >=5%.
func ByPacketLoss(pct float64) *color.Color {
	switch {
	case pct < 1:
		return Green
	case pct < 5:
		return Yellow
	default:
		return Red
	}
}
