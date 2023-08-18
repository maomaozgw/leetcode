from typing import List
import unittest


class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        num_len = len(nums)
        result = num_len+1
        left = 0
        win_sum = 0
        right = 0
        while right < num_len:
            if win_sum >= target:
                win_sum -= nums[left]
                left += 1
                result = min(result, right-left+1)
                continue
            win_sum += nums[right]
            right += 1
        while win_sum >= target:
            win_sum -= nums[left]
            left += 1
            result = min(result, num_len-left+1)
        if result > num_len:
            return 0
        return result


class Test(unittest.TestCase):
    def setUp(self) -> None:
        self.s = Solution()

    def test_example_1(self):
        self.assertEqual(self.s.minSubArrayLen(7, [2, 3, 1, 2, 4, 3]), 2)

    def test_example_2(self):
        self.assertEqual(self.s.minSubArrayLen(4, [1, 4, 4]), 1)

    def test_example_3(self):
        self.assertEqual(self.s.minSubArrayLen(
            11, [1, 1, 1, 1, 1, 1, 1, 1]), 0)


if __name__ == "__main__":
    unittest.main("p209")
