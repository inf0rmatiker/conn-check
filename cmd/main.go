package main

import (
	"fmt"

	"github.com/inf0rmatiker/conncheck/pkg/pinger"
)

// Example main() making use of the pinger package.
func main() {
	p := pinger.NewPinger()
	fmt.Printf("Created Pinger with default count %d, interval %s\n", p.PingCount(), p.PingInterval().String())
}
