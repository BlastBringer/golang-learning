package main 


import(
	"fmt"
	"net/http"
)

type Client struct{
	Id int 
	name string
}

type Server struct{
	room
}

func (c Client) Useradded (){
	fmt.Sprintf("User added to the server : ", c.name)
}


func main(){
	fmt.Println("Hello world! Server running on http://localhost:8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello there!")
	})

	http.ListenAndServe(":8080", nil)


}