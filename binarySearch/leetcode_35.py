"""
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/search-insert-position
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
"""

from typing import List


class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        left, right = 0, len(nums)
        while left < right:
            middle = left + ((right - left) >> 1)
            if nums[middle] < target:
                left = middle + 1
            elif nums[middle] > target:
                right = middle
            else:
                return middle

        return right


def main():
    nums = [1, 3, 5, 6, 7, 8, 9]
    target = 0
    target = 2
    s = Solution()
    print(s.searchInsert(nums, target))


if __name__ == '__main__':
    main()
