CC=clang
OPTIONS=-Wall
EXES=uds_client uds_server

all: $(EXES)

# Using automatic variables for shortcuts
# $@ - file name of the target of the rule
# $< - name of the first prerequisite
# $? - names of all prerequisites that newer than target
# $^ - name of all prerequisites

uds_server: uds_server.c
	$(CC) -o $@ $^

uds_client: uds_client.c
	$(CC) -o $@ $^
	
clean:
	rm $(EXES)
