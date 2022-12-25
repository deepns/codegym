import threading
import random
import time

# TODO
# - [ ] split this into two examples

NUM_THREADS=1

def thread_func(tid):
    print("starting thread:{}, name:{}".format(tid, threading.get_ident()))
    
    # do some work.
    sleep_time = random.randint(1, 10)
    print("sleeping for {} seconds".format(sleep_time))
    time.sleep(sleep_time)

    print("done thread:", tid)

# Start the worker threads
threads = []
for i in range(NUM_THREADS):
    # starting a non-daemon thread
    # threading module has a shutdown function that runs on program exit
    # it calls join() on any non-daemon thread that is still active
    t = threading.Thread(target=thread_func, args=(i, ))
    threads.append(t)
    t.start()

print("main: done")

# not needed in case of non-daemon threads, if the intention is wait at the end.
# if need to wait for the threads to finish before main ends, then call join()
# where needed

# Wait until all threads are complete
# for thread in threads:
#     thread.join()
