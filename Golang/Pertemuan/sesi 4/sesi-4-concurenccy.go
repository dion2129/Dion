package main
import ("fmt"
	// "time"
	"sync"
)

var wg sync.WaitGroup
func Printname(name string)string{
	for i := 0; i <5; i++{
		fmt.Println(name)
	}
	wg.Done()
	return name
}

func main(){
	wg.Add(2)

	go Printname("lila")
	
	Printname("daffa")
	
	go Printname("arif")
	// go func(){
	// 	fmt.Println("Bruh")
	// }()
	wg.Wait()

	// time.Sleep(5 * time.Second)
}
