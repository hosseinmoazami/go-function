package main

import (
	"fmt"

	"hossein.dev/simpleMath"
)

func main() {
	nums := []float64{8, 3, 5, 8, 10}
	fmt.Printf("8 + 3 + 5 + 8 + 10= %f\n", simpleMath.Sum(nums...))
	fmt.Printf("8 + 3 = %f\n", simpleMath.Add(8, 3))
	fmt.Printf("45 * 5 = %f\n", simpleMath.Multiply(45, 5))

	res, err := simpleMath.Divided(54, 8)
	if err != nil {
		fmt.Printf("An error occurred %s\n", err.Error())
	} else {
		fmt.Printf("54 / 8 = %f\n", res)
	}

	sv := simpleMath.NewSemanticVersion(1, 2, 3)
	fmt.Println(sv.String())
	sv.IncrementMajor()
	sv.IncrementMajor()
	sv.IncrementMinor()
	sv.IncrementPatch()
	fmt.Println(sv.String())

}
