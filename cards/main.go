package main

func main() {

	card := deck{"3 diamond", newCard()}
	card = append(card, "7 diamond")
	// for i , card :=range card{
	// 	fmt.Println(i,card)
	// }

	card.print()
}

func newCard() string {
	return "5 diamond"
}