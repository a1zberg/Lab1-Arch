default:test
clean:
	rmdir /s /q out
test:
	go vet ./implementation
	go test ./implementation
mainbuild:
	mkdir out
	git describe >> version.txt
	go build -o out/example
	out/example
	del version.txt