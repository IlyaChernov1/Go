package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {

	fmt.Println(s)

	const n = 50_000_000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(math.Sin(n))

}
