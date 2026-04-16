-- Placeholder implementation for two_sum
function two_sum(nums, target)
    return nil
end

local function test_two_sum()
    local cases = {
        {nums = {1, 2, 3}, target = 3, expected = {0, 1}},
        {nums = {1, 3, 2}, target = 3, expected = {0, 2}},
        {nums = {30, 15, 10, 15, 30}, target = 30, expected = {1, 3}},
    }

    for _, case in ipairs(cases) do
        local result = two_sum(case.nums, case.target)
        if not result or #result ~= #case.expected then
            error(string.format("test failed: expected %s, got %s", tostring(case.expected), tostring(result)))
        end
        for i = 1, #result do
            if result[i] ~= case.expected[i] then
                error(string.format("test failed at index %d: expected %d, got %d", i, case.expected[i], result[i]))
            end
        end
    end
    print("All tests passed!")
end

test_two_sum()
