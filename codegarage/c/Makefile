CC=clang
EXES=tcp_client tls_client

all: $(EXES)

tcp_client: tcp_client.c
	$(CC) -o tcp_client tcp_client.c

tls_client: tls_client.c
	$(CC) -lssl -lcrypto -o tls_client tls_client.c

clean:
	rm $(EXES)
