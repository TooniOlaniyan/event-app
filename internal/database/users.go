package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type UserModel struct {
	DB *sql.DB
}

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"`
}

func (m *UserModel) Insert(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := " INSERT INTO users (email , password , name) VALUES ($1,$2,$3) RETURNING id "

	return m.DB.QueryRowContext(ctx, query, user.Email, user.Password, user.Name).Scan(&user.Id)

}
func (m *UserModel) Get(userId int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := " SELECT id , email , name , password FROM users WHERE id = $1"

	var user User
	err := m.DB.QueryRowContext(ctx, query, userId).Scan(&user.Id, &user.Email, &user.Name, &user.Password)
	if err != nil {
		fmt.Println("Is this where the error is")
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil

}
func (m *UserModel) GetUserByEmail(userEmail string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := " SELECT id , email , name , password FROM users WHERE email = $1"

	var user User
	err := m.DB.QueryRowContext(ctx, query, userEmail).Scan(&user.Id, &user.Email, &user.Name, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil

}
