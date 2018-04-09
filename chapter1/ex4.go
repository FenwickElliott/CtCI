package chapter1

// PalendromePermutation checks if a string is a permutation of a palendrome
func PalendromePermutation(input string) bool {
	counts := make(map[rune]int)
	for _, r := range input {
		counts[r]++
	}
	odds := 0
	for _, v := range counts {
		if v%2 != 0 {
			odds++
			if odds == 2 {
				return false
			}
		}
	}
	return true
}
