package main
import (
	"fmt"
	"net/http"
)

var PORT = ":8080"

func main(){
	http.HandleFunc("/web", great)

	http.ListenAndServe(PORT, nil)
}

func great(w http.ResponseWriter, r *http.Request){
	msg := "hellado"
	fmt.Fprintln(w, msg)
}