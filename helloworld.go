package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"helloworld/shape"

	"github.com/justmaker/helloworld/hello"
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

	c := shape.Circle{2}
	s := shape.Square{5}
	fmt.Println("Circle{2}.Area= ", c.Area(), "Square{5}.Area()= ", s.Area())

	fmt.Scanln() // wait for Enter Key
}
