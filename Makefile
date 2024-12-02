EXEC_NAME=fakeapi.exe
run.debug:
	go run .

build.linux.amd64:
	@echo "Building for Linux"
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o ${EXEC_NAME} .

build.windows:
	@echo "Building for Windows"
	@GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o ${EXEC_NAME} .

build.darwin:
	@echo "Building for Darwin"
	@GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o ${EXEC_NAME} .

build.darwinM:
	@echo "Building for Darwin ARM"
	@GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o ${EXEC_NAME} .

build.linux.arm:
	@echo "Building for Linux"
	@GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-s -w" -trimpath -o ${EXEC_NAME} .