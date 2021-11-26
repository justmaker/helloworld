package main

import (
	"fmt"
	"helloworld/hello"
)

func main() {
	hello_vstr := "Hello String Variable"
	const hello_str = "Hello"
	hello.Hello()
	fmt.Println(hello_vstr, hello_str)
}
