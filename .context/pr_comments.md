# PR Review - Health check dashboard aggregator (by Deepak Gupta)

## Reviewer: Neha Sharma
---

**Overall:** Good foundation but critical bugs need fixing before merge.

### `dashboardAggregator.js`

> **Bug #1:** Service group status shows GREEN when any service is green instead of when all are green
> This is the higher priority fix. Check the logic carefully and compare against the design doc.

### `statusResolver.js`

> **Bug #2:** Last-check timestamp uses server time instead of actual check time and shows stale data as fresh
> This is more subtle but will cause issues in production. Make sure to add a test case for this.

---

**Deepak Gupta**
> Acknowledged. I have documented the issues for whoever picks this up.
