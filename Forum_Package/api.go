package Forum

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase(database string) *sql.DB {
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
							PRAGMA foreign_keys = ON;
							CREATE TABLE IF NOT EXISTS users (
								id INTEGER PRIMARY KEY AUTOINCREMENT,
								name TEXT NOT NULL,
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

func InsertIntoUsers(db *sql.DB, name string, email string, password string) (int64, error) {
	result, _ := db.Exec(`INSERT INTO users (name , email , password) VALUES (? , ? , ?) `, name, email, password)
	return result.LastInsertId()
}

func insertIntoPosts(db *sql.DB, user_id int, content string) (int64, error) {
	result, _ := db.Exec(`INSERT INTO posts (user_id , content) VALUES (?,?) `, user_id, content)
	return result.LastInsertId()
}
