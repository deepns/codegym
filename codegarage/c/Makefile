CC=clang
EXES=tls_server tcp_server tcp_client tls_client
OPENSSL_LIBS=-lssl -lcrypto

all: $(EXES)

# Using automatic variables for shortcuts
# $@ - file name of the target of the rule
# $< - name of the first prerequisite
# $? - names of all prerequisites that newer than target
# $^ - name of all prerequisites

tcp_server: tcp_server.c
	$(CC) -o $@ $^

tcp_client: tcp_client.c
	$(CC) -o $@ $^

tls_server: tls_server.c
	$(CC) $(OPENSSL_LIBS) -o $@ $^

tls_client: tls_client.c
	$(CC) $(OPENSSL_LIBS) -o $@ $^

clean:
	rm $(EXES)
