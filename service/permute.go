package service

func Permute(input string) []string {
	var result []string
	Backtrack([]rune(input), 0, &result)
	return result
}

func Backtrack(arr []rune, start int, result *[]string) {
	if start == len(arr)-1 {
		*result = append(*result, string(arr))
		return
	}

	for i := start; i < len(arr); i++ {
		arr[start], arr[i] = arr[i], arr[start]
		Backtrack(arr, start+1, result)
		arr[start], arr[i] = arr[i], arr[start]
	}
}
