// Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you can't use additional data structures?

package chapter1

// IsUniqueSpace tests whether the given string composed entirely of unique characters optimized for space
func IsUniqueSpace(str string) bool {
	for i := range str {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return false
			}
		}
	}
	return true
}

// IsUniqueComplexity tests whether the given string composed entirely of unique characters optimized for complexity
func IsUniqueComplexity(str string) bool {
	present := make(map[rune]bool)
	for _, r := range str {
		if present[r] == true {
			return false
		}
		present[r] = true
	}
	return true
}
