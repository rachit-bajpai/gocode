package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		s := fmt.Sprintf("%d : %s", index, arg)
		fmt.Println(s)
	}
}
