package main

import "fmt"

func main() {
	sp := NewStartPosition()
	fmt.Println(sp.String())
	fmt.Println(sp.WhitePawnBB.Draw())
}
