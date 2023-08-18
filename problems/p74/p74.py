from typing import List
import unittest


class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        m, n = len(matrix), len(matrix[0])
        left, right = 0, m*n
        while left < right:
            mid = int((left+right)/2)
            val = matrix[int(mid/n)][mid % n]
            if val == target:
                return True
            elif val > target:
                right = mid
            else:
                left = mid + 1

        return False


class Test(unittest.TestCase):
    def setUp(self) -> None:
        self.s = Solution()

    def test_example1(self) -> None:
        self.assertEqual(
            True,
            self.s.searchMatrix(
                [[1, 3, 5, 7], [10, 11, 16, 20], [23, 30, 34, 60]], 3
            )
        )

    def test_example2(self) -> None:
        self.assertEqual(
            False,
            self.s.searchMatrix(
                [[1, 3, 5, 7], [10, 11, 16, 20], [23, 30, 34, 60]], 13
            )
        )


if __name__ == "__main__":
    unittest.main("p74")
