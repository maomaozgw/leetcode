from typing import List


class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        num_len = len(nums)
        result = num_len+1
        left = 0
        win_sum = 0
        right = 0
        while right < num_len:
            if win_sum >= target:
                win_sum -= nums[left]
                left += 1
                result = min(result, right-left+1)
                continue
            win_sum += nums[right]
            right += 1
        while win_sum >= target:
            win_sum -= nums[left]
            left += 1
            result = min(result, num_len-left+1)
        if result > num_len:
            return 0
        return result
