package main 
func Sum(numbers []int) int{
	sum := 0
	for _, number := range numbers{
		sum += number
	}
	return sum
}

func SumAllTails(numberstoSum ...[]int)[]int{
	var sums []int
	for _, numbers := range numberstoSum{
		if len(numbers) == 0{
			sums = append(sums, 0)
		} else {
		tails := numbers[1:]
		sums = append(sums, Sum(tails))
	}
}
	return sums
}