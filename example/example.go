package main

import (
	cl "github.com/Coosis/cos-golog"
)

func pr() {
	cl.Debug("Hello, World!")
	cl.Info("Hello, World!")
	cl.Warn("Hello, World!")
	cl.Error("Hello, World!")
	cl.Panic("Hello, World!")
}

func main() {
	// default behavior
	pr()

	// set some properties
	cl.SetBrackets("_", "$")
	cl.SetDate(false)
	cl.SetCallerFullPath(true)
	pr()

	// omit the prefixes
	cl.SetDate(false)
	cl.SetCaller(false)
	cl.SetBracket(false)
	pr()
}
