package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func fact(n *big.Int) {
	var inc = big.NewInt(1)
	var res = big.NewInt(1)
	var cmp = big.NewInt(0)
	for n.Cmp(inc) >= 0 {
		res = res.Mul(res, inc)
		inc = cmp.Add(inc, big.NewInt(1))
	}

	fmt.Println(res)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	n, e := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	if e != nil {
		panic(e)
	}
	fact(big.NewInt(n))
}

func readLine(r *bufio.Reader) string {
	line, _, _ := r.ReadLine()
	return strings.TrimRight(string(line), "\r\n")
}
