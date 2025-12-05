package main

import (
	"fmt"
	"math"
	"strconv"
)

func main(){

	option := ""
	fmt.Scanln(option)




	n := 0
	m := 0
	fmt.Println("Enter the amount of digits to show:")
	fmt.Scanln(&n)
	if n > 32{
		fmt.Println("n too large")
		return
	}

	fmt.Println(picalc(n))

	fmt.Println("Enter the precision for e value:")
	fmt.Scanln(&m)

	fmt.Println(Ecalc(m))
	fmt.Println("Fibonacii numbers:")
	for i := 0; i<n; i++{
		fmt.Println(fibonacci(i))
	}
	
}

func picalc(n int)float64{

	pistr := strconv.FormatFloat(math.Pi, 'f', n, 64)
	val, _ := strconv.ParseFloat(pistr, 64)
	return val
}

func Ecalc(m int)float64{
	estr := strconv.FormatFloat(math.Exp(1), 'f', m, 64)
	val, _ := strconv.ParseFloat(estr, 64)
	return val

}

func fibonacci(n int) int {
	if n <= 1{
		return n
	}

	return fibonacci(n - 1) + fibonacci(n - 2)
}

