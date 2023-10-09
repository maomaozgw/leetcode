from typing import List, Callable
import unittest


class Solution:
    def binary_search(
        self,
        left: int,
        right: int,
        f: Callable[[int], bool],
    ):
        while left < right:
            mid = (left + right) // 2
            if f(mid):
                right = mid
            else:
                left = mid + 1
        return left

    def searchRange(self, nums: List[int], target: int) -> List[int]:
        first_idx, last_idx = -1, -1
        if len(nums) == 0 or nums[0] > target or nums[-1] < target:
            return [first_idx, last_idx]
        last_idx = self.binary_search(
            0, len(nums)-1,
            lambda x: nums[x] >= target and (x == len(nums) - 1 or nums[x+1] > target)
        )
        if nums[last_idx] != target:
            return [-1, -1]

        first_idx = self.binary_search(
            0, last_idx,
            lambda x: nums[x] >= target and (x == 0 or nums[x-1] <= target)
        )
        return [first_idx, last_idx]


class Test(unittest.TestCase):
    def setUp(self) -> None:
        self.s = Solution()

    def test_example1(self) -> None:
        self.assertEqual([3, 4], self.s.searchRange([5, 7, 7, 8, 8, 10], 8))

    def test_example2(self) -> None:
        self.assertEqual([-1, -1], self.s.searchRange([5, 7, 7, 8, 8, 10], 6))

    def test_example3(self) -> None:
        self.assertEqual([-1, -1], self.s.searchRange([], 6))
    
    def test_example4(self) -> None:
        self.assertEqual([0, 0], self.s.searchRange([1], 1))

    def test_example5(self) -> None:
        self.assertEqual([1, 1], self.s.searchRange([1, 4], 4))

    def test_example6(self) -> None:
        self.assertEqual([0, 1], self.s.searchRange([2, 2], 2))

    def test_example7(self) -> None:
        self.assertEqual([0, 0], self.s.searchRange([1, 4], 1))


if __name__ == "__main__":
    unittest.main("p34")
