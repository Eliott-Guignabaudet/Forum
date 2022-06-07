package Forum

import (
	"database/sql"
	"log"
	"net/http"
	"io/ioutil"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"encoding/json"
)

type UserParams struct {
	Id int
	Pseudo string `json:Pseudo`
	Email string `json:Email`
	Password string `json:Password` 
	ConfirmPW string `json:ConfirmPW`
}

func Register(rw http.ResponseWriter , r *http.Request) {
	var User UserParams
	// isTrue := true
	db := InitDatabase("ForumDB.db")
	defer db.Close()
	body , _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body , &User)

	//Check si tous les champs sont good
	if User.Pseudo == ""{
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

	if User.ConfirmPW == "" && User.Password != ""{
		fmt.Println("testGO3")
		rw.Write([]byte("{\"errorConfirmPW\" : \"confirmation du mot de passe nécessaire.\"}"))
		return
	}

	if User.ConfirmPW != User.Password{
		fmt.Println("testGO4")
		rw.Write([]byte("{\"errorNotSamePW\" : \"Le mot de passe n'est pas le même.\"}"))
		return
	}

	//insert into
	InsertIntoUsers(db , User.Pseudo , User.Email , User.Password)
}

// func Login(rw , http.ResponseWriter , r *http.Request) {

// }

func InitDatabase(database string) *sql.DB {
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
							PRAGMA foreign_keys = ON;
							CREATE TABLE IF NOT EXISTS users (
								id INTEGER PRIMARY KEY AUTOINCREMENT,
								pseudo TEXT NOT NULL,
								email TEXT NOT NULL UNIQUE,
								password TEXT NOT NULL
							);

							CREATE TABLE IF NOT EXISTS posts (
								id INTEGER PRIMARY KEY AUTOINCREMENT,
								user_id INTEGER NOT NULL,
								content TEXT NOT NULL,
								FOREIGN KEY (user_id) REFERENCES users(id)
							);
						 `
	_, err = db.Exec(sqlStmt)

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func InsertIntoUsers(db *sql.DB, pseudo string, email string, password string) (int64, error) {
	result, _ := db.Exec(`INSERT INTO users (name , email , password) VALUES (? , ? , ?) `, pseudo, email, password)
	return result.LastInsertId()
}

func insertIntoPosts(db *sql.DB, user_id int, content string) (int64, error) {
	result, _ := db.Exec(`INSERT INTO posts (user_id , content) VALUES (?,?) `, user_id, content)
	return result.LastInsertId()
}

// func selectUsersByEmail(db *sql.DB , email string) UserParams {
// 	var user UserParams
// 	db.QueryRow(`SELECT * FROM users WHERE email = ?` , email).Scan(&user.Id , &user.Name , &user.Email , &user.Password)
// 	return user.Id
// }