# const
IP := $(ip)
SERVER_PASSWORD := $(pass)
SCP_STORAGE_PATH := /root

all: cpuinfo cpuinfo_mac cpuinfo_linux_x86_64 cpuinfo_linux_aarch64 cpuinfo_windows

cpuinfo:
	mkdir -p build && cd build
	go build -o ./build/ci

cpuinfo_mac:
	mkdir -p build && cd build
	go build -o ./build/cpuinfo_mac

cpuinfo_linux_x86_64:
	mkdir -p build && cd build
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/cpuinfo_linux_x86_64

cpuinfo_linux_aarch64:
	mkdir -p build && cd build
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./build/cpuinfo_linux_aarch64

cpuinfo_windows:
	mkdir -p build && cd build
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./build/cpuinfo_windows