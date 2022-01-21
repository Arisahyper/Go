package lib

func Average(s []int) int {
	var sum int
	for _, v := range s {
		sum += v
	}
	return sum / len(s)
}
