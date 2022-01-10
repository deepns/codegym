#include <sys/uio.h> // for iovec
#include <stdio.h>
#include <stdlib.h>

// for open calls and flags
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>

#include <string.h> // for memset()

#include <unistd.h> // for close()

#define INFILE "vio_infile"
#define OUTFILE "vio_outfile"
#define BUFCOUNT 12
#define BUFLEN 32

/* 
 * Trying to see how to use vectored IO with readv and 
 * writev system calls.
 * 
 * we open a file, read a chunk of data into multiple buffers
 * using readv, and write the buffers into a new file.
 */

static int infile = -1;
static int outfile = -1;
struct iovec readbufs[BUFCOUNT];

/*
 * clean up the resources used in this file.
 */
static void cleanup()
{
    for (int i=0; i<BUFCOUNT; i++) {
        free(readbufs[i].iov_base);
    }

    close(infile);
    close(outfile);
}

/*
 * Return the size of the file pointed by the given descriptor
 */
off_t get_file_size(int fd)
{
    struct stat statbuf;
    if (fstat(fd, &statbuf) == -1) {
        perror("fstat failed");
        return -1;
    }
    return statbuf.st_size;
}

int main()
{
    infile = open(INFILE, O_RDONLY);
    if (infile == -1) {
        perror("infile open failed");
        exit(EXIT_FAILURE);
    }

    printf("sizeof infile=%ld\n", get_file_size(infile));
    
    for (int i=0; i<BUFCOUNT; i++) {
        readbufs[i].iov_base = malloc(sizeof(char) * BUFLEN);
        memset(readbufs[i].iov_base, 0, BUFLEN);
        readbufs[i].iov_len = BUFLEN;
    }

    ssize_t bytesread = readv(infile, readbufs, BUFCOUNT);
    printf("read %ld bytes from %s\n", bytesread, INFILE);

    for (int i=0; i<BUFCOUNT; i++) {
        printf("%s", readbufs[i].iov_base);
    }
    printf("\n");

    int flags = O_CREAT | O_WRONLY | O_TRUNC; // start from the beginning
    outfile = open(OUTFILE, flags);
    if (outfile == -1) {
        perror("outfile open failed");
        cleanup();
        exit(EXIT_FAILURE);
    }

    ssize_t byteswritten = writev(outfile, readbufs, BUFCOUNT);
    printf("wrote %ld bytes to %s\n", byteswritten, OUTFILE);

    cleanup();
    return 0;
}