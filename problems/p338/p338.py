from typing import List
import unittest


class Solution:
    def countBits(self, n: int) -> List[int]:
        result = [0] * (n + 1)
        for i in range(1, n + 1):
            result[i] = result[i >> 1] + (i & 1)
        return result


class Test(unittest.TestCase):
    def setUp(self) -> None:
        self.s = Solution()

    def test_example1(self) -> None:
        self.assertEqual([0, 1, 1], self.s.countBits(2))

    def test_example2(self) -> None:
        self.assertEqual([0, 1, 1, 2, 1, 2], self.s.countBits(5))


if __name__ == "__main__":
    unittest.main("p338")
