CC=clang
CFLAGS=-Wall
EXE=mmap_demo
OBJS=mmap_demo.o

$(EXE): $(OBJS)
	$(CC) $(CFLAGS) -o $@ $^

%.o : %.c
	$(CC) $(CFLAGS) -c $^

clean:
	rm $(EXE) *.o