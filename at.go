package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// 標準入力から整数2つを読み取る（今回は使わないので破棄）
	_, _ = readInt2()

	// 標準入力から整数スライスを読み取る
	A := readIntSlice()
	B := readIntSlice()

	// スライスAとBを昇順にソートする
	sort.Ints(A)
	sort.Ints(B)

	ans := 0 // 結果の累積和を保存する変数

	// Bの各要素bについて処理を行う
	for _, b := range B {
		ok := false // 条件を満たすAの要素が見つかったかどうかを示すフラグ

		// Aの各要素aについてb以上の要素を探す
		for j, a := range A {
			if b <= a {
				// b <= a ならば条件を満たす
				ans += a    // aを結果に加算
				A = A[j+1:] // Aの先頭を現在の要素以降に更新
				ok = true   // 条件を満たしたのでフラグを更新
				break       // 内側のループを終了
			}
		}

		if !ok {
			// 条件を満たすaが見つからなかった場合は-1を出力して終了
			fmt.Println("-1")
			return
		}
	}

	// 結果の累積和を出力
	fmt.Println(ans)
}

// 以下、補助関数

var rdr = bufio.NewReaderSize(os.Stdin, 10000000)

// readline は標準入力から1行を読み取って返す関数
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

// readIntSlice は標準入力からスペース区切りの整数を読み取ってスライスとして返す関数
func readIntSlice() []int {
	slice := make([]int, 0)
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		slice = append(slice, toI(v))
	}
	return slice
}

// readInt は標準入力から1つの整数を読み取って返す関数
func readInt() int {
	return toI(readline())
}

// readInt2 は標準入力からスペース区切りの2つの整数を読み取ってタプルとして返す関数
func readInt2() (int, int) {
	lines := strings.Split(readline(), " ")
	return toI(lines[0]), toI(lines[1])
}

// toI は文字列を整数に変換する関数。変換に失敗した場合はpanicを発生させる
func toI(s string) int {
	v, ok := strconv.Atoi(s)
	if ok != nil {
		panic("Faild : " + s + " can't convert to int")
	}
	return v
}
