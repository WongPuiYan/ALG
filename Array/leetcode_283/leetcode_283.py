from typing import List

"""
283. 移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意 ，必须在不复制数组的情况下原地对数组进行操作。
"""


class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        slow = 0
        nums_len = len(nums)

        for fast in range(nums_len):
            if nums[fast] != 0:
                nums[slow] = nums[fast]
                slow += 1

        while slow < nums_len:
            nums[slow] = 0
            slow += 1
