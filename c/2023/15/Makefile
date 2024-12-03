FLAGS = -Wall -W -O2

all:main

main: main.o hashtable.o
	gcc -o main main.o hashtable.o

test_main: main.o hashtable.o
	gcc -o test_main -DTEST main.c hashtable.o

main.o: main.c main.h
	gcc ${FLAGS} -c main.c

hashtable.o: hashtable.c hashtable.h
	gcc ${FLAGS} -c hashtable.c

run:main
	./main

test: test_main
	./test_main

clean:
	rm -f *.o main