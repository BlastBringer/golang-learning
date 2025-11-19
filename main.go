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
	room []Client 
	name string 

}

func (s *Server) Setroom(namestr string, roomsize int) {
	s.name = namestr 
	s.room = make([]Client, roomsize)
}

func (s *Server) Useradded (c Client){
	s.room = append(s.room, c)
	fmt.Sprintf("User added to the server : ", c.name)
	
}


func main(){
	fmt.Println("Hello world! Server running on http://localhost:8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello there!")
	})

	http.ListenAndServe(":8080", nil)


}