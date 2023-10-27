package main

import "fmt"

type deck []string

func(d deck) print(){
	for i, card := range d {
		fmt.Println(i,card)
	}
}

package main
 
import "fmt"
 
type book string
 
func (b book) printTitle() {
    fmt.Println(b)
}
 
func main() {
    var b book = "Harry Potter"
    b.printTitle()
}