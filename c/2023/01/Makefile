all:main

main:main.o
	gcc -o main main.o

main.o:main.c
	gcc -Wall -W -c main.c

run:main
	./main

clean:
	rm -f main.o main