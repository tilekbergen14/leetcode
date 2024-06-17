package main

import (
	"math"
)

func judgeSquareSum(c int) bool {
	for i := 0; i <= int(math.Ceil(math.Sqrt(float64(c)))); i++ {
		res := c - i*i
		res2 := int(math.Ceil(math.Sqrt(float64(res))))
		if c-(res2*res2+i*i) == 0 {
			return true
		}
	}
	return false
}
