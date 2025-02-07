package main

import (
	"MathLib/lib/linalg"
	"MathLib/lib/consts"
	"fmt"
)

func main() {
	fmt.Println("In main")
	fmt.Println(linalg.Test())
	const x = consts.T
	fmt.Printf("Const T: %d", x)
}
