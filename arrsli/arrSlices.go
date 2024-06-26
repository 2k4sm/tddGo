package arrsli

func Sum(values []int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}

	return sum
}

func SumAll(valuesArr ...[]int) (retArr []int) {
	retArr = make([]int, len(valuesArr))

	for idx, values := range valuesArr {
		sum := 0
		for _, value := range values {
			sum += value
		}
		retArr[idx] = sum
	}

	return retArr
}

func SumAllTails(valuesArr ...[]int) (retArr []int) {
	retArr = make([]int, len(valuesArr))
	var tail []int
	for idx, values := range valuesArr {
		if len(values) < 1 {
			tail = []int{}
		} else {
			tail = values[1:]
		}

		retArr[idx] = Sum(tail)
	}

	return retArr
}
