local M = {}

-- O(n²)
---
---@param nums number[]
---@param target number
---@return number[]|nil, string|nil
function M.two_sum(nums, target)
	for i, n in ipairs(nums) do
		for ii = i + 1, #nums do
			local nn = nums[ii]
			if target == n + nn then
				return { i - 1, ii - 1 }, nil
			end
		end
	end

	return nil, 'not found'
end

return M
