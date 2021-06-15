package arrays

// Sum sums the passed array in a for range loop
func Sum(array [6]int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

// SumAlt sums the passed slice in a simple for loop
func SumAlt(slice []int) int {
	result := 0
	for i := 0; i < len(slice); i++ {
		result += slice[i]
	}
	return result
}

// SumAll takes varying numbers of slices, returning a new slice containing the totals for each slice passed in
func SumAll(slices [][]int) []int{
	sum := 0
	var result = []int {}
	for _, slice := range slices {
		for _, integer := range slice {
			sum += integer
		}
		result = append(result, sum)
		sum = 0
	}
	return result
}

// SumAllAlt takes varying numbers of slices, returning a new slice containing the totals for each slice passed in via a different implementation

func SumAllAlt(slices ...[]int) []int{
	var result = make([]int, 0)
	for _, v := range slices {
		result = append(result, SumAlt(v))
	}
	return result
}

func SumAllTails(numbersToSum...[]int) []int {
	var sum = make([]int, 0)
	for _,numbers := range numbersToSum {
		sum = append(sum, SumAlt(numbers[1:]))
	}
	return sum
}