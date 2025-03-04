package models

import (
	"RenomachiBack/db"
	"RenomachiBack/utils"
	"database/sql"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	UID       string    `json:"uid"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ユーザを追加
func (user *User) AddUser() error {

	UID := utils.GenerateUserID()

	query := "INSERT INTO users (user_id, user_name, password, email) VALUES (?, ?, ?, ?)"
	result, err := db.DB.Exec(query, UID, user.Name, user.Password, user.Email)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)
	return nil
}

// ユーザを編集
func UpdateUser(user *User, user_id string) error {
	query := "UPDATE users SET user_name = ?, password = ?, email = ? WHERE user_id = ?"
	_, err := db.DB.Exec(query, user.Name, user.Password, user.Email, user.UID)
	return err
}

// ユーザを削除
func DeleteUser(user_id string) error {
	query := "DELETE FROM users WHERE user_id = ?"
	_, err := db.DB.Exec(query, user_id)
	return err
}

// ユーザ一覧を表示
func GetUsers() ([]User, error) {
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.UID, &user.Name, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// UIDによるユーザを取得
func GetUser(user_id string) (*User, error) {
	query := "SELECT id, user_id, user_name, password, email FROM users WHERE user_id = ?"
	row := db.DB.QueryRow(query, user_id)

	var user User
	err := row.Scan(&user.ID, &user.UID, &user.Name, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // 該当するユーザーがいない場合
		}
		return nil, err // その他のエラー
	}

	return &user, nil
}
