package sample

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rdr = bufio.NewReaderSize(os.Stdin, 10000000)

func readline() string {
	buf := make([]byte, 0, 16)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			fmt.Println(e.Error())
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func readIntSlice() []int {
	slice := make([]int, 0)
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		slice = append(slice, toI(v))
	}
	return slice
}

func readInt() int {
	return toI(readline())
}

func readInt2() (int, int) {
	lines := strings.Split(readline(), " ")
	return toI(lines[0]), toI(lines[1])
}

func readInt3() (int, int, int) {
	lines := strings.Split(readline(), " ")
	return toI(lines[0]), toI(lines[1]), toI(lines[2])
}

func toI(s string) int {
	v, ok := strconv.Atoi(s)
	if ok != nil {
		panic("Faild : " + s + " can't convert to int")
	}
	return v
}

func toS(i int) string {
	return strconv.Itoa(i)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
