# netspeed

A CLI tool to test your internet connection speed from the terminal. Measures download speed, upload speed, ping, jitter, and packet loss with color-coded output.

## Requirements

- Go 1.21 or later

## Installation

```bash
go install github.com/netspeed@latest
```

Or build from source:

```bash
git clone https://github.com/<your-username>/netspeed.git
cd netspeed
go build -o netspeed .
```

## Usage

```bash
netspeed
```

The tool automatically finds the nearest Speedtest.net server and runs all tests. Example output:

```
⏱  Testing internet speed...

   Server:    Tokyo (OPEN Project) [12.34 km]
   Ping:      8.23 ms
   Jitter:    1.45 ms
   Download:  245.67 Mbps
   Upload:    98.32 Mbps
   Pkt Loss:  0.00%
```

Results are color-coded based on quality:

| Metric | Green | Yellow | Red |
|--------|-------|--------|-----|
| Download | > 50 Mbps | > 10 Mbps | <= 10 Mbps |
| Upload | > 25 Mbps | > 5 Mbps | <= 5 Mbps |
| Ping | < 30 ms | < 100 ms | >= 100 ms |
| Jitter | < 5 ms | < 20 ms | >= 20 ms |
| Packet Loss | < 1% | < 5% | >= 5% |

## License

MIT
