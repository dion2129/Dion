package main

import "fmt"

type Person struct{
	name string
	age int
}

type member struct {
	person Person
	position string
}

func main() {
	// // operator &
	// var orang1 person = person("daffa", "backend")
	// var orang2 person = orang1

	// orang2.name = "yaya"

	// *orang2 = &person("lila", "frontend")
	// fmt.Println(orang1)
	// fmt.Println(orang2)

	// embeded struct

	orang1 := member{}

	orang1.person.name = "daffa"
	orang1.person.age = 19
	orang1.position = "backend"

	fmt.Printf("%+v\n",orang1)
}