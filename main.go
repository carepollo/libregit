package main

import "fmt"

type Maiki struct{}

func (m *Maiki) Lol() {
	fmt.Println("a")
}

func main() {
	var maiki *Maiki
	maiki.Lol()
}
