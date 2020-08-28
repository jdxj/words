fileName=words.out

build:
	go build -o $(fileName) *.go

clean:
	rm -rvf $(fileName)