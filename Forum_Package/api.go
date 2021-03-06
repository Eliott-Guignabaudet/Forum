package Forum

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
)

type Post struct {
	Id            int
	UserId        int
	UserName      string
	Category      string
	Title         string
	Content       string
	Likes         int
	UsersWholiked string
}

type UserParams struct {
	Id        int
	Pseudo    string
	Email     string
	Password  string
	ConfirmPW string
}

type Commentaires struct {
	Id      int
	UserId  int
	PostId  int
	Content string
}

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func GetPostHandlefunc(w http.ResponseWriter, r *http.Request) {
	var posts []Post
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &posts)
	posts = tranlateSQLrowPosts(selectAllFromTable("posts"))
	for i := range posts {
		posts[i].UserName = selectUserNameById(posts[i].UserId)
	}

	jsonPosts, _ := json.Marshal(posts)
	w.Write(jsonPosts)
}

func Register(rw http.ResponseWriter, r *http.Request) {

	var User UserParams
	if r.Method != "POST" {
		// page d'erreur
		fmt.Println("erreur methode : ", r.Method)
		http.Redirect(rw, r, "/", http.StatusFound)
		return
	}
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
	if Rq.Method != "POST" {
		// page d'erreur
		fmt.Println("erreur methode : ", Rq.Method)
		http.Redirect(Rw, Rq, "/", http.StatusFound)
		return
	}
	isCorrectUser := false
	db := InitDatabase("ForumDB.db")
	defer db.Close()
	fmt.Println(Users.Email)
	body, _ := ioutil.ReadAll(Rq.Body)
	json.Unmarshal(body, &Users)

	UserInfo := selectUsersByEmailAndPW(db, Users.Email, Users.Password)
	if UserInfo.Email == Users.Email && UserInfo.Password == Users.Password {
		isCorrectUser = true
	} else {
		isCorrectUser = false
		fmt.Println("Pleure bébou")
	}
	fmt.Println(UserInfo.Id)
	if isCorrectUser {
		session, _ := store.Get(Rq, "session")
		res, _ := json.Marshal(UserInfo)
		session.Values["authenticated"] = string(res)
		session.Save(Rq, Rw)
		Rw.Write([]byte("{\"resp\":\"login!\",\"id\":" + strconv.Itoa(UserInfo.Id) + ",\"pseudo\":\"" + UserInfo.Pseudo + "\", \"email\":\"" + UserInfo.Email + "\" }"))
		//Rw.Write([]byte("{\"resp\":\"login!\"}"))

	}

}

func AddPostHandlefunc(w http.ResponseWriter, r *http.Request) {
	var post Post
	if r.Method != "POST" {
		// page d'erreur
		fmt.Println("erreur methode : ", r.Method)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &post)
	if err != nil {
		println("ERREUR : ", err)
	}
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
		addPost(post)
		w.Write([]byte("{\"resp\":\"poste créer\"}"))
	}
}

func GetComsByPostId(w http.ResponseWriter, r *http.Request) {
	var comments []Commentaires
	var comment Commentaires
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &comment)
	fmt.Println("hey postId : ", comment.PostId)
	comments = tranlateSQLrowComs(selectComByPostId(comment.PostId))
	jsonComs, _ := json.Marshal(comments)
	w.Write(jsonComs)
}
func AddCommsHandleFunc(w http.ResponseWriter, r *http.Request) {
	var comment Commentaires
	if r.Method != "POST" {
		// page d'erreur
		fmt.Println("erreur methode : ", r.Method)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &comment)
	if err != nil {
		println("ERREUR : ", err)
	}
	if comment.Content == "" {
		w.Write([]byte("{\"error\":\"contenu vide\"}"))
	} else {
		addComms(comment)
	}
}

func LikeAPost(w http.ResponseWriter, r *http.Request) {
	var postliked Post
	println("1", postliked.Id)
	if r.Method != "POST" {
		// page d'erreur
		fmt.Println("erreur methode : ", r.Method)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &postliked)
	println("2", postliked.Id)
	incrementPostsLikes(postliked.Id)
}
