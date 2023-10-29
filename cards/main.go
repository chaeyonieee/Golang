package main

func main() {

	// cards := newDeck()

	// // for i , card :=range card{
	// // 	fmt.Println(i,card)
	// // }

	// cards.print()

	// greeting := "Hi there"
	// fmt.Println([]byte(greeting))
// 	cards:= newDeck()
// cards.saveToFile("my_cards")
cards:= newDeck()
cards.shuffle()
cards.print()
}
