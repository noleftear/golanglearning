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
	fmt.Println(josh.firstName, josh.lastName, "is a", josh.info.age, "year old", josh.info.gender, "that weighs", josh.info.weight, "pounds.")
	fmt.Println(jeff.firstName, jeff.lastName, "is a", jeff.info.age, "year old", jeff.info.gender, "that weighs", jeff.info.weight, "pounds.")

}
