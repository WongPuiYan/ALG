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
