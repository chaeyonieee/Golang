package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func newDeck() deck {

		cards := deck{}

		cardSuits := []string{"spades","diamond","hearts"}
	cardValues := []string{"Ace","two","three"}

	for _, suit := range cardSuits{
		for _, value := range cardValues {
			cards = append(cards,suit+" of "+value)

		}

	}
	return cards
	}




func(d deck) print(){
	for i, cards := range d {
		fmt.Println(i,cards)
	}
}

func(d deck) toString() string{
	return strings.Join([]string(d), ",")
	
}

func(d deck) saveToFile(filename string) error{
	return ioutil.WriteFile(filename,[]byte(d.toString()),0666)
}