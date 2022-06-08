package algorithms

func CheckIfPrime(isPrime int) bool {
	if isPrime == 0 || isPrime == 1 {
		return false
	}

	if isPrime == 2 {
		return true
	}

	for i := isPrime - 1; i > 1; i-- {
		if isPrime%i == 0 {
			return false
		}
	}
	return true
}

func GiveAllPrimes(n int) []int {
	r := []int{}
	for i := 0; i <= n; i++ {
		if res := CheckIfPrime(i); res {
			r = append(r, i)
		}
	}
	return r
}
