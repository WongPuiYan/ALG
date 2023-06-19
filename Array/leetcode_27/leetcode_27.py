from typing import List

"""
27. 移除元素
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
"""


class Solution1:
    def removeElement(self, nums: List[int], val: int) -> int:
        l, r = 0, len(nums) - 1
        while l <= r:
            if nums[l] == val and nums[r] != val:
                nums[l] = nums[l] ^ nums[r]
                nums[r] = nums[l] ^ nums[r]
                nums[l] = nums[l] ^ nums[r]
                l, r = l + 1, r - 1
            elif nums[r] == val:
                r -= 1
            else:
                l += 1
        return l


class Solution2:
    def removeElement(self, nums: List[int], val: int) -> int:
        slow = 0

        for fast in range(len(nums)):
            if nums[fast] != val:
                nums[slow] = nums[fast]
                slow += 1

        return slow
