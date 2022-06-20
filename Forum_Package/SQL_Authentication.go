package Forum

import "database/sql"

func InsertIntoGoogleUsers(db *sql.DB, name string, email string) (int64, error) {
	result, _ := db.Exec(`INSERT INTO googleUsers (name , email) VALUES (? , ?) `, name, email)
	return result.LastInsertId()
}

func InsertIntoFacebookUsers(db *sql.DB, name string, email string) (int64, error) {
	result, _ := db.Exec(`INSERT INTO facebookUsers (name , email) VALUES (? , ?) `, name, email)
	return result.LastInsertId()
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
