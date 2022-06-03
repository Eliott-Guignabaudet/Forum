package Forum

import (
	"database/sql"
	"log"
)

func tranlateSQLrowPosts(rows *sql.Rows) []Post {
	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.Id, &post.UserId, &post.Category, &post.Content, &post.Title)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, post)
	}

	return posts

}
