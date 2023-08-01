from typing import List


class Solution:
    def asteroidCollision(self, asteroids: List[int]) -> List[int]:
        if len(asteroids) == 0:
            return []
        s = [asteroids[0]]
        last = asteroids[0]
        for i in range(1, len(asteroids)):
            if last < 0:
                s.append(asteroids[i])
                last = asteroids[i]
                continue
            if asteroids[i] > 0:
                s.append(asteroids[i])
                last = asteroids[i]
                continue
            need_apped = True
            while len(s) > 0:
                top = s[len(s)-1]
                if top < 0:
                    break
                if top > -asteroids[i]:
                    need_apped = False
                    break
                if top == -asteroids[i]:
                    need_apped = False
                    s.pop()
                    break
                s.pop()
            if need_apped:
                s.append(asteroids[i])
                last = asteroids[i]
        return s
