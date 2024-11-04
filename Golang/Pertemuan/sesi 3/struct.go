package main

import "fmt"

type person struct{

	name string
	age int
	position string
}
// initializing struct

func main() {
	orang2 := person{name: "daffa", age: 19, position: "backend"}
	var orang = person{}

	orang.name = "daffa"
	orang.age = 19
	orang.position = "backend"

	fmt.Println(orang.name)
	fmt.Println(orang.age)
	fmt.Println(orang.position)
	fmt.Printf("%+v",orang2) // pakai field yang ditulis
	fmt.Printf(orang2)

}

