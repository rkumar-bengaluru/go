package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

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
func main() {
	file, err := os.Open("test.txt")
	reader := bufio.NewReaderSize(file, 16*1024*1024)

	stdout, err := os.Create("t.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := isValid(s)
	fmt.Println(result)
	fmt.Println(1 % 2)
	result = isValid("aaaaabc")
	fmt.Println(result)
	result = isValid("aabbccddeefghi")
	fmt.Println(result)
	result = isValid("abbccc")
	fmt.Println(result)
	result = isValid("aabbc")
	fmt.Println(result)
	result = isValid("aaaabbcc")
	fmt.Println(result)
	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}
func check(n []int) bool {
	// 4 2 2
	var m = make(map[int]int)
	for _, v := range n {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v] = m[v] + 1
		}
	}
	
	fmt.Println(m)
	return true
}

func isValid(s string) string {
	// Write your code here
	if len(s) == 1 {
		return "YES"
	}
	var m = make(map[rune]int)

	r := []rune(s)
	for _, v := range r {
		if _, ok := m[v]; ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}
	fmt.Println(m)
	var c = make([]int, len(m))
	var i int
	for _, v := range m {
		c[i] = v
		i++
	}
	if check(c) {
		return "YES"
	}
	return "NO"
}
