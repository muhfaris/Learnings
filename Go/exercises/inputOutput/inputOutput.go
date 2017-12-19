package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputreader := bufio.NewReader(os.Stdin)
	input, _ := inputreader.ReadString('\n')
	fmt.Print(input)
}
