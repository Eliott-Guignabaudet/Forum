package main

import (
	"html/template"
	"fmt"
	"net/http"
	Forum "Forum/Forum_Package"
)


func main() {
	db := Forum.InitDatabase("ForumDB.db")
	register := template.Must(template.ParseFiles("HTML/inscription.html"))
	Home := template.Must(template.ParseFiles("HTML/Accueil.html"))
	login := template.Must(template.ParseFiles("HTML/connexion.html"))
	defer db.Close()
	fmt.Println("test")

	http.HandleFunc("/" , func(rw http.ResponseWriter , r *http.Request) {
		Home.Execute(rw , nil)
	})

	http.HandleFunc("/register" , func(rw http.ResponseWriter , r *http.Request){
		if r.Method != http.MethodPost {
			register.Execute(rw , nil)
			return
		}
		name := r.FormValue("Name")
		email := r.FormValue("Email")
		password := r.FormValue("Password")
		
		Forum.InsertIntoUsers(db , name , email, password)
		// Forum.InsertIntoUsers(db , "Mathieu" , "Mathieu@gmail.com" , "abcde")
		// Forum.InsertIntoUsers(db , "Lucas" , "Lulu@gmail.com" , "klmno")

		// http.Redirect(rw , r , "/" , http.StatusFound)
	})

	http.HandleFunc("/login" , func(rw http.ResponseWriter , r *http.Request) {
		login.Execute(rw , nil)
	})


	fmt.Println("test1")
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("test2")
	http.ListenAndServe(":8080", nil)
	fmt.Println("test3")
}