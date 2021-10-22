package main

import(
	"fmt"
	"net/http"
	"html/template"
	"encoding/json"

)
type Student struct{
	Id int
	Name string
}
func main(){
	fmt.Println("Server Running")
	http.HandleFunc("/home",MyHome)
	http.HandleFunc("/api",API)
	http.ListenAndServe(":3000",nil)
}
func MyHome(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	// fmt.Fprint(w,"My home")
	tmpl.Execute(w,nil)
}
func API(w http.ResponseWriter, r *http.Request){
	listStudent := []Student{
		Student{1,"Danh"},
		Student{2,"TheShy"},
	}
	json.NewEncoder(w).Encode(listStudent)
}