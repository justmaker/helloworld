package main

import (
	"bufio"
	"fmt"
	"helloworld/hello"
	"os"
	"strings"
)

func main() {
	fmt.Print("Ouput mode? [i]mport/[c]onst/[v]ariable: ")
	reader := bufio.NewReader(os.Stdin)
	delim := byte('\n')
	text, err := reader.ReadString(delim)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Print("text is: ", text)
	}

	hello_vstr := "Hello String Variable"
	const hello_cstr = "Hello"

	switch text = strings.TrimSpace(text); text {
	case "i", "I", "import", "Import":
		hello.Hello()
	case "c", "C", "const", "Const":
		fmt.Println(hello_cstr)
	case "v", "V", "variable", "Variable":
		fmt.Println(hello_vstr)
		fallthrough
	default:
		fmt.Println("Program exit.")
	}
}
