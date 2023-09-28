from typing import List
import unittest


class Solution:
    def sortArrayByParity(self, nums: List[int]) -> List[int]:
        left, right = 0, len(nums) - 1
        while left < right:
            if nums[left] % 2 == 0:
                left += 1
                continue
            if nums[right] % 2 == 1:
                right -= 1
                continue
            nums[left], nums[right] = nums[right], nums[left]
            left += 1
            right -= 1
        return nums


class Test(unittest.TestCase):
    def setUp(self) -> None:
        self.s = Solution()

    def test_example1(self) -> None:
        self.assertEqual([4, 2, 1, 3], self.s.sortArrayByParity([3, 1, 2, 4]))

    def test_exmaple2(self) -> None:
        self.assertEqual([0], self.s.sortArrayByParity([0]))


if __name__ == "__main__":
    unittest.main("p905")
