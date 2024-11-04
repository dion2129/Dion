package main
import ("fmt"
	"os"
)
func main(){
	defer println("ini defer")
	fmt.Println("ini bukan defer")
	defer fmt.Println("ini defer ke 2")
	os.Exit(0)
}