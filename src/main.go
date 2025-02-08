package main

import (
	"MathLib/lib/linalg"
	"MathLib/lib/consts"
	"fmt"
)

func main() {
	fmt.Println("In main")
	const x = consts.T
	fmt.Printf("Const T: %d\n", x)

	// Test linear algebra package
	v1 := linalg.VecFrom(2, 4, 6)
	fmt.Println("v1: ", v1)
	v2 := linalg.VecFrom(1, 3, 5)
	fmt.Println("v2: ", v2)

	fmt.Println("v1 norm: ", v1.Norm())

	v3, err := v1.Add(v2)

	if err != nil {
		fmt.Println("Shouldn't be possible")
	}
	
	fmt.Println("v1 + v2: ", v3)
	v3.NormInPlace()
	fmt.Println("v3 norm in place: ", v3)

	v4, err := v1.Dot(v2)

	if err != nil {
		fmt.Println("Shouldn't be possible")
	}

	fmt.Println("v2 dot v2: ", v4)
	
}
