#! /usr/bin/env python3

import time
from multiprocessing import shared_memory

# To share the same namespace if the server and client
# are in different containers

# start first container.
# docker run --rm --ipc=shareable --name shmserver --volume $PWD:/home python:latest python /home/shm_server.py

# start second container
# docker run --rm --ipc=container:shmserver --name shmclient --volume $PWD:/home python:latest python /home/shm_client.py

def run_server():
    # raises FileExistsError if sharedmemory segment with the given name
    # exists already
    shm = shared_memory.SharedMemory(name="shmdemo", create=True, size=256)

    # sleep until client says wake up
    while shm.buf[:6] != b'wakeup':
        print("sleeping...")
        time.sleep(5)

    print("Received wake up call from client")
    shm.close()
    shm.unlink()

if __name__ == "__main__":
    run_server()
    