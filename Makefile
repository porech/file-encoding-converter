# All
all: linux windows darwin

# Client build commands
linux-i386:
	GOOS=linux GOARCH=386 ./build-scripts/build.sh
linux-amd64:
	GOOS=linux GOARCH=amd64 ./build-scripts/build.sh
linux-arm:
	GOOS=linux GOARCH=arm ./build-scripts/build.sh
windows-i386:
	GOOS=windows GOARCH=386 ./build-scripts/build.sh
windows-amd64:
	GOOS=windows GOARCH=amd64 ./build-scripts/build.sh
darwin-amd64:
	GOOS=darwin GOARCH=amd64 ./build-scripts/build.sh

# Platform-specific builds
linux: linux-i386 linux-amd64 linux-arm
windows: windows-i386 windows-amd64
darwin: darwin-amd64
