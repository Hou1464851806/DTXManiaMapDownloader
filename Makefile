.PHONY: clean all
all: build-linux build-windows build-macos
build-windows:
	GOOS=windows GOARCH=amd64 go build -o build/DTXMapDownloader_windows.exe
build-macos:
	GOOS=darwin GOARCH=arm64 go build -o build/DTXMapDownloader_darwin
build-linux:
	GOOS=linux GOARCH=amd64 go build -o build/DTXMapDownloader_linux
clean:
	rm -rf build