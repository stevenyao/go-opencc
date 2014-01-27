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

	fmt.Println("Test Convert function:")
	s := opencc.Convert(input, config_s2t)
	fmt.Println(s)
	fmt.Println(opencc.Convert(s, config_t2s))

	fmt.Println("Test ConvertAsync function:")
	retChan := make(chan string)
	opencc.ConvertAsync(input, config_s2t, func(output string) {
		retChan <- output
	})

	fmt.Println(<- retChan)
}