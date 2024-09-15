package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	S := readline()
	Rindex := strings.Index(S, "R")
	Mindex := strings.Index(S, "M")
	if Rindex < Mindex {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func readline() string {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000000)
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
