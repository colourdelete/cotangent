all: compile test clean;

compile:
	go build -buildmode=c-archive cotangent.go
	go build -buildmode=c-shared cotangent.go

test:
	gcc -o test -pthread test.c cotangent.a

clean:
	rm -f main *.o

.PHONY: compile clean
