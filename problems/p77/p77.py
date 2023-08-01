from typing import List


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


if __name__ == "__main__":
    s = Solution()
    print(s.combine(4, 2))
