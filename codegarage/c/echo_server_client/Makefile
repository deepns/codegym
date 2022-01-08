CC=clang
CFLAGS=-g -Wall
EXES=echo_server echo_client

all: $(EXES)

echo_server: echo_server.c
	$(CC) $(CFLAGS) -o $@ $^

echo_client: echo_client.c
	$(CC) $(CFLAGS) -o $@ $^

clean:
	rm -rf $(EXES) *.dSYM
