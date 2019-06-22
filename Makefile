build: main.go
	go build -o main.out main.go

run:
	./main.out

clean:
	rm *.out