package repository

import (
	"database/sql"
	"golang-api/internal/models"
)

type Posts struct {
	db *sql.DB
}

func NewPostsRepository(db *sql.DB) *Posts {
	return &Posts{db: db}
}

func (p Posts) CreatePost(post models.Post) (uint64, error) {
	statement, err := p.db.Prepare("INSERT INTO posts (title, content, author_id) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertId), nil
}

func (p Posts) GetPostById(postId uint64) (models.Post, error) {
	rows, err := p.db.Query("SELECT p.*, u.nickname, count(l.id) as likes FROM posts p INNER JOIN users u ON u.id = p.author_id LEFT JOIN likes l on l.post_id = p.id WHERE p.id = ? GROUP BY p.id", postId)
	if err != nil {
		return models.Post{}, err
	}
	defer rows.Close()

	var post models.Post

	if rows.Next() {
		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.CreatedAt,
			&post.AuthorNickname,
			&post.Likes); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

func (p Posts) GetPosts(user uint64) ([]models.Post, error) {
	rows, err := p.db.Query(`
		SELECT distinct p.*, u.nickname, (SELECT COUNT(*) FROM likes WHERE post_id = p.id) as likes
		FROM posts p 
		INNER JOIN users u ON u.id = p.author_id 
		INNER JOIN followers f ON f.user_id = p.author_id 
		WHERE u.id = ? or f.follower_id = ?
		ORDER BY 1 DESC
	`, user, user)
	if err != nil {
		return []models.Post{}, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.CreatedAt,
			&post.AuthorNickname,
			&post.Likes,
		); err != nil {
			return []models.Post{}, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (p Posts) UpdatePost(post models.Post) error {
	statement, err := p.db.Prepare("UPDATE posts SET title = ?, content = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(post.Title, post.Content, post.ID); err != nil {
		return err
	}

	return nil
}

func (p Posts) DeletePost(postId uint64) error {
	statement, err := p.db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(postId); err != nil {
		return err
	}

	return nil
}

func (p Posts) GetPostsByUser(userId uint64) ([]models.Post, error) {
	rows, err := p.db.Query(`
		SELECT p.*, u.nickname, count(l.id) as likes 
		FROM posts p 
		INNER JOIN users u ON u.id = p.author_id 
		LEFT JOIN likes l on l.post_id = p.id 
		WHERE p.author_id = ? 
		GROUP BY p.id
	`, userId)
	if err != nil {
		return []models.Post{}, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.CreatedAt,
			&post.AuthorNickname,
			&post.Likes,
		); err != nil {
			return []models.Post{}, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (p Posts) LikePost(postId uint64, userId uint64) error {
	statement, err := p.db.Prepare("INSERT INTO likes (post_id, user_id) SELECT ?, ? WHERE NOT EXISTS (SELECT 1 FROM likes WHERE post_id = ? AND user_id = ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postId, userId, postId, userId); err != nil {
		return err
	}

	return nil
}

func (p Posts) UnlikePost(postId uint64, userId uint64) error {
	statement, err := p.db.Prepare("DELETE FROM likes WHERE post_id = ? AND user_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postId, userId); err != nil {
		return err
	}

	return nil
}
