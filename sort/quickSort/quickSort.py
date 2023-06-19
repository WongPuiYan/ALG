from typing import List
import random


def quick_sort(nums: List[int]) -> List[int]:
    """
    快速排序
    """
    def qsort(left: int, right: int) -> None:
        if left >= right:
            return

        p_index = partition(left, right)
        if p_index > left + 1:
            while p_index > left + 1 and nums[p_index] == nums[p_index - 1]:
                p_index -= 1
            qsort(left, p_index - 1)
        if p_index < right - 1:
            while p_index < right - 1 and nums[p_index] == nums[p_index + 1]:
                p_index += 1
            qsort(p_index + 1, right)

    def partition(left: int, right: int) -> int:
        pos = random.randint(left, right)
        tar = nums[pos]
        nums[left], nums[pos] = nums[pos], nums[left]

        while left < right:
            while right > left and nums[right] >= tar:
                right -= 1
            nums[left] = nums[right]
            while left < right and nums[left] <= tar:
                left += 1
            nums[right] = nums[left]
        nums[left] = tar

        return left

    qsort(0, len(nums) - 1)

    return nums


if __name__ == "__main__":
    nums = [9, 8, 7, 6, 5, 4, 3, 2, 1]
    nums = [3, 2, 1, 8, 2, 1, 6, 7, 9]
    func = quick_sort
    res = func(nums)
    print(res)
