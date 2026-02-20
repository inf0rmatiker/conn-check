package pinger

import (
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

// Total report of endpoint ping results and their statistics
type PingReport struct {

	// Mapping of endpoints to ping result
	responsive map[string]probing.Statistics

	// Mapping of endpoints to ping result
	unresponsive map[string]probing.Statistics
}

// Default behavior for determining if an endpoint is 'responsive' or not.
// Very liberal criteria: if even a single packet was received, consider
// the endpoint responsive.
func defaultIsResponsive(result *probing.Statistics) bool {
	return result.PacketsRecv > 0
}

// Function that determines the criteria for whether an endpoint
// is responsive or not, based on probing.Statistics results.
// Return true if the endpoint should be considered responsive,
// false otherwise.
type CriteriaFunc func(*probing.Statistics) bool

type Pinger struct {
	// Slice of endpoints to ping (DNS names or IP addresses).
	endpoints []string

	// Function for determining if an endpoint is responsive or not.
	criteriaFunc func(*probing.Statistics) bool

	// Amount of times to ping each endpoint.
	pingCount int

	// Interval between pings for each endpoint.
	pingInterval time.Duration
}

// Constructs and returns a default Pinger instance.
func NewPinger() *Pinger {
	return &Pinger{
		endpoints:    []string{},
		criteriaFunc: defaultIsResponsive,
		pingCount:    1,
		pingInterval: time.Second,
	}
}

// Sets multiple endpoints to ping. Each endpoint could be
// an IPv4 address, IPv6 address, or DNS name.
func (p *Pinger) ForEndpoints(endpoints []string) *Pinger {
	p.endpoints = append(p.endpoints, endpoints...)
	return p
}

// Sets how many times to try to ping each endpoint.
func (p *Pinger) WithPingCount(count int) *Pinger {
	p.pingCount = count
	return p
}

// Sets how many times to try to ping each endpoint.
func (p *Pinger) WithPingInterval(t time.Duration) *Pinger {
	p.pingInterval = t
	return p
}

// Sets the criteria function to determine an endpoint's responsivity.
func (p *Pinger) WithResponsivityCriteria(critFunc CriteriaFunc) *Pinger {
	p.criteriaFunc = critFunc
	return p
}

func (p Pinger) PingCount() int {
	return p.pingCount
}

func (p Pinger) PingInterval() time.Duration {
	return p.pingInterval
}

func (p Pinger) Endpoints() []string {
	return p.endpoints
}
