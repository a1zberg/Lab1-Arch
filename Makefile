default:mainbuild
clean:
	rmdir /s /q out
test:
	go vet ./implementation
	go test ./implementation
mainbuild:	implementation/implementation.go main.go
	mkdir out
	git describe >> version.txt
	go build -o out/example
	out/example
	del version.txt