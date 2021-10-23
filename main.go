package main

import(
	// "fmt"
	"net/http"
	"html/template"
	"encoding/json"
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"

)
type Student struct{
	Id int
	Name string
}

func main(){
	// fmt.Println("Server Running")
	// http.HandleFunc("/home",MyHome)
	http.HandleFunc("/list",AllStudent)
	http.ListenAndServe(":3000",nil)
	// getData()
}
func getData() []Student{
	var(
		id int
		name string
	)
	var listStudent []Student
	db,err := sql.Open("mysql","root:admin12345@tcp(127.0.0.1:3306)/EmployeeManager")
	if err != nil{
		log.Fatal(err)
	}else{
		log.Println("Sucessful")
	}
	defer db.Close()
	rows, err := db.Query("select id,name from employee")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
	for rows.Next(){
		err := rows.Scan(&id,&name)
		if err != nil{
			log.Fatal(err)
		}
		listStudent = append(listStudent,Student{id,name})
	}
	return listStudent
}
func MyHome(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	// fmt.Fprint(w,"My home")
	tmpl.Execute(w,nil)
}
func AllStudent(w http.ResponseWriter, r *http.Request){
	
	listStudent := getData()
	json.NewEncoder(w).Encode(listStudent)
}
