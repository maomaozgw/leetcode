from typing import List


class Solution:
    def peakIndexInMountainArray(self, arr: List[int]) -> int:
        left = 0
        right = len(arr)
        while left < right:
            mid = int((left + right)/2)
            if mid == left:
                return left
            if arr[mid+1] > arr[mid] > arr[mid-1]: 
                left = mid
            elif arr[mid+1] < arr[mid] < arr[mid-1]:
                right = mid
            else:
                return mid
        return left 