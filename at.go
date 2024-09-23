package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	ST := readline()
	words := strings.Fields(ST)
	S := words[0]
	T := words[1]
	// 文字分割を全パターン試す
	for w := 1; w < len(S); w++ {
		for c := 0; c < w; c++ {
			name := ""
			for i := c; i < len(S); i += w {
				name += string(S[i])
			}
			if name == T {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
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
