"""
Try using local bubble sort and remote bubble sort, show difference
"""

from random import randint
from time import time

import ray

def random_list_generator(elem_num: int, start: int = 0, stop: int = 2**64) -> list:
    """Returns the random integers list"""
    return [randint(start, stop) for _ in range(elem_num)]

def bubble_sort(unsorted_list: list) -> list:
    """Simple bubble sort algorithm"""
    result_list = unsorted_list.copy()
    len_ = len(result_list)

    for i in range(len_):
        for j in range(0, len_-i-1):
            if result_list[j] > result_list[j+1]:
                result_list[j], result_list[j+1] = result_list[j+1], result_list[j]

    return result_list

@ray.remote
def compare_and_swap(list_: list, idx1: int, idx2: int) -> list:
    """Remotely delegated compare and swap function"""
    if list_[idx1] > list_[idx2]:
        list_[idx1], list_[idx2] = list_[idx2], list_[idx1]
    return list_

@ray.remote
def remote_bubble_sort(unsorted_list: list) -> list:
    """Bubble sort with remotely delegated compare and swam function"""
    result_list = unsorted_list.copy()
    len_ = len(result_list)

    for i in range(len_):
        for j in range(0, len_-i-1):
            result_list = ray.get(compare_and_swap.remote(result_list, j, j+1))

    return result_list

def comparator() -> None:
    """Solution test"""
    unsorted_list_1 = random_list_generator(2**7)
    unsorted_list_2 = unsorted_list_1.copy()

    bubble_start = time()
    sorted_list_1 = bubble_sort(unsorted_list_1)
    print(f"Bubble sort local\nExecution time: {time()-bubble_start}\n\
        Sort passed: {sorted_list_1 == sorted(unsorted_list_1)}\n")

    remote_start = time()
    sorted_list_2 = ray.get(remote_bubble_sort.remote(unsorted_list_2))
    print(f"Bubble sort remote\nExecution time: {time()-remote_start}\n\
        Sort passed: {sorted_list_2 == sorted(unsorted_list_2)}\n")

if __name__ == "__main__":
    ray.init()
    comparator()
    ray.shutdown()

"""
Results:
Data size 2^7:
2023-04-06 20:35:54,225 INFO worker.py:1553 -- Started a local Ray instance.
Bubble sort local
Execution time: 0.0009701251983642578
        Sort passed: True

Bubble sort remote
Execution time: 10.256896734237671
        Sort passed: True
"""
