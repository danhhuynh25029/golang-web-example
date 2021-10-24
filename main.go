package main

import(
	"fmt"
	"net/http"
	"html/template"
	"encoding/json"
	// "database/sql"
	// "log"
	"module/models"
	// _ "github.com/go-sql-driver/mysql"

)
func main(){
	fmt.Println("Server Running")
	http.HandleFunc("/home",MyHome)
	http.HandleFunc("/list",AllStudent)
	http.HandleFunc("/check",check)
	http.ListenAndServe(":3000",nil)
	// s := models.String()
	// fmt.Println(s)
	// // getData()
}

func MyHome(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	// fmt.Fprint(w,"My home")
	listUser := models.GetAll()
	tmpl.Execute(w,listUser)
}
func AllStudent(w http.ResponseWriter, r *http.Request){
	listUser := models.GetAll()
	json.NewEncoder(w).Encode(listUser)
}
func check(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		fmt.Println(r.FormValue("name"))
	}else{
		tmpl := template.Must(template.ParseFiles("templates/form.html"))
		tmpl.Execute(w,nil)
	}
}