default:mainbuild
clean:
	rmdir /s /q out
	rmdir /s /q -p
test:
	go vet ./implementation
	go test ./implementation
mainbuild:	implementation/implementation.go main.go
	mkdir -p out
	go build -o out/example
	out/example