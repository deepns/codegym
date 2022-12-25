CC=gcc
CFLAGS=-Wall
EXE=mmap_with_file
OBJS=mmap_with_file.o

$(EXE): $(OBJS)
	$(CC) $(CFLAGS) -o $@ $^

%.o : %.c
	$(CC) $(CFLAGS) -c $^

clean:
	rm $(EXE) *.o
