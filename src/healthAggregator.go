package health

// Health Aggregator — checks all services and determines system health.
//
// YOU MUST IMPLEMENT the methods marked with TODO.
// ServiceRegistry is working — use it to get service list.
//
// Author: Ravi Krishnan (SRE team)
// Last Modified: 2026-03-25

import (
	"time"
)

type CheckResult struct {
	ServiceName string
	Healthy     bool
	Message     string
	ResponseMs  int64
	Critical    bool
	Timestamp   time.Time
}

type SystemStatus string

const (
	StatusHealthy   SystemStatus = "healthy"
	StatusDegraded  SystemStatus = "degraded"
	StatusUnhealthy SystemStatus = "unhealthy"
)

type AggregatedHealth struct {
	Status       SystemStatus
	TotalChecks  int
	Healthy      int
	Unhealthy    int
	Results      []CheckResult
	CheckedAt    time.Time
}

type HealthAggregator struct {
	registry *ServiceRegistry
	history  []AggregatedHealth
}

func NewHealthAggregator(registry *ServiceRegistry) *HealthAggregator {
	return &HealthAggregator{
		registry: registry,
		history:  make([]AggregatedHealth, 0),
	}
}

// Check all registered services and return aggregated health.
// 1. Get all services from registry
// 2. For each service, call its CheckFunc and measure response time
// 3. Build CheckResult for each
// 4. Count healthy vs unhealthy
// 5. Determine overall status:
//    - "unhealthy" if any CRITICAL service is down
//    - "degraded" if any non-critical service is down
//    - "healthy" if all pass
// 6. Store result in history
// 7. Return AggregatedHealth
func (ha *HealthAggregator) CheckAll() AggregatedHealth {
	return AggregatedHealth{Status: StatusHealthy}
}

// Check a single service by name.
// 1. Look up service in registry
// 2. Call its CheckFunc, measure time
// 3. Return CheckResult
func (ha *HealthAggregator) CheckService(name string) (CheckResult, error) {
	return CheckResult{}, nil
}

// Get uptime percentage for a service from history.
// 1. Count how many times the service was healthy vs total checks
// 2. Return percentage (0-100)
func (ha *HealthAggregator) GetUptime(serviceName string) float64 {
	return 0
}

func (ha *HealthAggregator) GetHistory() []AggregatedHealth {
	return ha.history
}
