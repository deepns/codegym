CC=clang
OPTIONS=-Wall

all: logger_test

logger_test: logger_test.c logger.c
	$(CC) $(OPTIONS) -o $@ $^

logger.o : logger.c
	$(CC) $(OPTIONS) -c $^

clean:
	rm *.o logger_test