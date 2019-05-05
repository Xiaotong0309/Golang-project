package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	/*
		type person struct {
			firstName string
			lastName  string
		}
		//assign 1
		p1 := person{firstName: "Alex", lastName: "Stark"}
		fmt.Println(p1)
		//assign 2
		var p2 person
		p2.firstName = "Alex"
		p2.lastName = "Stark"
		fmt.Printf("%+v\n", p2)
	*/

	p3 := person{
		firstName: "Alex",
		lastName:  "Stark",
		contact: contactInfo{
			email:   "sgwgw@css.com",
			zipCode: 45252,
		},
	}
	fmt.Printf("%+v\n", p3)

	p3.update("Bill")
	fmt.Printf("%+v\n", p3)

}

func (p *person) update(name string) { // must outside main
	(*p).firstName = name
}
