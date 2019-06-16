all:
	go build -o acmdl *.go

ensure:
	dep ensure

test:
	go test ./lib

install:
	go install github.com/d2verb/acmdl

clean:
	rm -f acmdl
