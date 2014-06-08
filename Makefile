all: movii

movii:
	go build -o movii ./src

clean:
	rm -f movii
