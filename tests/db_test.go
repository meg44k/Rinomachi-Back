package models_test

import (
	"RenomachiBack/db"
	"RenomachiBack/models"
	"testing"
)

func TestUserModel(t *testing.T) {
	// データベース接続を初期化
	db.InitDB()
	defer db.DB.Close()

	// 1. ユーザ追加テスト
	t.Run("AddUser", func(t *testing.T) {
		user := &models.User{
			UID:      "root",
			Name:     "root",
			Password: "root",
			Email:    "root@example.com", // 仮のメールアドレス
		}

		err := user.AddUser()
		if err != nil {
			t.Fatalf("Failed to add user: %v", err)
		}

		// IDがセットされているか確認
		if user.ID == 0 {
			t.Fatalf("Expected user ID to be set, got: %d", user.ID)
		}

		t.Logf("User added successfully with ID: %d", user.ID)
	})

	// 2. ユーザ更新テスト
	t.Run("UpdateUser", func(t *testing.T) {
		user := &models.User{
			ID:       1, // 適宜、テスト環境に合わせたIDを設定
			Name:     "updated_root",
			Password: "updated_password",
			Email:    "updated_email@example.com",
		}

		err := user.UpdateUser()
		if err != nil {
			t.Fatalf("Failed to update user: %v", err)
		}

		t.Log("User updated successfully")
	})

	// 3. ユーザ削除テスト
	t.Run("DeleteUser", func(t *testing.T) {
		user := &models.User{
			ID: 1, // 適宜、テスト環境に合わせたIDを設定
		}

		err := user.DeleteUser()
		if err != nil {
			t.Fatalf("Failed to delete user: %v", err)
		}

		t.Log("User deleted successfully")
	})

	// 4. ユーザ一覧取得テスト
	t.Run("GetUsers", func(t *testing.T) {
		users, err := models.GetUsers()
		if err != nil {
			t.Fatalf("Failed to get users: %v", err)
		}

		if len(users) == 0 {
			t.Log("No users found")
		} else {
			t.Logf("Users retrieved successfully: %v", users)
		}
	})
}
