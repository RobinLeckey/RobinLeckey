package main

import "fmt"
import "log"
import "net/http"
import "io/ioutil"
import "html/template"
import "regexp"

//First handler
func handler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hello friends! This is the first demo for the standup meeting today %s!", r.URL.Path[1:])
}

//Page data structure
type Page struct {
  Title string
  Body []byte
}

//Storage for the Page structure
func (p *Page) save() error{
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
}

//Save Handler that edits the pages
func saveHandler(w http.ResponseWriter, r *http.Request, title string){
  body := r.FormValue("body")
  p := &Page{Title: title, Body: []byte(body)}
  err := p.save()
  if err != nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

//Render template
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page){
  err := templates.ExecuteTemplate(w, tmpl + ".html", p)
  if err != nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

//Load the pages
func loadPage(title string) (*Page, error){
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)
  //if ReadFile encounters an error
  if err != nil {
    return nil, err
  }
  return &Page{Title: title, Body: body}, nil
}

//View Handler that views the pages
func viewHandler(w http.ResponseWriter, r *http.Request, title string){
  p, err := loadPage(title)
  if err != nil {
    http.Redirect(w, r, "/edit/"+title, http.StatusFound)
    return
  }
  renderTemplate(w, "view", p)
}

//Edit Handler that edits the pages
func editHandler(w http.ResponseWriter, r *http.Request, title string){
  p, err := loadPage(title)
  if err != nil {
    p = &Page{Title: title}
  }
  renderTemplate(w, "edit", p)
}

//Valid Path
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

//Make Handler
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc{
  return func(w http.ResponseWriter, r *http.Request) {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
      http.NotFound(w, r)
      return
    }
    fn(w, r, m[2])
  }
}

//Beginning of main
func main(){
  //First handler
  http.HandleFunc("/view/", makeHandler(viewHandler))
  http.HandleFunc("/edit/", makeHandler(editHandler))
  http.HandleFunc("/save/", makeHandler(saveHandler))
  //Fatal warning
  log.Fatal(http.ListenAndServe(":8080", nil))
}
//End of main
