CC=clang
CFLAGS=-g -Wall -Werror
EXES=signal_ex signal_multiple signal_with_return sigaction_ex sigio_socket
RM=rm -rf

all: $(EXES)

# $@ - file name of the target of the rule
# $^ - name of all prerequisites
signal_ex: signal_ex.c
	$(CC) $(CFLAGS) -o $@ $^

signal_multiple: signal_multiple.c
	$(CC) $(CFLAGS) -o $@ $^

signal_with_return: signal_with_return.c
	$(CC) $(CFLAGS) -o $@ $^

sigaction_ex: sigaction_ex.c
	$(CC) $(CFLAGS) -o $@ $^

sigio_socket: sigio_socket.c
	$(CC) $(CFLAGS) -o $@ $^

clean:
	$(RM) $(EXES) *.dSYM