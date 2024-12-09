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
			t.Fatalf("ユーザ追加に失敗: %v", err)
		}

		// IDがセットされているか確認
		if user.ID == 0 {
			t.Fatalf("ユーザIDがセットされていません, ID: %d", user.ID)
		}

		t.Logf("ユーザ追加に成功 ID: %d", user.ID)
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
			t.Fatalf("ユーザの編集に成功: %v", err)
		}

		t.Log("ユーザの編集に失敗")
	})

	// 3. ユーザ削除テスト
	t.Run("DeleteUser", func(t *testing.T) {
		user := &models.User{
			ID: 1, // 適宜、テスト環境に合わせたIDを設定
		}

		err := user.DeleteUser()
		if err != nil {
			t.Fatalf("ユーザの削除に失敗: %v", err)
		}

		t.Log("ユーザの削除に成功")
	})

	// 4. ユーザ一覧取得テスト
	t.Run("GetUsers", func(t *testing.T) {
		users, err := models.GetUsers()
		if err != nil {
			t.Fatalf("ユーザ一覧取得に失敗: %v", err)
		}

		if len(users) == 0 {
			t.Log("ユーザが見つかりませんでした")
		} else {
			t.Logf("ユーザ一覧取得に成功: %v", users)
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
			BID:            "TEST3",
			Address:        "TEST3",
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
			t.Fatalf("建物追加に失敗: %v", err)
		}

		// IDがセットされているか確認
		if building.ID == 0 {
			t.Fatalf("建物IDがセットされていません: %d", building.ID)
		}

		t.Logf("建物追加に成功 ID: %d", building.ID)
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
			t.Fatalf("建物編集に失敗: %v", err)
		}

		t.Log("建物編集に成功")
	})

	// 3. 建物削除テスト
	t.Run("DeleteBuilding", func(t *testing.T) {
		user := &models.Building{
			ID: 2, // 適宜、テスト環境に合わせたIDを設定
		}

		err := user.DeleteBuilding()
		if err != nil {
			t.Fatalf("建物削除に失敗: %v", err)
		}

		t.Log("建物削除に成功")
	})

	// 4. 建物覧取得テスト
	t.Run("GetBuildings", func(t *testing.T) {
		buildings, err := models.GetBuildings()
		if err != nil {
			t.Fatalf("建物一覧の取得に失敗: %v", err)
		}

		if len(buildings) == 0 {
			t.Log("建物が見つかりませんでした")
		} else {
			t.Logf("建物一覧に取得に成功: %v", buildings)
		}
	})
}

func TestFavoriteModel(t *testing.T) {
	// データベース接続を初期化
	db.InitDB()
	defer db.DB.Close()

	// 1. お気に入り追加テスト
	t.Run("AddFavorite", func(t *testing.T) {
		favorite := &models.Favorite{
			UID: "root",
			BID: "TEST3",
		}

		err := favorite.AddFavorite()
		if err != nil {
			t.Fatalf("お気に入り追加に失敗: %v", err)
		}

		// IDがセットされているか確認
		if favorite.ID == 0 {
			t.Fatalf("お気に入りIDがセットされていません: %d", favorite.ID)
		}

		t.Logf("お気に入り追加に成功 ID: %d", favorite.ID)
	})

	// 2. 建物削除テスト
	t.Run("DeleteFavorite", func(t *testing.T) {
		favorite := &models.Favorite{
			ID: 1, // 適宜、テスト環境に合わせたIDを設定
		}

		err := favorite.DeleteFavorite()
		if err != nil {
			t.Fatalf("お気に入り削除に失敗: %v", err)
		}

		t.Log("お気に入り削除に成功")
	})

	// 3. お気に入り覧取得テスト
	t.Run("GetFavorites", func(t *testing.T) {
		favorites, err := models.GetFavorites()
		if err != nil {
			t.Fatalf("お気に入り一覧の取得に失敗: %v", err)
		}
		if len(favorites) == 0 {
			t.Log("お気に入りはありませんでした")
		} else {
			t.Logf("お気に入り一覧の取得に成功: %v", favorites)
		}
	})

	t.Run("GetFavoritesByUser", func(t *testing.T) {
		favorites, err := models.GetFavoritesByUserID("root")
		if err != nil {
			t.Fatalf("ユーザIDによるお気に入り一覧の取得に失敗: %v", err)
		}
		if len(favorites) == 0 {
			t.Log("お気に入りはありませんでした")
		} else {
			t.Logf("ユーザIDによるお気に入り一覧の取得に成功: %v", favorites)
		}
	})

	t.Run("GetFavoritesByBuildingID", func(t *testing.T) {
		count, err := models.GetFavoritesByBuildingID("TEST")
		if err != nil {
			t.Fatalf("建物お気に入り数の取得に失敗: %v", err)
		}
		if count == 0 {
			t.Log("お気に入りは0でした")
		} else {
			t.Logf("建物お気に入り数の取得に成功: %v", count)
		}
	})
}
