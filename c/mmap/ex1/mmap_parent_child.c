#include <stdio.h>
#include <sys/wait.h> // for waitpid
#include <sys/mman.h> // for mmap
#include <unistd.h> // for fork

/*
 * An example of communicating between parent and child processes through
 * shared memory using mmap
 */

int main(void) {
    void *shared_mem = NULL;
    int data = 10;
    pid_t child_pid = -1;
    int child_state = -1;

    shared_mem = mmap(NULL,
                      sizeof(int),
                      PROT_READ | PROT_WRITE,
                      MAP_ANONYMOUS | MAP_PRIVATE,
                      -1 /* fd */,
                      0 /* offset */);
                                            
    if (shared_mem == MAP_FAILED) {
        perror("mmap failed");
        return -1;
    }

    child_pid = fork();
   
    if (child_pid == 0) {
        // write to shared memory from the child
        printf("(child) Writing data=%d at addr=%p\n", data, shared_mem);
        *(int *)shared_mem = 10;
    } else {
        // wait for child to complete before reading
        waitpid(child_pid, &child_state /*stat_loc*/, 0 /* options */ );
        printf("(parent) Reading from addr=%p returned data=%d\n", shared_mem, *(int *)shared_mem);
    }
    return 0;
}