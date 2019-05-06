package main

import (
	"fmt"
	"strings"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1:])
	fmt.Println(strings.Join(os.Args[1:], ","))

	for i,e := range os.Args {
		//Itoa指的是integer to ASCII string。
		fmt.Println(strconv.Itoa(i) + "=" + e)
	}

}
