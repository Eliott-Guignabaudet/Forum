package Forum

import (
	"database/sql"
	"log"
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
								Id INTEGER PRIMARY KEY AUTOINCREMENT,
								UserId INTEGER NOT NULL,
								Category TEXT NOT NULL,
								Title TEXT NOT NULL,
								Content TEXT NOT NULL,
								FOREIGN KEY (UserId) REFERENCES users(id)
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

func InsertIntoPosts(db *sql.DB, user_id int, content string) (int64, error) {
	result, _ := db.Exec(`INSERT INTO posts (user_id , content) VALUES (?,?) `, user_id, content)
	return result.LastInsertId()
}

func selectAllFromTable(table string) *sql.Rows {
	db, err := sql.Open("sqlite3", "ForumDB.db")
	if err != nil {
		log.Fatal(err)
	}
	result, _ := db.Query("SELECT * FROM " + table)
	return result
}

func addPost(post Post) (int64, error) {
	db, err := sql.Open("sqlite3", "ForumDB.db")
	if err != nil {
		log.Fatal(err)
	}
	result, _ := db.Exec(`INSERT INTO posts (UserId, Category, Title, Content) VALUES (?, ?, ?, ?)`, post.UserId, post.Category, post.Title, post.Content)
	return result.LastInsertId()
}
