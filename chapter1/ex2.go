// Check Permutation: Given two strings, write a method to decide if one is a permutation of the other.

package chapter1

import "reflect"

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
