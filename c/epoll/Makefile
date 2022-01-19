CC=clang
CFLAGS=-g -Wall -Werror
EXES=epoll_server epoll_server_et

all: $(EXES)

epoll_server_et: epoll_server_et.c
	$(CC) $(CFLAGS) -o $@ $^

epoll_server: epoll_server.c
	$(CC) $(CFLAGS) -o $@ $^

clean:
	rm -rf $(EXES) *.dSYM

