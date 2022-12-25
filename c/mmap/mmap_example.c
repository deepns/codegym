#include <stdio.h>
#include <stdlib.h>
#include <sys/mman.h>

int main(void) {
    void *shared_mem = NULL;
    int data = 10;

    shared_mem = mmap(NULL, /* addr */
                        sizeof(int), /* size of mapped region */
                        PROT_READ | PROT_WRITE, /* protection */
                        MAP_ANONYMOUS | MAP_SHARED, /* flags specifying type of mapped object */
                        -1, /* fd */
                        0 /* offset */);

    if (shared_mem == MAP_FAILED) {
        perror("mmap failed");
        return EXIT_FAILURE;
    }

    printf("Writing data=%d at addr=%p\n", data, shared_mem);
    *(int *)shared_mem = 10;
    printf("Reading from addr=%p returned data=%d\n", shared_mem, *(int *)shared_mem);
    return 0;
}