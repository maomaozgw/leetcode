from typing import List
import unittest

from queue import Queue

offset = 100000
next_steps = [[0, 1], [1, 0], [0, -1], [-1, 0]]


class Solution:
    def updateMatrix(self, mat: List[List[int]]) -> List[List[int]]:
        l = len(mat)
        n = len(mat[0])
        q = Queue()
        result = [0]*l

        for i in range(0, l):
            result[i] = [0]*n
            for j in range(0, n):
                if mat[i][j] == 0:
                    q.put(i*offset+j)
                    result[i][j] = 0
        step = 0
        while q.qsize() > 0:
            ql = q.qsize()
            step += 1
            for i in range(0, ql):
                v = q.get()
                x = int(v/offset)
                y = int(v % offset)
                for j in range(0, 4):
                    nx, ny = x+next_steps[j][0], y+next_steps[j][1]
                    if nx < 0 or ny < 0 or nx >= l or ny >= n:
                        continue
                    if mat[nx][ny] == 0:
                        continue
                    result[nx][ny] = step
                    mat[nx][ny] = 0
                    q.put(nx*offset+ny)
        return result


class Test(unittest.TestCase):
    def setUp(self) -> None:
        self.s = Solution()

    def test_example1(self):
        self.assertEqual(
            [[0, 0, 0], [0, 1, 0], [0, 0, 0]],
            self.s.updateMatrix([[0, 0, 0], [0, 1, 0], [0, 0, 0]]))

    def test_example2(self):
        self.assertEqual(
            [[0, 0, 0], [0, 1, 0], [1, 2, 1]],
            self.s.updateMatrix([[0, 0, 0], [0, 1, 0], [1, 1, 1]]))


if __name__ == "__main__":
    unittest.main("p542")
