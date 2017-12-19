package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Print(input)

	var i int
	if _, err := fmt.Scan(&i); err != nil {
		log.Print("failed", err)
		return
	}
	fmt.Println(i)

	var i, j, k int
	if _, err := fmt.Scan(&i, &j, &k); err != nil {
		log.Print("  Scan for i, j & k failed, due to ", err)
		return
	}

	fmt.Println(i, j, k)
}
