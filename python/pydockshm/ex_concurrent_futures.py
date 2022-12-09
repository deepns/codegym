import concurrent.futures
import random
import time

MAX_THREADS=3
NUM_JOBS=10

def thread_func(tid):
    print("starting thread:", tid)
    time.sleep(random.randint(1, 10))
    print("done thread:", tid)

with concurrent.futures.ThreadPoolExecutor(max_workers=MAX_THREADS) as executor:
    executor.map(thread_func, range(NUM_JOBS))
