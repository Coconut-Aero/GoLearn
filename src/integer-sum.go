package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	var result = 0
	for i := 1; i < len(args); i++ {
		var num, err = strconv.ParseInt(args[i], 10, 64)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if i != len(args)-1 {
			fmt.Print(args[i], "+")
		}
		if i == len(args)-1 {
			fmt.Print(args[i], "=")
		}
		result += int(num)
	}
	fmt.Print(result)
}
