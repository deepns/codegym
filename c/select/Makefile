CC=clang
CFLAGS=-g -Wall
EXES=select_server select_client

all: $(EXES)

select_server: select_server.c
	$(CC) $(CFLAGS) -o $@ $^

select_client: select_client.c
	$(CC) $(CFLAGS) -o $@ $^

clean:
	rm -rf $(EXES) *.dSYM

