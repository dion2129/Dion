package main

import "fmt"

// func hello(name string, age int) string {
// 	var result string = fmt.Sprintf("nama %s %d", name, age)
// 	return result
// }

func calculate(number float64) (float64, float64){
	multiple := number * number
	divide := number /number
	return multiple, divide
}
func main(){
	// respon := hello("daffa", 19)
	// fmt.Println(respon)

	number = 5.0
	multipleresult, divideresult := calculate(number)
	fmt.Println("Multiple: ", multipleresult)
	fmt.Println("divide: ", divideresult)
}