

def binary_search(nums: list, target: int) -> int:
    left, right = 0, len(nums)
    while left < right:
        middle = left + (right - left) >> 1
        if nums[middle] < target:
            left = middle + 1
        elif nums[middle] > target:
            right = middle
        else:
            return middle
    return

def main():
    test_list = [1,2,3,4,5,6,7,8]
    target = 3
    print(test_list, target)
    print(binary_search(test_list, target))

if __name__ == '__main__':
    main()
