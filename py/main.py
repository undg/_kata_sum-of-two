# O(n*n) ish
def two_sum_enum_range(nums: list[int], target: int) -> list[int] | str:
    for i, n in enumerate(nums):
        for ii in range(i + 1, len(nums)):
            nn = nums[ii]
            if target == n + nn:
                return [i, ii]
    return "not found"


# O(n*n) ish
def two_sum_enum_enum(nums: list[int], target: int) -> list[int] | str:
    for i, n in enumerate(nums):
        for ii, nn in enumerate(nums[i + 1 :], i + 1):
            if target == n + nn:
                return [i, ii]
    return "not found"


# O(n)
def two_sum(nums: list[int], target: int) -> list[int] | str:
    hash_rest: dict[int, int] = {}
    for i, n in enumerate(nums):
        if n in hash_rest:
            return [hash_rest[n], i]
        hash_rest[target - n] = i

    return "not found"
