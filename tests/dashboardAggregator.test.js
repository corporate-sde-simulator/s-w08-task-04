const DashboardAggregator = require("../src/dashboardAggregator.js");
const StatusResolver = require("../src/statusResolver.js");

describe("Health check dashboard aggregator", () => {
    test("should process valid input", () => {
        const obj = new DashboardAggregator();
        expect(obj.process({ key: "val" })).not.toBeNull();
    });
    test("should handle null", () => {
        const obj = new DashboardAggregator();
        expect(obj.process(null)).toBeNull();
    });
    test("should track stats", () => {
        const obj = new DashboardAggregator();
        obj.process({ x: 1 });
        expect(obj.getStats().processed).toBe(1);
    });
    test("support should work", () => {
        const obj = new StatusResolver();
        expect(obj.process({ data: "test" })).not.toBeNull();
    });
});
