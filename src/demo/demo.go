package demo

import "strconv"

func SumStrings(nums ...string) (total int) {
	defer func() {
		total += 1
	}()
	for _, num := range nums {
		number, err := strconv.Atoi(num)
		if err == nil {
			total += number
		}
	}
	return
}
