CC=clang
CFLAGS=-g -Wall -Werror
EXES=epoll_server

all: $(EXES)

epoll_server: epoll_server.c
	$(CC) $(CFLAGS) -o $@ $^

clean:
	rm -rf $(EXES) *.dSYM

