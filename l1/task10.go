package main

import (
	"fmt"
	"math"
)

func main() {
	degrees := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)
	for _, d := range degrees {
		var key int
		if d >= 0 {
			key = int(math.Floor(d/10)) * 10
		} else {
			key = int(math.Ceil(d/10)) * 10
		}
		groups[key] = append(groups[key], d)
	}
	for key, values := range groups {
		fmt.Println(key, values)
	}
}
