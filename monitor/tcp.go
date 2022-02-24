package monitor

import (
	"strconv"

	"github.com/chenjiandongx/pinger"
)

func MonitorTcpPing(path string) (checkSuccess bool, elapsed int) {
	opts := pinger.DefaultTCPPingOpts
	opts.PingCount = 1
	// TCP
	stats, _ := pinger.TCPPing(opts, path)
	pktSent := stats[0].PktSent
	pktSentStr := strconv.FormatInt(int64(pktSent), 10)
	pktSentFloat64, _ := strconv.ParseFloat(pktSentStr, 64)
	pktLosttRate := stats[0].PktLossRate

	timePing := stats[0].Mean.Milliseconds()
	timePingStr := strconv.FormatInt(timePing, 10)
	timePingInt, _ := strconv.Atoi(timePingStr)
	result := false
	if pktSentFloat64 > pktLosttRate {
		result = true
	}
	return result, timePingInt
}
