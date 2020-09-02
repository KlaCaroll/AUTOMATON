package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(len(os.Args))
	fmt.Printf("%#v\n", os.Args)

	var NumberOfTurns = os.Args[1]

	for i := 2; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}

	_ = NumberOfTurns
}

