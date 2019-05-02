package main

import "fmt"

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
	type contactInfo struct {
		email   string
		zipCode int
	}
	type person struct {
		firstName string
		lastName  string
		contact   contactInfo
	}
	p3 := person{
		firstName: "Alex",
		lastName:  "Stark",
		contact: contactInfo{
			email:   "sgwgw@css.com",
			zipCode: 45252,
		},
	}
	fmt.Printf("%+v\n", p3)
}
