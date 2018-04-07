// Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you can't use additional data structures?

package ex1

// IsUnique tests whether the given string composed entirely of unique characters
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
