local M = {}

--- O(n²)
---
---@param nums number[]
---@param target number
---@return number[]|nil, string|nil
function M.two_sum_skip(nums, target)
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

--- O(n)
---
---@param nums number[]
---@param target number
---@return number[]|nil, string|nil
function M.two_sum(nums, target)
	local hash_map = {}
	for i, n in ipairs(nums) do
		if hash_map[target - n] ~= nil then
			return { hash_map[target - n], i - 1 }, nil
		end
		hash_map[n] = i - 1
	end

	return nil, 'not found'
end

return M
