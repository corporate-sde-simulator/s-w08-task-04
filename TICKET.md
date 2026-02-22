# FINSERV-4258: Build service health check aggregator

**Status:** In Progress · **Priority:** High
**Sprint:** Sprint 30 · **Story Points:** 5
**Reporter:** Ravi Krishnan (SRE Lead) · **Assignee:** You (Intern)
**Due:** End of sprint (Friday)
**Labels:** `backend`, `golang`, `monitoring`, `health`
**Task Type:** Feature Ship

---

## Description

The `ServiceRegistry` tracks registered services. Build the `HealthAggregator` that periodically checks all registered services, aggregates their health status, and determines overall system health. Implement TODOs in `healthAggregator.go`.

## Acceptance Criteria

- [ ] `CheckAll()` runs health checks on all registered services
- [ ] Overall status is `healthy` only if all services pass
- [ ] Overall status is `degraded` if any non-critical service is down
- [ ] Overall status is `unhealthy` if any critical service is down
- [ ] Check results include response times
- [ ] All unit tests pass
