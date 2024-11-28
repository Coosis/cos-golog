# A simple lib for golang logging
## Usage
```bash
go get github.com/Coosis/cos-golog
```
```go
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
```
You can also find the example in the example folder as well.
