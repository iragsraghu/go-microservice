package services

func FilterSlice(input []int, filter func(int) bool) []int {
	var result []int
	for _, v := range input {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}
