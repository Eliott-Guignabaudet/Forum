package Forum

import (
	"database/sql"
	"log"
)

func tranlateSQLrowPosts(rows *sql.Rows) []Post {
	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.Id, &post.UserId, &post.Category, &post.Title, &post.Content) //&post.Likes, &post.UsersWholiked

		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, post)
	}

	return posts

}

func tranlateSQLrowComs(rows *sql.Rows) []Commentaires {
	var comments []Commentaires
	for rows.Next() {
		var comment Commentaires
		err := rows.Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.Content)
		if err != nil {
			log.Fatal(err)
		}
		comments = append(comments, comment)
	}

	return comments
}
