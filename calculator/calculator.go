package calculator

import (
	"math"
)
func Sum(firstnum, secondnum float64)float64 {
	return firstnum + secondnum
}
func Sub(firstnum, secondnum float64)float64{
	return firstnum - secondnum
}
func Multiplication(firstnum, secondnum float64)float64{
	return firstnum * secondnum
}
func Division(firstnum, secondnum float64)float64{
	return firstnum / secondnum
}
func power(firstnum, secondnum float64)float64{
	return firstnum ** secondnum
}
func sqrt(firstnum, secondnum float64)float64{
	return math.sqrt(firstnum)
}