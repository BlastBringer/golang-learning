package main 
import "fmt"

func sumandaverage(a, b int) (int, int){
	return a+b , (a+b)/2
}

func makestarpyramid(n int){
	for i := 0; i <= n; i++{
		for j := 0; j <= i; j++{
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func Simpleinterest(principal, rate, time float64) float64{
	return (principal * rate * time) / 100
}

func Discount(total float64, discountpercentage float64) float64{
	return total * (discountpercentage / 100)
}

func swap(a, b *int) {
	temp := *a 
	*a = *b
	*b = temp 
}

