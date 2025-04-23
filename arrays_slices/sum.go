package main

func SumArray(numbers [5]int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumSlice(numbers []int) int{
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(nums ...[]int) []int { // nums ...[]int <- this takes into consideration that you can have multiple lists as inputs
	var sums []int

	for _, numbers  := range nums {
		sums = append(sums, SumSlice(numbers)) // appending to sums
	}
	return sums
}

func SumAllTails(nums ...[]int) []int {
	var sums []int

	for _, numbers := range nums {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, SumSlice(tail))
		}
	}
	return sums
}
