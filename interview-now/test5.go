package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Printf("%f\n", 0.3)
	fmt.Printf("%.10f\n", 0.3)
	fmt.Printf("%.20f\n", 0.3)

	var a float32 = 1 << 24
	var b float32 = a+1
	fmt.Println(math.Float32bits(a)) // 1266679808
	fmt.Println(math.Float32bits(b)) // 1266679808
	fmt.Println(a == b)
	var c float64 = 0.1
	var d float64 = 0.2
	fmt.Println(c+d)
	fmt.Printf("%.20f\n", c+d)


	var f int = 123_456_321
	var g float64 = 3.1415_926
	fmt.Println(f, g)
}
