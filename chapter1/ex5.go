package chapter1

// OneAway evaluates whether two strings are within on 1 permutation of eachother
func OneAway(in1, in2 string) bool {
	if len(in1) > len(in2) {
		return (OneAway(in2, in1))
	}
	if len(in2)-len(in1) > 1 {
		return false
	}

	muts := 0
	shift := 0
	for i, r := range in1 {
		if r != []rune(in2)[i+shift] {
			muts++
			if muts == 2 {
				return false
			}
			if r == []rune(in2)[i+1] {
				shift++
			}
		}
	}
	return true
}
