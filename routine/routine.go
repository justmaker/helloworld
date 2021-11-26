package main

import (
	"fmt"
	"sync"
	"time"
)

func take_money() {
	balance := money
	time.Sleep(3000 * time.Millisecond)
	balance -= 100
	money = balance
}

func foo() {
	take_money()
	fmt.Println("Foo: ", money)
	wg.Done()
}

func bar() {
	take_money()
	fmt.Println("Bar: ", money)
	wg.Done()
}

var wg sync.WaitGroup
var money int = 2000

func main() {
	wg.Add(2)
	go foo()
	//time.Sleep(0.0005 * time.Millisecond)
	go bar()
	wg.Wait()
	//fmt.Scanln() // wait for Enter Key
}
