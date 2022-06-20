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
								Content TEXT NOT NULL,` +
		// Likes INTEGER,
		// UsersWholiked TEXT,
		`
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

func selectUserNameById(id int) string {
	db, err := sql.Open("sqlite3", "ForumDB.db")
	if err != nil {
		log.Fatal(err)
	}
	var name string
	err = db.QueryRow(`SELECT name FROM users WHERE id = ?`, id).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	return name
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

func incrementPostsLikes(post_id int) *sql.Rows {
	db, err := sql.Open("sqlite3", "ForumDB.db")
	if err != nil {
		log.Fatal(err)
	}
	result, _ := db.Query(`UPDATE posts SET Likes = Likes + 1 WHERE PostId = ?`, post_id)
	return result
}
