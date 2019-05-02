package main

import "fmt"

func main() {
	/*
		var should be together with type
		var card string = "Ace of Spades"
		fmt.Println(card)
		temp := "jkgkgk"
		fmt.Println(temp)
	*/
	/*
		cards := createDeck()
		cards.shuffle()
		cards.print()
	*/
	/*
		hand, remain := deal(cards, 5)
		hand.print()
		remain.print()
		hand.save("hand.txt")
		read("hand.txt").print()
	*/
	//assignment 1
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, e := range s {
		if e%2 == 0 {
			fmt.Printf("%v is even\n", e)
		} else {
			fmt.Printf("%d is odd\n", e)
		}
	}

}
