go-opencc
=========

opencc wrapper for Golang

### Install

Run 'go get github.com/stevenyao/go-opencc'

### Example

```go
package main

import "fmt"
import "github.com/stevenyao/go-opencc"

const (
	input = "中国鼠标软件打印机"
	config_s2t = "/usr/share/opencc/s2t.json"
	config_t2s = "/usr/share/opencc/t2s.json"
)

func main() {
	fmt.Println("Test Converter class:")
	c := opencc.NewConverter(config_s2t)
	defer c.Close()
	output := c.Convert(input)
	fmt.Println(output)
}
```
