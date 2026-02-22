package health

// Service Registry — tracks registered services and their metadata.
//
// This module is COMPLETE. Your task is in healthAggregator.go.
//
// Author: Ravi Krishnan (SRE team)
// Last Modified: 2026-03-25

type ServiceInfo struct {
	Name       string
	Endpoint   string
	Critical   bool
	Team       string
	CheckFunc  func() (bool, string) // returns (healthy, message)
}

type ServiceRegistry struct {
	services map[string]ServiceInfo
}

func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{
		services: make(map[string]ServiceInfo),
	}
}

func (sr *ServiceRegistry) Register(info ServiceInfo) {
	sr.services[info.Name] = info
}

func (sr *ServiceRegistry) Deregister(name string) {
	delete(sr.services, name)
}

func (sr *ServiceRegistry) GetService(name string) (ServiceInfo, bool) {
	s, ok := sr.services[name]
	return s, ok
}

func (sr *ServiceRegistry) GetAllServices() []ServiceInfo {
	result := make([]ServiceInfo, 0, len(sr.services))
	for _, s := range sr.services {
		result = append(result, s)
	}
	return result
}

func (sr *ServiceRegistry) GetCriticalServices() []ServiceInfo {
	result := make([]ServiceInfo, 0)
	for _, s := range sr.services {
		if s.Critical {
			result = append(result, s)
		}
	}
	return result
}

func (sr *ServiceRegistry) Count() int {
	return len(sr.services)
}
