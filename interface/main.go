package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type engBot struct{}
type spaBot struct{}

//type logWriter struct{}

func main() {
	e := engBot{}
	print(e)

	//http
	/*
		resp, err := http.Get("http://example.com/")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		bs := make([]byte, 9999)
		resp.Body.Read(bs)
		fmt.Println(string(bs))

		b := logWriter{}
		io.Copy(b, resp.Body)
	*/
}

func (engBot) getGreeting() string {
	return "Hello"
}
func (spaBot) getGreeting() string {
	return "Oala"
}
func print(b bot) {
	fmt.Println(b.getGreeting())
}

/*
func (logWriter) write(bs []byte) (int, err) {
	return 1, nil
}
*/
