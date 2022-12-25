#include <stdio.h>
#include <sys/mman.h>
#include <unistd.h>
#include <fcntl.h>

/*
 * map contents of a file into memory using mmap
 */

#define PROTECTION (PROT_READ)
#define MAP_FLAGS (MAP_PRIVATE | MAP_FILE)
#define MMAP_SIZE 4096
#define BUF_SIZE 256

#define FILEPATH "testfile"

int main(int argc, char *argv[]) {

    int fd;
    int read_length = BUF_SIZE; // read 64 bytes at a time
    void *maddr;

    printf("Mapping %s to memory\n", FILEPATH);
    fd = open(FILEPATH, O_RDONLY);
    if (fd == -1) {
        perror("Unable to open the file");
    }

    // mmap the file just opened, starting at offset 0
    maddr = mmap(NULL, MMAP_SIZE, PROTECTION, MAP_FLAGS, fd, 0);
    if (maddr == MAP_FAILED) {
        perror("Unable to complete mmap");
        return 1;
    }

    printf("Reading from the memory...\n");
    for (int i=0; i < read_length; i++) {
        printf("%c", ((char *)maddr)[i]);
    }
    printf("\n");

    printf("Unmap the memory...\n");
    if (munmap(maddr, MMAP_SIZE) != 0) {
        perror("Unable to complete mumap");
        return 1;
    }

    close(fd);
    return 0;
}
