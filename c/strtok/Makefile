CC=gcc
CFLAGS=-Wall -g
EXES=strtok_ex
RM=rm -rf

all: $(EXES)

strtok_ex: strtok_ex.c
	$(CC) $(CFLAGS) -o $@ $^

clean:
	$(RM) $(EXES) *.dSYM

