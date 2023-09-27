import unittest


class Solution:
    def decodeAtIndex(self, s: str, k: int) -> str:
        size = 0
        for i in range(len(s)):
            if s[i].isdigit():
                size *= int(s[i])
            else:
                size += 1
        for i in range(len(s) - 1, -1, -1):
            k = k % size
            if k == 0 and s[i].isalpha():
                return s[i]
            if s[i].isdigit():
                size = int(size/int(s[i]))
            else:
                size -= 1
        return ""

    def decodeAtIndexOOM(self, s: str, k: int) -> str:
        before = ""
        k -= 1
        for i in range(len(s)):
            if "0" > s[i] or s[i] > "9":
                before += s[i]
                continue
            c = int(s[i])
            if c * len(before) > k:
                print(before)
                return before[k % len(before)]
            before = before * int(s[i])
        if len(before) > k:
            return before[k]
        return ""


class Test(unittest.TestCase):
    def setUp(self) -> None:
        self.s = Solution()

    def test_example1(self) -> None:
        self.assertEqual("o", self.s.decodeAtIndex("leet2code3", 10))

    def test_example2(self) -> None:
        self.assertEqual("h", self.s.decodeAtIndex("ha22", 5))

    def test_example3(self) -> None:
        self.assertEqual("a", self.s.decodeAtIndex("a2345678999999999999999", 1))

    def test_example4(self) -> None:
        self.assertEqual("a", self.s.decodeAtIndex("abc", 1))

    def test_example5(self) -> None:
        self.assertEqual("b", self.s.decodeAtIndex("a2b3c4d5e6f7g8h9", 3))


if __name__ == "__main__":
    unittest.main("p880", verbosity=True)
