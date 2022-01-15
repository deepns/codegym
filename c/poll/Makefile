CC=clang
CFLAGS=-g -Wall
EXES=poll_server poll_client

all: $(EXES)

poll_server: poll_server.c
	$(CC) $(CFLAGS) -o $@ $^

poll_client: poll_client.c
	$(CC) $(CFLAGS) -o $@ $^

clean:
	rm -rf $(EXES) *.dSYM

