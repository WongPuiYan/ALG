import random


def generate_random_array(max_size: int, max_value: int) -> list:
    arr = []
    list_size = int((max_size + 1) * random.random())
    for i in range(list_size):
        arr.append(int((max_value + 1) * random.random() - (max_value + 1) * random.random()))
    return arr


def copy_array(arr:list)->list:
    copy_arr = []
    for i in arr:
        copy_arr.append(i)
    return copy_arr


def is_equal(arr1:list, arr2:list)-> bool:
    if len(arr1) != len(arr2):
        return False
    for i in range(len(arr1)):
        if arr1[i] != arr2[i]:
            return False
    return True


if __name__ == "__main__":
    test_time = 1
    max_size = 10
    max_value = 10
    arr = generate_random_array(max_size, max_value)
    print(arr) 