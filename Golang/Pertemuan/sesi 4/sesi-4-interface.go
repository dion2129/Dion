package main

import ("fmt"
	"math"
)
type lingkaran struct{
	Diameter float64
}
type Hitung interface{
	luasLingkaran() float64
	Keliling() float64
}

func (1 lingkaran) luasLingkaran() float64{
	r := float64(l.Diameter / 2)
	luas := math.Pi * (r * r)

	return luas
}

func (1 lingkaran) Keliling() float64{
	r := float64 (l.Diameter / 2)
	luas := 2 * math.Pi * r
	return luas
}

func main(){
	var person interface{}

	person = "daffa"
	person = 4
	fmt.Println(person)

	number := map[string]any{ //Any sama dengan interface
		"age":1,
		"age2":"12",
	}
	var lingkaran Hitung = lingkaran(2)
	fmt.Println(lingkaran.luasLingkaran)
}