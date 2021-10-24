package models

import(
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
type User struct{
	Id int
	Name string
}
func GetAll() []User{
	var(
		id int
		name string
	)
	var listUser []User
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
		listUser = append(listUser,User{id,name})
	}
	return listUser
}