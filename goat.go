package main

//process HTML templates
import "html/template"
//builds web applications
import "net/http"


var tpl *template.Template

func init(){
  tpl = template.Must(template.ParseGlob("C:/Users/rleckey/go/src/hello/templates/*.gohtml"))
}

func main(){
  //Handles all requests to the web root
  http.HandleFunc("/", hom)
  //http.HandleFunc("/about", abo)
  http.Handle("/stuff/", http.StripPrefix("/stuff", http.FileServer(http.Dir("./pics"))))
  //Tells the server to isten on the TCP network address
  http.ListenAndServe(":8080", nil)
}

//Sends data to the HTTP client
func hom(w http.ResponseWriter, r *http.Request){
  tpl.ExecuteTemplate(w, "default.gohtml", nil)
}

//func abo(w http.ResponseWriter, r *http.Request){
  //tpl.ExecuteTemplate(w, "about.gohtml", nil)
//}
