package main 


import (
	"net/http"
	"fmt"
)

func apiHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello world!")
}
func main(){
	mux :+ http.newM
}