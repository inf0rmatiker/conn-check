package main

import (
	"fmt"

	probing "github.com/prometheus-community/pro-bing"
)

func main() {
	pinger, err := probing.NewPinger("www.google.com")
	if err != nil {
		panic(err)
	}

	pinger.SetPrivileged(true) // Allow use of raw sockets (run as root)

	pinger.Count = 3
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
	fmt.Printf("stats.MaxRtt=%v\n", stats.MaxRtt)
}
