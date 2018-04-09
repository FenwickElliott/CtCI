package chapter1

import "reflect"

// Permutation checks if two strings are permutations of eachother.
func Permutation(in1, in2 string) bool {
	in1m := make(map[rune]int)
	in2m := make(map[rune]int)
	for _, r := range in1 {
		in1m[r]++
	}
	for _, r := range in2 {
		in2m[r]++
	}
	res := reflect.DeepEqual(in1m, in2m)
	if res {
		return true
	}
	return false
}
