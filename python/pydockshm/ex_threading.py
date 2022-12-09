import threading
import random
import time

NUM_THREADS=10

def thread_func(tid):
    print("starting thread:", tid)
    # do some work.
    time.sleep(random.randint(1, 10))
    print("done thread:", tid)

# Start the worker threads
threads = []
for i in range(NUM_THREADS):
    t = threading.Thread(target=thread_func, args=(i, ))
    threads.append(t)
    t.start()

# Wait until all threads are complete
for thread in threads:
    thread.join()
