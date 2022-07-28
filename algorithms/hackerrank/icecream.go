package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'whatFlavors' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY cost
 *  2. INTEGER money
 */

type CostIndex struct {
	c int32
	i int
}
type CostIndexList struct {
	list []CostIndex
}

func (l CostIndexList) Len() int {
	return len(l.list)
}

func (l CostIndexList) Less(i, j int) bool {
	return l.list[i].c > l.list[j].c
}

func (l CostIndexList) Swap(i, j int) {
	l.list[i], l.list[j] = l.list[j], l.list[i]
}

func whatFlavors(cost []int32, money int32) {
	// Write your code here
	var cil CostIndexList
	for i, c := range cost {
		cil.list = append(cil.list, CostIndex{c: c, i: i + 1})
	}
	sort.Sort(&cil)
	// fmt.Println(cil.list)
	for i := 0; i < len(cil.list); i++ {
		if cil.list[i].c < money {
			for j := i + 1; j < len(cil.list); j++ {
				if cil.list[i].c+cil.list[j].c == money {
					if cil.list[i].i > cil.list[j].i {
						fmt.Println(cil.list[j].i, cil.list[i].i)
					} else {
						fmt.Println(cil.list[i].i, cil.list[j].i)
					}

				}
			}
		}
	}
	//fmt.Println(cil.list)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		moneyTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		money := int32(moneyTemp)

		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		costTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var cost []int32

		for i := 0; i < int(n); i++ {
			costItemTemp, err := strconv.ParseInt(costTemp[i], 10, 64)
			checkError(err)
			costItem := int32(costItemTemp)
			cost = append(cost, costItem)
		}

		whatFlavors(cost, money)
	}
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
