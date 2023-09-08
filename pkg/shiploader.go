package shiploader

func Utils(input string) (result string) {
	for _, c := range input {
		result = string(c) + result
	}
	return result
}
