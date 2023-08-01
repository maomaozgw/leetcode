from typing import List


class Solution:
    def findNumberOfLIS(self, nums: List[int]) -> int:
        largest = 0
        dp = [1]*len(nums)
        cnt_map = {}
        for i in range(0, len(nums)):
            cnt_map[i] = 1
        for i in range(len(nums)):
            for j in range(i):
                if nums[i] > nums[j] and dp[i] < dp[j]+1:
                    dp[i] = dp[j]+1
                    cnt_map[i] = cnt_map[j]
                elif nums[i] > nums[j] and dp[i] == dp[j]+1:
                    cnt_map[i] += cnt_map[j]
            largest = max(largest, dp[i])
        result = 0
        for i in range(len(nums)):
            if dp[i] == largest:
                result += cnt_map[i]

        return result
