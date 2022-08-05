package main

import "fmt"

func saveThePrisoner(n int32, m int32, s int32) int32 {
	r := (((m + n) % n) + (s % n) - 1) % n
	if r == 0 {
		return n
	}

	return ((r + n) % n)
}
func main() {
	fmt.Println(8 % 5)
	saveThePrisoner(3, 7, 3)
}
