# const
IP := $(ip)
SERVER_PASSWORD := $(pass)
SCP_STORAGE_PATH := /root

all: filelock filelock_mac filelock_linux_x86_64 filelock_linux_aarch64 filelock_windows

filelock:
	mkdir -p build && cd build
	go build -o ./build/ci

filelock_mac:
	mkdir -p build && cd build
	go build -o ./build/filelock_mac

filelock_linux_x86_64:
	mkdir -p build && cd build
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/filelock_linux_x86_64

filelock_linux_aarch64:
	mkdir -p build && cd build
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./build/filelock_linux_aarch64

filelock_windows:
	mkdir -p build && cd build
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./build/filelock_windows