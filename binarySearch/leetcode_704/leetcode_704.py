"""
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/binary-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
"""

from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        l, r = 0, len(nums)
        while l < r:
            m = l + (r - l) // 2
            if nums[m] < target:
                l = m + 1
            elif nums[m] > target:
                r = m
            else:
                return m
        return -1


if __name__ == "__main__":
    s = Solution()
    nums = [-1, 0, 3, 5, 9, 12]
    target = 9
    print(s.search(nums, target))
