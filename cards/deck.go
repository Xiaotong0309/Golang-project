package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func createDeck() deck {
	cards := deck{}
	cardSuits := deck{"Diamonds", "Spades", "Hearts", "Clubs"}
	cardValues := deck{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, numHand int) (deck, deck) {
	return d[:numHand], d[numHand:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) save(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func read(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
