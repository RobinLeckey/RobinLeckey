package main


//builds web applications
import "net/http"
//process HTML templates
import "html/template"
//prints statements
import "fmt"
//Goroutines
import "time"


func adopt(a string){
  for i := 0; i < 10; i++ {
    time.Sleep(100 * time.Millisecond)
    fmt.Println(a)
  }
}

type GoatsPage struct{
  Title string
  Goats string
}

//First handler ~ indexHandler
func indexHandler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "<h1>Goats are cool animals<h2>")
}

//Second handler ~ goatsHandler
func goatsHandler(w http.ResponseWriter, r *http.Request){
  p := GoatsPage{Title: "Baby Goats in Colorado for Adoption", Goats: "Here are the goats that are available:"}
  t, err := template.ParseFiles("index2.gohtml")
  fmt.Println(err)
  t.Execute(w, p)
}

//Beginning of main
func main(){
  go adopt("Please adopt these baby goats")
  adopt("!!!")
  http.FileServer(http.Dir("./demo"))
  //Handles all requests to the web root
  http.HandleFunc("/", indexHandler)
  //Displays all information about adopting goats
  http.HandleFunc("/goat/", goatsHandler)
  //Tells the server to isten on the TCP network address
  http.ListenAndServe(":8080", nil)
}
//End of main
