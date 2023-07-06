from problems.p209.p209 import Solution

import unittest


class Test_TestP209(unittest.TestCase):
    def test_example_1(self):
        s = Solution()
        self.assertEqual(s.minSubArrayLen(7, [2, 3, 1, 2, 4, 3]), 2)

    def test_example_2(self):
        s = Solution()
        self.assertEqual(s.minSubArrayLen(4, [1, 4, 4]), 1)

    def test_example_3(self):
        s = Solution()
        self.assertEqual(s.minSubArrayLen(11, [1, 1, 1, 1, 1, 1, 1, 1]), 0)
