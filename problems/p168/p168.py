import unittest

start = ord('A')


class Solution:
    def convertToTitle(self, columnNumber: int) -> str:
        result = ''
        while columnNumber > 0:
            idx = int((columnNumber - 1) % 26)
            result = chr(start + idx) + result
            columnNumber = int((columnNumber-idx)/26)
        return result


class Test(unittest.TestCase):
    def setUp(self) -> None:
        self.s = Solution()

    def test_example1(self) -> None:
        self.assertEqual('A', self.s.convertToTitle(1))

    def test_example2(self) -> None:
        self.assertEqual('AB', self.s.convertToTitle(28))

    def test_example3(self) -> None:
        self.assertEqual("ZY", self.s.convertToTitle(701))


if __name__ == "__main__":
    unittest.main("p168")
