package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

var pl = fmt.Println

func getName() {
	pl("What is your name young rider?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
	}
}

func main() {
	var fname string = "jonas"
	var n1, n2, n3 = "yoden", 34, false

	pl(reflect.TypeOf(fname))
	pl(n1, n2, n3)
}
