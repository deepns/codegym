#include <stdio.h>
#include <unistd.h>
#include <sys/mman.h>
#include <string.h>

int main(void) {
    int page_size = getpagesize();
    char *test_string = "Hello, World";
    char *maddr = NULL;

    printf("System pagesize=%d\n", page_size);
    printf("My pid=%u, ppid=%u\n", getpid(), getppid());

    // Get a page from mmap to write to.
    maddr = mmap(NULL,
                 page_size,
                 PROT_READ | PROT_WRITE,
                 MAP_ANONYMOUS | MAP_PRIVATE,
                 -1,
                 0);
    
    if (maddr == MAP_FAILED) {
        perror("mmap failed to allocate");
        return 1;
    }
    printf("Mapping string=\"%s\" to the address=0x%p returned by mmap\n", test_string, maddr);

    // write some contents to the allocated address
    strncpy(maddr, test_string, strlen(test_string));

    // read from the address
    printf("Reading from allocated address\n");
    for (int i=0; i<strlen(test_string); i++) {
        printf("%c", maddr[i]);
    }
    printf("\n...Done\n");

    // unmap the allocated address
    if (munmap(maddr, page_size) != 0) {
        perror("munmap failed");
    }

    return 0;
}
