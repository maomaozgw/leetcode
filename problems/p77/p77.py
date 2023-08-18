from typing import List
import unittest


class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        result = []

        def combineF(start: int, current: List[int]):
            if len(current) == k:
                tmp = current.copy()
                result.append(tmp)
            for i in range(start, n+1):
                current.append(i)
                combineF(i+1, current)
                current.pop()
        combineF(1, [])
        return result


class SolutionTest(unittest.TestCase):
    def setUp(self) -> None:
        self.s = Solution()

    def test_example1(self):
        self.assertEqual([[1, 2], [1, 3], [1, 4], [2, 3],
                         [2, 4], [3, 4]], self.s.combine(4, 2))


if __name__ == "__main__":
    unittest.main("p77")
