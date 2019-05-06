package main

import (
	"bufio"
	"os"
	"fmt"
)

//计算重复的行
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	var i = 1
	for input.Scan() {
		counts[input.Text()]++
		if i == 4 {
			break
		}
		i++;
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
