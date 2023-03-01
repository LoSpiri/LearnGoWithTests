package arrays

func Sum(numbers []int) int {
	var result int
	for _, v := range numbers {
		result += v
	}
	return result
}

func SumAll(arrays ...[]int) []int {
	resultArray := []int{}
	for _, j := range arrays {
		resultArray = append(resultArray, Sum(j))
	}
	return resultArray
}

func SumAllTails(arrays ...[]int) []int {
	resultArray := []int{}
	for _, j := range arrays {
		if len(j) == 0 {
			resultArray = append(resultArray, 0)
		} else {
			tail := j[1:]
			resultArray = append(resultArray, Sum(tail))
		}
	}
	return resultArray
}
