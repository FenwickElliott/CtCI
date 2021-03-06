package chapter1

// URLify replaces all spaces in a string with '%20'
func URLify(input string) string {
	res := []rune{}
	subOut := []rune(" ")[0]
	subIn := []rune("%20")
	for _, r := range input {
		if r == subOut {
			res = append(res, subIn...)
		} else {
			res = append(res, r)
		}
	}
	return string(res)
}
