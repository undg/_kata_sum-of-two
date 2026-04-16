import unittest

from main import two_sum


class TestTwoSum(unittest.TestCase):
    def test_two_sum_simple(self):
        self.assertEqual(two_sum([1, 2, 3], 3), [0, 1])

    def test_two_sum_reorder(self):
        self.assertEqual(two_sum([1, 3, 2], 3), [0, 2])

    def test_two_sum_larger_numbers(self):
        self.assertEqual(two_sum([30, 15, 10, 15, 30], 30), [1, 3])

    def test_two_sum_larger_not_found(self):
        self.assertEqual(two_sum([30, 14, 10, 15, 30], 30), "not found")


if __name__ == "__main__":
    unittest.main()
