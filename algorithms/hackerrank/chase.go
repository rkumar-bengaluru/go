package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the catAndMouse function below.
func catAndMouse(x int32, y int32, z int32) string {
	// A--M--B
	// x--y--z
	// xd = x-z
	var xd int32
	if x > z {
		xd = x - z
	} else if z > x {
		xd = z - x
	}
	var yd int32
	if y > z {
		yd = y - z
	} else if z > y {
		yd = z - y
	}

	if xd > yd {
		return "Cat B"
	} else if yd > xd {
		return "Cat A"
	}
	return "Mouse C"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create("tmp.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		xyz := strings.Split(readLine(reader), " ")

		xTemp, err := strconv.ParseInt(xyz[0], 10, 64)
		checkError(err)
		x := int32(xTemp)

		yTemp, err := strconv.ParseInt(xyz[1], 10, 64)
		checkError(err)
		y := int32(yTemp)

		zTemp, err := strconv.ParseInt(xyz[2], 10, 64)
		checkError(err)
		z := int32(zTemp)

		result := catAndMouse(x, y, z)
		fmt.Println(result)

		fmt.Fprintf(writer, "%s\n", result)
	}

	writer.Flush()
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
