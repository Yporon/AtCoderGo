package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func main10() {
	N := readInt()
	scores := make([][2]int, N+1) // インデックス0を初期状態として利用するために、サイズをN+1にしています。
	for i := 1; i <= N; i++ {
		row := readIntSlice()
		c, p := row[0]-1, row[1] // cを0ベースのインデックスに調整
		// scores[i]へのポインタを取得し、そのポインタを通じて値を更新
		current := &scores[i]
		*current = scores[i-1] // まず、前の要素の値をコピー
		(*current)[c] += p     // 次に、適切なカテゴリにポイントを加算
	}
	q := readInt()
	for i := 0; i < q; i++ {
		l, r := readInt2()
		fmt.Println(scores[r][0]-scores[l-1][0], scores[r][1]-scores[l-1][1])
	}
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func main14() {
	_ = readInt()
	//fmt.Println(N)
	aSlice := readIntSlice()
	bSlice := readIntSlice()
	sort.Slice(aSlice, func(i, j int) bool { return aSlice[i] < aSlice[j] })
	sort.Slice(bSlice, func(i, j int) bool { return bSlice[i] < bSlice[j] })
	count := 0
	for i, _ := range aSlice {
		count += abs(aSlice[i] - bSlice[i])
	}
	fmt.Println(count)
}
