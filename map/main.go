package main

import "fmt"

func main() {
	/*
		m := map[string]string{
			"red":   "fgwg",
			"green": "fsgsg",
		}

		fmt.Println(m)
	*/
	m := make(map[int]string)
	m[10] = "egege"
	delete(m, 10)
	m[1] = "egege"
	m[11] = "egfde"
	m[12] = "egbtre"
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
}
