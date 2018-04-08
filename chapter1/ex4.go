// Palendrome Permutation: Is the given string a permutation of a palendrome

package chapter1

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
