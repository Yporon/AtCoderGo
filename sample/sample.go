package sample

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func readInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
func readInt2() (int, int) {
	return readInt(), readInt()
}

func readInt3() (int, int, int) {
	return readInt(), readInt(), readInt()
}

func readInt4() (int, int, int, int) {
	return readInt(), readInt(), readInt(), readInt()
}

func readIntSlice(arg ...int) []int {
	n := arg[0]
	t := 0
	if len(arg) == 2 {
		t = arg[1]
	}

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = readInt() - t
	}
	return a
}

func readStr() string {
	sc.Scan()
	return sc.Text()
}

//
// NOTE:変換
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

// NOTE:便利関数
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

// UnionFind
// NOTE:UnionFind
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
