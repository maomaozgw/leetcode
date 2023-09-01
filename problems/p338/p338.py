from typing import List
import unittest


class Solution:
    def countBits(self, n: int) -> List[int]:
        result = []
        for i in range(0, n + 1):
            result.append(self.cntBits(i))
        return result

    def cntBits(self, n: int) -> int:
        c = 0
        while n != 0:
            n = n & (n - 1)
            c += 1
        return c


class Test(unittest.TestCase):
    def setUp(self) -> None:
        self.s = Solution()

    def test_example1(self) -> None:
        self.assertEqual([0, 1, 1], self.s.countBits(2))

    def test_example2(self) -> None:
        self.assertEqual([0, 1, 1, 2, 1, 2], self.s.countBits(5))


if __name__ == "__main__":
    unittest.main("p338")
