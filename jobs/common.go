package jobs

func InArray(arr []string, str string) bool {
	var result bool = false
	for _, item := range arr {
		if item == str {
			result = true
		}
	}
	return result
}
