package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
 * Complete the 'libraryFine' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER d1
 *  2. INTEGER m1
 *  3. INTEGER y1
 *  4. INTEGER d2
 *  5. INTEGER m2
 *  6. INTEGER y2
 */

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
	// Write your code here
	fmt.Printf("d1=%v d2=%v m1=%v m2=%v y1=%v y2=%v\n", d1, d2, m1, m2, y1, y2)
	var res int32
	// check for year
	if (y1 - y2) > 0 {
		res = 10000
	} else {
		// same year
		if (y1 - y2) < 0 {
			res = 0
		} else if (m1 - m2) > 0 {
			res = (m1 - m2) * 500
		} else {
			if (m1 - m2) < 0 {
				res = 0
			} else if (d1 - d2) > 0 {
				res = (d1 - d2) * 15
			} else {
				res = 0
			}
		}
	}
	fmt.Println(res)
	return res
}

func main() {

	libraryFine(2, 7, 1014, 1, 1, 1015)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
