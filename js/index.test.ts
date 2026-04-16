import { twoSum } from '.'

describe(`twoSum()`, () => {
    it(`should return indices of target sum`, () => {
        expect(twoSum([1, 2, 3], 3)).toEqual([0, 1])
        expect(twoSum([1, 3, 2], 3)).toEqual([0, 2])
        expect(twoSum([30, 15, 10, 15, 30], 30)).toEqual([1, 3])
    })
})
