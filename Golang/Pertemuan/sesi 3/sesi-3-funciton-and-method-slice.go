package main

import "fmt"

type people2 struct {
	name string
	position string
}

func main(){
	// slice
	var orang1= []people2{
		{name: "daffa", position: "BE"},
		{name: "Gopal", position: "BE"},
		{name: "bruh", position: "BE"},
	}
	
	for _, value := range orang1 {
		fmt.Println(value)
	}

	// anonymus struct of slcie
	orang2 := []struct{
		position string
		people people2
	} {
		{		
			{ position: "BE", people: people2{name: "bruh"}}
			{ position: "BE", people: people2{name: "bruh"}}
			{ position: "BE", people: people2{name: "bruh"}}
		}	
	}
	for _, value := range orang2{
		fmt.Println(value)
	}
}