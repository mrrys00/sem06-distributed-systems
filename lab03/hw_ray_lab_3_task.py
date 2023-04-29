"""
3.0 start remote cluster settings and observe actors in cluster
a) make screenshot of dependencies
3.1. Modify the Actor class MethodStateCounter and add/modify methods that return the following:
a) - Get number of times an invoker name was called
b) - Get a list of values computed by invoker name
c) - Get state of all invokers
"""

import logging
from fractions import Fraction
from math import pi as PI
from random import choice, randint, random
from time import time

import ray


@ray.remote
class MethodStateCounter32:
    """
    Class 3.2
    """

    def __init__(self) -> None:
        self.invokers: dict[str:int] = {"A": 0, "B": 0, "C": 0}
        self.vals: dict[str:int] = {
            "A": [],
            "B": [],
            "C": []
        }

    def invoke(self, name: str) -> int:
        """return the state of that invoker"""
        sleep(0.5)
        self.invokers[name] += 1
        val = randint(5, 25)
        self.vals[name].append(val)
        return val

    def get_invoker_state(self, name: str) -> int:
        """return the state of the named invoker"""
        return self.invokers[name]

    def get_all_invoker_state(self) -> dict:
        """return the state of all invokers"""
        return self.invokers

    def get_all_invorek_vals(self) -> dict:
        """return all invoker vals"""
        return self.vals


def ex_3_2(callers_: list):
    """
    3.2 Modify method invoke to return a random int value between [5, 25]
    """
    worker_invoker = MethodStateCounter32.remote()
    print(worker_invoker)

    for _ in range(10):
        name = choice(callers_)
        worker_invoker.invoke.remote(name)

    print("method callers")
    for _ in range(5):
        random_name_invoker = choice(callers_)
        val = ray.get(worker_invoker.invoke.remote(random_name_invoker))
        print(f"caller: {random_name_invoker} returned {val}")

    print(
        f"State of all invoker: {ray.get(worker_invoker.get_all_invoker_state.remote())}")
    print(
        f"All return values: {ray.get(worker_invoker.get_all_invorek_vals.remote())}")


@ray.remote
class MethodStateCounter33:
    """
    Class 3.3
    """

    def __init__(self, samples_):
        self.samples = samples_
        self.invokers: dict[str:int] = {"A": 0, "B": 0, "C": 0}
        self.vals: dict[str:int] = {
            "A": [],
            "B": [],
            "C": []
        }

    def invoke(self, name: str) -> int:
        """
        Almost copy paste pi4_sample
        """
        in_count = 0
        for _ in range(self.samples):
            x, y = random(), random()
            if x * x + y * y <= 1:
                in_count += 1

        val = Fraction(in_count, self.samples)
        self.invokers[name] += 1
        self.vals[name].append(val)
        return val

    def get_invoker_state(self, name: str) -> int:
        """return the state of the named invoker"""
        return self.invokers[name]

    def get_all_invoker_state(self) -> dict:
        """return the state of all invokers"""
        return self.invokers

    def get_all_invorek_vals(self) -> dict:
        """return all invoker vals"""
        return self.vals


def ex_3_3(callers_: list, samples_: int):
    """
    3.3 Take a look on implement parralel Pi computation
    based on https://docs.ray.io/en/master/ray-core/examples/highly_parallel.html

    Implement calculating pi as a combination of actor (which keeps the
    state of the progress of calculating pi as it approaches its final value)
    and a task (which computes candidates for pi)
    """
    worker_invoker = MethodStateCounter33.remote(samples_)
    print(worker_invoker)

    start_time = time()

    cal_len = len(callers_)
    ran_ = 2**12
    for i in range(ran_):
        worker_invoker.invoke.remote(callers_[i % cal_len])

    vals: dict = ray.get(worker_invoker.get_all_invorek_vals.remote())

    pi_sum: float = 0.0
    pi_len: int = 0
    for val in vals.values():
        pi_sum += sum(val)
        pi_len += len(val)

    pi_val = pi_sum * 4 / pi_len

    print(
        f"Library π:\t{PI}\nCalculated π:\t{pi_val}\nInaccuracy: {abs(PI/pi_val-1)*100}%")
    print(f"Duration: \t{time() - start_time}s")


if __name__ == "__main__":
    callers, samples = ["A", "B", "C"], 2**18

    ray.init(num_cpus=2**3, ignore_reinit_error=True,
             logging_level=logging.ERROR)
    ex_3_2(callers)
    ex_3_3(callers, samples)
    ray.shutdown()
