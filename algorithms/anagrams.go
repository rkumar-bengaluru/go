package algorithms

/**
 * 2 string are anagrams if they have same characters
 * with same number of occurance.
 * Example:- danger & garden
 * d :=1, a :=1, n := 1 , g :=1, e:=1, r :=1
 * find occurance of each character in each string.
 *
 */

func checkIfAnagrams(s1 string, s2 string) bool {
	s1map := make(map[rune]int)
	s1slice := []rune(s1)
	for i := 0; i < len(s1slice); i++ {
		if s1map[s1slice[i]] != 0 {
			s1map[s1slice[i]] = s1map[s1slice[i]] + 1
		} else {
			s1map[s1slice[i]] = 1
		}
	}
	s2map := make(map[rune]int)
	s2slice := []rune(s2)
	for i := 0; i < len(s2slice); i++ {
		if s2map[s2slice[i]] != 0 {
			s2map[s2slice[i]] = s2map[s2slice[i]] + 1
		} else {
			s2map[s2slice[i]] = 1
		}
	}
	for k, v := range s1map {
		if s2map[k] != v {
			return false
		}
	}

	return true
}
