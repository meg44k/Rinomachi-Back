package models

import "RenomachiBack/db"

type User struct {
	ID       int    `json:"id"`
	UID      string `json:"uid"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// ユーザを追加
func (user *User) AddUser() error {
	query := "INSERT INTO users (user_id, user_name, password, email) VALUES (?, ?, ?, ?)"
	result, err := db.DB.Exec(query, user.UID, user.Name, user.Password, user.Email)
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
func (user *User) UpdateUser() error {
	query := "UPDATE users SET user_name = ?, password = ?, email = ? WHERE id = ?"
	_, err := db.DB.Exec(query, user.Name, user.Password, user.Email, user.ID)
	return err
}

// ユーザを削除
func (user *User) DeleteUser() error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := db.DB.Exec(query, user.ID)
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
		if err := rows.Scan(&user.ID, &user.UID, &user.Name, &user.Password, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
