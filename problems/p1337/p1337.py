from typing import List
from functools import cmp_to_key
import unittest


class item:
    idx: int
    cnt: int

    def __init__(self, idx: int, cnt: int) -> None:
        self.idx = idx
        self.cnt = cnt


def compare(first: item, second: item) -> int:
    if first.cnt < second.cnt:
        return -1
    elif first.cnt > second.cnt:
        return 1
    elif first.idx > second.idx:
        return 1
    else:
        return -1
        

class Solution:
    def bSearch(self, l: List[int]) -> int:
        left, right = 0, len(l)
        while left < right:
            mid = int((left + right) / 2)
            if l[mid] == 0:
                right = mid
            else:
                left = mid + 1
        return left

    def kWeakestRows(self, mat: List[List[int]], k: int) -> List[int]:
        items = []
        for idx in range(0, len(mat)):
            items.append(item(idx, self.bSearch(mat[idx])))
        items.sort(key=cmp_to_key(compare))
        result = []
        for i in range(0, k):
            result.append(items[i].idx)
        return result


class Test(unittest.TestCase):
    def setUp(self) -> None:
        self.s = Solution()

    def test_example1(self) -> None:
        self.assertEqual(
            [2, 0, 3],
            self.s.kWeakestRows(
                [
                    [1, 1, 0, 0, 0],
                    [1, 1, 1, 1, 0],
                    [1, 0, 0, 0, 0],
                    [1, 1, 0, 0, 0],
                    [1, 1, 1, 1, 1],
                ],
                3,
            ),
        )

    def test_example2(self) -> None:
        self.assertEqual(
            [0, 2],
            self.s.kWeakestRows(
                [[1, 0, 0, 0], [1, 1, 1, 1], [1, 0, 0, 0], [1, 0, 0, 0]], 2
            ),
        )


if __name__ == "__main__":
    unittest.main("p1337")
