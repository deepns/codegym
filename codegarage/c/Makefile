CC=clang
OPTIONS=-Wall
EXES=mtls_client mtls_server tls_server tls_client tcp_server tcp_client
OPENSSL_LIBS=-lssl -lcrypto

all: $(EXES)

# Using automatic variables for shortcuts
# $@ - file name of the target of the rule
# $< - name of the first prerequisite
# $? - names of all prerequisites that newer than target
# $^ - name of all prerequisites

mtls_client: mtls_client.c
	$(CC) $(OPTIONS) $(OPENSSL_LIBS) -o $@ $^

mtls_server: mtls_server.c
	$(CC) $(OPTIONS) $(OPENSSL_LIBS) -o $@ $^

tls_server: tls_server.c
	$(CC) $(OPENSSL_LIBS) -o $@ $^

tls_client: tls_client.c
	$(CC) $(OPENSSL_LIBS) -o $@ $^

tcp_server: tcp_server.c
	$(CC) -o $@ $^

tcp_client: tcp_client.c
	$(CC) -o $@ $^

clean:
	rm $(EXES)
