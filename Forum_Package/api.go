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

type UserParams struct {
	Id        int
	Pseudo    string `json:Pseudo`
	Email     string `json:Email`
	Password  string `json:Password`
	ConfirmPW string `json:ConfirmPW`
}

func GetPostHandlefunc(w http.ResponseWriter, r *http.Request) {
	var posts []Post
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &posts)
	posts = tranlateSQLrowPosts(selectAllFromTable("posts"))
	jsonPosts, _ := json.Marshal(posts)
	w.Write(jsonPosts)
}

func Register(rw http.ResponseWriter, r *http.Request) {
	var User UserParams
	// isTrue := true
	db := InitDatabase("ForumDB.db")
	defer db.Close()
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &User)

	//Check si tous les champs sont good
	if User.Pseudo == "" {
		fmt.Println("testGO")
		rw.Write([]byte("{\"errorPseudo\" : \"Pseudo nécessaire\"}"))
		return
	}

	if User.Email == "" {
		fmt.Println("testGO1")
		rw.Write([]byte("{\"errorEmail\" : \"Email requis.\"}"))
		return
	}

	if User.Password == "" {
		fmt.Println("testGO2")
		rw.Write([]byte("{\"errorPassword\" : \"password requis.\"}"))
		return
	}

	if User.ConfirmPW == "" && User.Password != "" {
		fmt.Println("testGO3")
		rw.Write([]byte("{\"errorConfirmPW\" : \"confirmation du mot de passe nécessaire.\"}"))
		return
	}

	if User.ConfirmPW != User.Password {
		fmt.Println("testGO4")
		rw.Write([]byte("{\"errorNotSamePW\" : \"Le mot de passe n'est pas le même.\"}"))
		return
	}

	//insert into
	InsertIntoUsers(db, User.Pseudo, User.Email, User.Password)

}

func Login(Rw http.ResponseWriter, Rq *http.Request) {
	var Users UserParams
	isTrue := false
	db := InitDatabase("ForumDB.db")
	defer db.Close()
	Id := selectUsersByEmailAndPW(db, Users.Email, Users.Password)
	body, _ := ioutil.ReadAll(Rq.Body)
	json.Unmarshal(body, &Users)

	if Id > 0 {
		isTrue = true
	}

	if isTrue {
		fmt.Println("Sa marche")
	}

	fmt.Println("test", Users.Email, Users.Password, isTrue)
}

func AddPostHandlefunc(w http.ResponseWriter, r *http.Request) {
	var post Post

	if r.Method != "POST" {
		// page d'erreur
		fmt.Println("erreur methode : ", r.Method)
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &post)
	if err != nil {
		println("ERREUR : ", err)
	}
	fmt.Println(post)
	fmt.Println("Title", post.Title)
	fmt.Println("Content", post.Content)
	fmt.Println("Category", post.Category)
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
		fmt.Println("Aucune categorie")
		w.Write([]byte("{\"error\":\"aucune categorie\"}"))
	} else {
		fmt.Println("api: ", post)
		addPost(post)
	}
}
