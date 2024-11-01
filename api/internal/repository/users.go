package repository

import (
	"database/sql"
	"fmt"
	"golang-api/internal/models"
)

type users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

func (repo users) Get(nameOrNick string) ([]models.User, error) {

	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)
	rows, err := repo.db.Query("SELECT id, name, nickname, email, created_at FROM users where name LIKE ? OR nickname LIKE ?", nameOrNick, nameOrNick)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repo users) GetById(id uint64) (models.User, error) {

	row := repo.db.QueryRow("SELECT id, name, nickname, email, created_at FROM users WHERE id = ?", id)

	var user models.User
	if err := row.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (repo users) Create(user models.User) (uint64, error) {

	statement, err := repo.db.Prepare("INSERT INTO users (name, nickname, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

func (repo users) Update(id uint64, user models.User) error {

	statement, err := repo.db.Prepare("UPDATE users SET name = ?, nickname = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nickname, user.Email, id); err != nil {
		return err
	}

	return nil
}

func (repo users) Delete(id uint64) error {

	statement, err := repo.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(id); err != nil {
		return err
	}

	return nil
}

func (repo users) GetByEmail(email string) (models.User, error) {

	row := repo.db.QueryRow("SELECT id, password FROM users WHERE email = ?", email)

	var user models.User
	if err := row.Scan(&user.ID, &user.Password); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (repo users) Follow(userId, followerId uint64) error {
	statement, err := repo.db.Prepare("INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

func (repo users) Unfollow(userId, followerId uint64) error {
	statement, err := repo.db.Prepare("DELETE FROM followers WHERE user_id = ? AND follower_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

func (repo users) GetFollowers(userId uint64) ([]models.User, error) {
	rows, err := repo.db.Query(
		"SELECT u.id, u.name, u.nickname, u.email, u.created_at FROM users u "+
			"INNER JOIN followers f ON u.id = f.follower_id WHERE f.user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repo users) GetFollowing(userId uint64) ([]models.User, error) {
	rows, err := repo.db.Query(
		"SELECT u.id, u.name, u.nickname, u.email, u.created_at FROM users u "+
			"INNER JOIN followers f ON u.id = f.user_id WHERE f.follower_id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repo users) GetPassword(userId uint64) (string, error) {
	row := repo.db.QueryRow("SELECT password FROM users WHERE id = ?", userId)

	var password string
	if err := row.Scan(&password); err != nil {
		return "", err
	}

	return password, nil
}

func (repo users) UpdatePassword(userId uint64, password string) error {
	statement, err := repo.db.Prepare("UPDATE users SET password = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(password, userId); err != nil {
		return err
	}

	return nil
}
