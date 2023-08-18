from typing import List
import unittest


class Solution:
    def maximalNetworkRank(self, n: int, roads: List[List[int]]) -> int:
        road_map = {}
        for i in range(0, n):
            road_map[i] = set()
        for pair in roads:
            road_map[pair[0]].add(pair[1])
            road_map[pair[1]].add(pair[0])
        result = 0
        for start in range(0, n):
            for end in range(start, n):
                if start == end:
                    continue
                rank = len(road_map[start]) + len(road_map[end])
                if end in road_map[start]:
                    rank -= 1
                result = max(result, rank)
        return result


class SolutionTest(unittest.TestCase):

    def setUp(self) -> None:
        self.s = Solution()

    def test_example1(self):
        self.assertEqual(4, self.s.maximalNetworkRank(
            4, [[0, 1], [0, 3], [1, 2], [1, 3]]))

    def test_example2(self):
        self.assertEqual(5, self.s.maximalNetworkRank(
            5, [[0, 1], [0, 3], [1, 2], [1, 3], [2, 3], [2, 4]]))

    def test_example3(self):
        self.assertEqual(5, self.s.maximalNetworkRank(
            8, [[0, 1], [1, 2], [2, 3], [2, 4], [5, 6], [5, 7]]))


if __name__ == '__main__':
    unittest.main()
