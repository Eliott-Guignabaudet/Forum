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

							CREATE TABLE IF NOT EXISTS commentaires (
								Id INTEGER PRIMARY KEY AUTOINCREMENT,
								UserId INTEGER NOT NULL,
								PostId INTEGER NOT NULL,
								Content TEXT NOT NULL,
								FOREIGN KEY (UserId) REFERENCES users(id)
								FOREIGN KEY (PostId) REFERENCES posts(Id)
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

func InsertIntoCommentaire(db *sql.DB, user_id int, post_id int, content string) (int64, error) {
	result, _ := db.Exec(`INSERT INTO commentaires (user_id, post_id, content) VALUES (? , ?, ? )`, user_id, post_id, content)
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
	result, err := db.Exec(`INSERT INTO posts (UserId, Category, Title, Content) VALUES (?, ?, ?, ?)`, post.UserId, post.Category, post.Title, post.Content)
	if err != nil {
		log.Fatal(err)
	}
	return result.LastInsertId()
}

func addComms(comment Commentaires) (int64, error) {
	db, err := sql.Open("sqlite3", "ForumDB.db")
	if err != nil {
		log.Fatal(err)
	}
	result, err := db.Exec(`INSERT INTO commentaires (UserId, PostId, Content) VALUES (?,?,?)`, comment.UserId, comment.PostId, comment.Content)
	if err != nil {
		log.Fatal(err)
	}
	return result.LastInsertId()
}

func selectUsersByEmailAndPW(db *sql.DB, email string, password string) UserParams {
	var user UserParams
	db.QueryRow(`SELECT * FROM users WHERE email = ? and password = ?`, email, password).Scan(&user.Id, &user.Pseudo, &user.Email, &user.Password)
	return user
}

// func selectUserByEmailAndPW(db *sql.DB, email string, password string) UserParams {
// 	var user UserParams
// 	db.QueryRow(`SELECT * FROM users WHERE email = ? and password = ?`, email, password).Scan(&user.Id, &user.Pseudo, &user.Email, &user.Password)
// 	fmt.Println("to sql:", user)
// 	return user
// }

func selectComByPostId(post_id int) *sql.Rows {
	db, err := sql.Open("sqlite3", "ForumDB.db")
	if err != nil {
		log.Fatal(err)
	}
	result, _ := db.Query(`SELECT * FROM commentaires WHERE PostId = ?`, post_id)
	return result
}

// func selectPostbyCategory(Category string) *sql.Rows {

// }
