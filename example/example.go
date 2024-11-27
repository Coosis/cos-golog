package main

import (
	cl "github.com/Coosis/cos-golog"
)

func main() {
	cl.Info("Hello, World!")
	cl.Warn("Hello, World!")
	cl.Error("Hello, World!")
	cl.Panic("Hello, World!")
	cl.Debug("Hello, World!")
}
