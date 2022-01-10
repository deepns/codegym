CC=clang
OPTIONS=-Wall
EXES=vio
TESTFILES=vio_outfile

all: $(EXES)

# Using automatic variables for shortcuts
# $@ - file name of the target of the rule
# $< - name of the first prerequisite
# $? - names of all prerequisites that newer than target
# $^ - name of all prerequisites

vio: vio.c
	$(CC) -o $@ $^
	
clean:
	rm $(EXES) $(TESTFILES)
