CC=clang
CFLAGS=-g -Wall
EXES=nb_server nb_client

all: $(EXES)

nb_server: nb_server.c
	$(CC) $(CFLAGS) -o $@ $^

nb_client: nb_client.c
	$(CC) $(CFLAGS) -o $@ $^

clean:
	rm -rf $(EXES) *.dSYM

