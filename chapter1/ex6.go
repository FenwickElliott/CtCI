package chapter1

import "strconv"

// StringCompression implements a simple character count algorithm. If a character is returned once followed by the number of times it appeared in situ. If the compression is not shorter the input is returned.
func StringCompression(input string) string {
	if len(input) < 3 {
		return input
	}
	res := []byte{input[0]}
	c := 1
	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			c++
		} else {
			res = append(res, []byte(strconv.Itoa(c))...)
			res = append(res, input[i])
			c = 1
		}
	}
	res = append(res, []byte(strconv.Itoa(c))...)
	if len(res) < len(input) {
		return string(res)
	}
	return input
}
