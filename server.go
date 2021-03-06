package main

import (
	Forum "Forum/Forum_Package"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	Forum.InitDatabase("ForumDB.db")
	register, _ := template.ParseFiles("HTML/inscription.html")
	Home := template.Must(template.ParseFiles("HTML/Accueil.html"))
	login := template.Must(template.ParseFiles("HTML/connexion.html"))
	profil := template.Must(template.ParseFiles("HTML/profil.html"))

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		Home.Execute(rw, nil)
	})

	http.HandleFunc("/Profil", func(rw http.ResponseWriter, r *http.Request) {
		profil.Execute(rw, nil)
	})

	http.HandleFunc("/registerPage", func(rw http.ResponseWriter, r *http.Request) {
		register.Execute(rw, nil)
	})
	http.HandleFunc("/register", Forum.Register)

	http.HandleFunc("/loginPage", func(rw http.ResponseWriter, r *http.Request) {
		login.Execute(rw, nil)
	})
	http.HandleFunc("/login", Forum.Login)
	http.HandleFunc("/Logout", Forum.DeleteSession)

	http.HandleFunc("/GetPosts", Forum.GetPostHandlefunc)
	http.HandleFunc("/CreatePost", Forum.AddPostHandlefunc)
	http.HandleFunc("/GetSession", Forum.GetSession)

	http.HandleFunc("/GetComms", Forum.GetComsByPostId)
	http.HandleFunc("/CreateComms", Forum.AddCommsHandleFunc)

	http.HandleFunc("/LikePost", Forum.LikeAPost)

	fmt.Println("test1")
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("Start Server")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)

}
