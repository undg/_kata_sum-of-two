local function two_sum(nums, target)
    return nil
end

describe("two_sum", function()
    it("finds two numbers that sum to target", function()
        assert.are.same({0, 1}, two_sum({1, 2, 3}, 3))
        assert.are.same({0, 2}, two_sum({1, 3, 2}, 3))
        assert.are.same({1, 3}, two_sum({30, 15, 10, 15, 30}, 30))
    end)
end)
