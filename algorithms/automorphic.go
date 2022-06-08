package algorithms

// a number is automorphic if the quare of the number ends
// with the same number. Example := 25*25 = 625
func CheckIfAutomorphic(n int) bool {
	sq := n * n

	for n > 0 {
		remFromSq := sq % 10
		remFromN := n % 10
		if remFromN != remFromSq {
			return false
		}
		n = n / 10
		sq = sq / 10
	}
	return true
}
