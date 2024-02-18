package main

import (
	"fmt"

	"github.com/fleezesd/go-advanced/Concurrency"
)

func main() {
	info, _ := Concurrency.WeightBoundedCities("us", "zh")
	fmt.Println(info[0].TempC, info[1].TempC)
}
