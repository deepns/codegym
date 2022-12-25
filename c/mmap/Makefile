CC=gcc
CFLAGS=-Wall
EXES=mmap_example.app mmap_parent_child.app

all: $(EXES)

%.app : %.c
	$(CC) $(CFLAGS) -o $@ $<

clean:
	rm -rf $(EXES)