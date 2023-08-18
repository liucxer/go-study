package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/shirou/gopsutil/v3/cpu"
)

func main() {
	info, _ := cpu.Info()
	spew.Dump(info)
}
