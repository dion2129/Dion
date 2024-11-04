package main

import (
	"reflect"
	"fmt"
	"math"
)

func main(){
	number := 1
	valuenumb := reflect.ValueOf(number)
	fmt .Println(valuenumb)
	fmt .Println(valuenumb)

	student := struct {
		Name string
	}{}
	u := models.getUser()

	student.Name= "dafa"
	fmt.Println(student.ame)
	
	fmt.Println(reflect.TypeOf(student))
	
}