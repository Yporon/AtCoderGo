package main

import (
	"fmt"
	"sort"
)

func main() {
	N := ni()
	AList := nis(N)
	WList := nis(N)

	WMap := make(map[int]int)
	for i, v := range WList {
		WMap[i] = v
	}

	WAListMap := map[int][]int{
		0: {},
		1: {},
		2: {},
		3: {},
	}

	for i := 0; i < N; i++ {
		Ai := AList[i]
		WAListMap[Ai-1] = append(WAListMap[Ai-1], WMap[i])
	}
	sum := 0
	for _, v := range WAListMap {
		if len(v) < 2 {
			continue
		}
		sort.Ints(v)
		for _, vv := range v[:len(v)-1] {
			sum += vv
		}
	}
	fmt.Println(sum)
}

// ==================================================
// io
// ==================================================
