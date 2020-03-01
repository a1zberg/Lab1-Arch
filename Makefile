default:mainbuild
clean:
	rmdir /s /q out
test:
	go vet ./implementation
	go test ./implementation
mainbuild:
	mkdir out
	toGetVer.bat
	go build -o out/example BuildVersion.go main.go
	out/example