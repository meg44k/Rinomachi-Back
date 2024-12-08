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

func TestBuildingModel(t *testing.T) {
	// データベース接続を初期化
	db.InitDB()
	defer db.DB.Close()

	// 1. 建物追加テスト
	t.Run("AddBuilding", func(t *testing.T) {
		building := &models.Building{
			BID:            "TEST",
			Address:        "TEST",
			Structure:      "TEST",
			Floors:         2,
			Age:            2,
			Area:           2.1213,
			Contract:       "TEST",
			Description:    "TEST",
			IsAvailable:    true,
			Price:          1230,
			Favorites:      1,
			Transportation: "TEST",
		}

		err := building.AddBuilding()
		if err != nil {
			t.Fatalf("Failed to add building: %v", err)
		}

		// IDがセットされているか確認
		if building.ID == 0 {
			t.Fatalf("Expected user ID to be set, got: %d", building.ID)
		}

		t.Logf("User added successfully with ID: %d", building.ID)
	})

	// 2. 建物更新テスト
	t.Run("UpdateBuilding", func(t *testing.T) {
		user := &models.Building{
			ID:             1,
			BID:            "TESTED",
			Address:        "TESTED",
			Structure:      "TESTED",
			Floors:         21,
			Age:            21,
			Area:           2.12133123,
			Contract:       "TESTED",
			Description:    "TESTED",
			IsAvailable:    false,
			Price:          123000,
			Favorites:      100,
			Transportation: "TESTED",
		}

		err := user.UpdateBuilding()
		if err != nil {
			t.Fatalf("Failed to update building: %v", err)
		}

		t.Log("Building updated successfully")
	})

	// 3. 建物削除テスト
	t.Run("DeleteBuilding", func(t *testing.T) {
		user := &models.Building{
			ID: 2, // 適宜、テスト環境に合わせたIDを設定
		}

		err := user.DeleteBuilding()
		if err != nil {
			t.Fatalf("Failed to delete building: %v", err)
		}

		t.Log("Building deleted successfully")
	})

	// 4. 建物覧取得テスト
	t.Run("GetBuildings", func(t *testing.T) {
		buildings, err := models.GetBuildings()
		if err != nil {
			t.Fatalf("Failed to get buildings: %v", err)
		}

		if len(buildings) == 0 {
			t.Log("No buildings found")
		} else {
			t.Logf("Buildings retrieved successfully: %v", buildings)
		}
	})
}
