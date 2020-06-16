package main


//builds web applications
import "net/http"
//process HTML templates
import "html/template"
import "fmt"


type GoatsPage struct{
  Title string
  Goats string
}

//First handler
func indexHandler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "<h1>Goats are cool animals<h2>")
}

func goatsHandler(w http.ResponseWriter, r *http.Request){
  p := GoatsPage{Title: "Cute Goats in Colorado Springs", Goats: "Here are some goats"}
  t, err := template.ParseFiles("index2.gohtml")
  fmt.Println(err)
  t.Execute(w, p)
}
//Beginning of main
func main(){
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/goat/", goatsHandler)
  http.ListenAndServe(":8080", nil)
}
//End of main
