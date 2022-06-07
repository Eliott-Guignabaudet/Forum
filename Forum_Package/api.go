package Forum

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Post struct {
	Id       int
	UserId   int
	Category string
	Title    string
	Content  string
}

func GetPostHandlefunc(w http.ResponseWriter, r *http.Request) {
	var posts []Post
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &posts)
	posts = tranlateSQLrowPosts(selectAllFromTable("posts"))
	jsonPosts, _ := json.Marshal(posts)
	w.Write(jsonPosts)
	fmt.Println(string(jsonPosts))
}

func AddPostHandlefunc(w http.ResponseWriter, r *http.Request) {
	var post Post

	if r.Method != "POST" {
		// page d'erreur
		fmt.Println("erreur methode : ", r.Method)
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	err := json.Unmarshal(body, &post)
	if err != nil {
		println("ERREUR : ", err)
	}
	fmt.Println(post)
	if post.Title == "" {
		// erreur titre vide
		fmt.Println(post)
		fmt.Println("titre vide")
		w.Write([]byte("{\"error\":\"titre vide\"}"))
	} else if post.Content == "" {
		// erreur contenu vide
		fmt.Println("contenu vide")
		w.Write([]byte("{\"error\":\"contenu vide\"}"))
	} else if post.Category == "" {
		// erreur contenu vide
		fmt.Println("contenu vide")
		w.Write([]byte("{\"error\":\"aucune category\"}"))
	} else {
		fmt.Println("good")
		addPost(post)
	}

}
