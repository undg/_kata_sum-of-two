// O(n(n-1)/2) nearly O(n*n)
export function twoSum_skip(nums: number[], target: number) {
    for (const [i, n] of nums.entries()) {
        const numsRest = nums.slice(i + 1)
        for (const [ii, nn] of numsRest.entries()) {
            const i2 = ii + i + 1
            const t = n + nn
            //console.log({ i, i2, target, t, numsRest })
            if (target === n + nn) {
                return [i, i2]
            }
        }
    }

    return 'not found'
}

// o(n)
export function twoSum(nums: number[], target: number) {
    const map = new Map<number, number>()

    for (const [i, n] of nums.entries()) {
        const minusTarget = target - n
        //console.log({ i, n, target, minusTarget, map })
        if (map.has(n)) {
            return [map.get(n), i]
        } else {
            map.set(minusTarget, i)
        }
    }

    return 'not found'
}
