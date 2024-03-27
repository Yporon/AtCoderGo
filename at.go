package main

import "fmt"

type Dog struct {
	Name string
}

func (d *Dog) Speak() {
	fmt.Printf("Woof")
}

func main() {
	var nilPoint *Dog
	nilPoint.Speak()
}
