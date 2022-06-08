package algorithms

func Reverse(stringtoreverse string) (string, error) {
	b := []rune(stringtoreverse)
	l := len(b)
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil
}
