#! /usr/bin/env python3
import time
from multiprocessing import shared_memory

def run_client():
    # attach to the shared memory created in the server instance
    shm = shared_memory.SharedMemory(name="shmdemo")
    shm.buf[:6] = b'wakeup'
    shm.close()

if __name__ == "__main__":
    run_client()