package main

import "fmt"

type personInfo struct {
	age    int
	weight int
	gender string
}

type person struct {
	firstName string
	lastName  string
	info      personInfo
}

func main() {
	var jeff person
	jeff.firstName = "Jeff"
	jeff.lastName = "Hurst"
	jeff.info = personInfo{age: 31, weight: 180, gender: "male"}
	josh := person{firstName: "Joshua", lastName: "Hurst", info: personInfo{age: 29, weight: 225, gender: "male"}}
	josh.printPerson()
	jeff.printPerson()
}

func (p person) printPerson() {
	fmt.Println(p.firstName, p.lastName, "is a", p.info.age, "year old", p.info.gender, "that weighs", p.info.weight, "pounds.")
}
