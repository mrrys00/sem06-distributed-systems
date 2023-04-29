"""
Create large lists and python dictionaries, put them in object store.
Write a Ray task to process them.
"""

from random import randint
from ray import ObjectRef

import ray

def data_creator(num_of_sets: int, elements_per_set: int) -> \
    tuple[list[list[int]], list[dict[int:int]]]:
    """Data creator produces list and dictionary"""
    lists_: list[list[int]] = \
        [[randint(0, 2**64) for _ in range(elements_per_set)] for _ in range(num_of_sets)]

    dicts_: list[dict[int:int]] = []
    for _ in range(num_of_sets):
        tmp_dict = dict()
        for i in range(elements_per_set):
            tmp_dict[i] = randint(0, 2**64)

        dicts_.append(tmp_dict)

    return lists_, dicts_

@ray.remote
def data_processor(lists_: list[ObjectRef], dicts_: list[ObjectRef]) -> tuple[int, int]:
    """Data processor"""
    return sum(x for y in ray.get(lists_) for x in y), \
        sum(x for y in ray.get(dicts_) for x in y.values())

def resolver() -> None:
    """Simple resolver"""
    num_lists, num_dicts = data_creator(2**3, 2**8)

    num_list_id = [ray.put(i) for i in num_lists]
    num_dict_id = [ray.put(i) for i in num_dicts]

    result_id = data_processor.remote(num_list_id, num_dict_id)

    list_sum, dict_sum = ray.get(result_id)

    print(f"List sum: {list_sum}\nDict sum: {dict_sum}\n")

if __name__ == "__main__":
    ray.init(num_cpus=8)
    resolver()
    ray.shutdown()
