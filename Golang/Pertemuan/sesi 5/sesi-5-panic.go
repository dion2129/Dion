package main

import "fmt"

func endApp(){
	fmt.Println("APK ERROR")
}
func runApp(error bool){
	defer endApp()

	if error{
		panic("APK error")
	}
	fmt.Println("Apk berjalan")
}

func main(){
	runApp(true)
}