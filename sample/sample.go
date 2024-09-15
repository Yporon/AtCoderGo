package sample

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rdr = bufio.NewReaderSize(os.Stdin, 10000000)

//
// 入力
//

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

//
// 変換
//

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

// UnionFind
type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent: parent, rank: rank}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] == x {
		return x
	}
	uf.parent[x] = uf.Find(uf.parent[x])
	return uf.parent[x]
}

func (uf *UnionFind) Unite(x, y int) {
	xRoot := uf.Find(x)
	yRoot := uf.Find(y)
	if xRoot == yRoot {
		return
	}

	if uf.rank[xRoot] < uf.rank[yRoot] {
		uf.parent[xRoot] = yRoot
	} else {
		uf.parent[yRoot] = xRoot
		if uf.rank[xRoot] == uf.rank[yRoot] {
			uf.rank[xRoot]++
		}
	}
}

func (uf *UnionFind) Same(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

//
// 便利関数
//

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
