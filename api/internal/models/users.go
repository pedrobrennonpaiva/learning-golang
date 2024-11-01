package models

import (
	"errors"
	"golang-api/internal/pkg/security"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (u *User) Prepare(step string) error {
	if err := u.validate(step); err != nil {
		return err
	}

	if err := u.format(step); err != nil {
		return err
	}

	return nil
}

func (u *User) validate(step string) error {
	if u.Name == "" {
		return errors.New("name is required and cannot be empty")
	}
	if u.Nickname == "" {
		return errors.New("nickname is required and cannot be empty")
	}
	if u.Email == "" {
		return errors.New("email is required and cannot be empty")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("email is invalid")
	}

	if step == "register" && u.Password == "" {
		return errors.New("password is required and cannot be empty")
	}
	return nil
}

func (u *User) format(step string) error {
	u.Name = formatString(u.Name)
	u.Nickname = formatString(u.Nickname)
	u.Email = formatString(u.Email)

	if step == "register" {
		hashedPassword, err := security.Hash(u.Password)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}

	return nil
}

func formatString(s string) string {
	return strings.TrimSpace(s)
}
