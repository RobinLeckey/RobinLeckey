package main


//builds web applications
import "net/http"
//process HTML templates
import "html/template"
//prints statements
import "fmt"
//To demonstrate goroutine
import "time"

//Goroutine
func adopt(a string){
  for i := 0; i < 10; i++ {
    time.Sleep(100 * time.Millisecond)
    fmt.Println(a)
  }
}

//Datastucture
type GoatsPage struct{
  Title string
  Goats string
  Note string
}

//First handler ~ indexHandler
// localhost:8080/
func indexHandler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "<h1>Goats are cool animals<h2>")
}

//Second handler ~ goatsHandler
// localhost:8080/goat/
func goatsHandler(w http.ResponseWriter, r *http.Request){
  p := GoatsPage{Title: "Baby Goats in Colorado for Adoption", Goats: "Here are the goats that are available:"}
  t, err := template.ParseFiles("index2.gohtml")
  fmt.Println(err)
  t.Execute(w, p)
}

//Beginning of main
func main(){
  //"go" starts a goroutine
  go adopt("This is a demo")
  adopt("!!!")
  fmt.Println("Please adopt a goat or two.")
  http.FileServer(http.Dir("./demo"))
  //Handles all requests to the web root
  http.HandleFunc("/", indexHandler)
  //Displays all information about adopting goats
  http.HandleFunc("/goat/", goatsHandler)
  //Tells the server to isten on the TCP network address
  http.ListenAndServe(":8080", nil)
}
//End of main
