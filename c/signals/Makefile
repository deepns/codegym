CC=clang
CFLAGS=-g -Wall -Werror
EXES=signal signal_multiple
RM=rm -rf

all: $(EXES)

# $@ - file name of the target of the rule
# $^ - name of all prerequisites
signal: signal.c
	$(CC) $(CFLAGS) -o $@ $^

signal_multiple: signal_multiple.c
	$(CC) $(CFLAGS) -o $@ $^

clean:
	$(RM) $(EXES) *.dSYM