package main

import "fmt"

func main() {
	//var cards string = "Ace of spaces"
	cards := []string{newCard(), "heart", "diamond"}
	fmt.Println(cards)

}

func newCard() string {
	return "Five of diamonds"
}
