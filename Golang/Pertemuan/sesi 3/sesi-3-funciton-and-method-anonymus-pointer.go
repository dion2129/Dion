package main

import "fmt"

type people struct{
	name string
	
}

func main(){
	orang1 := struct{
		position string
		person people
	} {}

	orang1.person.name = "daffa"
	orang1.position = "backend"

	orang2 := struct{
		position string
		person people
	} {
		person: people{name: "ado"}
		position : "frontend"
	}
	fmt.Printf(orang1)
}