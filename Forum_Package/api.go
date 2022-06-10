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
	Pseudo    string
	Email     string
	Password  string
	ConfirmPW string
}

// type Lgin struct {
// 	EmailLog    string
// 	PasswordLog string
// }

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

	db := InitDatabase("ForumDB.db")
	defer db.Close()
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &User)
	fmt.Println("Register : ", User)

	//Check si tous les champs sont good
	if User.Pseudo == "" {
		fmt.Println("testGO")
		rw.Write([]byte("{\"errorPseudo\" : \"Pseudo nécessaire\"}"))
		return
	} else if User.Email == "" {
		fmt.Println("testGO1")
		rw.Write([]byte("{\"errorEmail\" : \"Email requis.\"}"))
		return
	} else if User.Password == "" {
		fmt.Println("testGO2")
		rw.Write([]byte("{\"errorPassword\" : \"password requis.\"}"))
		return
	} else if User.ConfirmPW == "" && User.Password != "" {
		fmt.Println("testGO3")
		rw.Write([]byte("{\"errorConfirmPW\" : \"confirmation du mot de passe nécessaire.\"}"))
		return
	}

	if User.ConfirmPW != User.Password {
		fmt.Println("testGO4")
		rw.Write([]byte("{\"errorNotSamePW\" : \"Le mot de passe n'est pas le même.\"}"))
		return
	} else {
		fmt.Println("c'est good")
		rw.Write([]byte("{\"CorrectRegister\" : \"true\"}"))
	}

	//insert into
	// if registerError(User, rw) {
	// 	http.Redirect(rw, r, "/loginPage", http.StatusFound)
	// }
	InsertIntoUsers(db, User.Pseudo, User.Email, User.Password)

}

func Login(Rw http.ResponseWriter, Rq *http.Request) {
	var Users UserParams
	isCorrectUser := false
	db := InitDatabase("ForumDB.db")
	defer db.Close()
	fmt.Println(Users.Email)
	body, _ := ioutil.ReadAll(Rq.Body)
	json.Unmarshal(body, &Users)

	UserInfo := selectUsersByEmailAndPW(db, Users.Email, Users.Password)
	fmt.Println("Log In:", Users, Users.Email, Users.Password)
	fmt.Println(UserInfo, UserInfo.Id)
	if UserInfo.Email == Users.Email && UserInfo.Password == Users.Password {
		isCorrectUser = true
		fmt.Println("OUI OUI OUI")
	} else {
		isCorrectUser = false
		fmt.Println("Pleure bébou")
	}

	if isCorrectUser {
		fmt.Println("Sa marche")
	}

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
