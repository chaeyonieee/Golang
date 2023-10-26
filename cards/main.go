package main

import "fmt"

func main() {
	
	card :=[]string{"3 diamond", newCard()}
	card = append(card, "7 diamond")
for i , card :=range card{
	fmt.Println(i,card)
}

}


func newCard() string{
	return "5 diamond"
} 