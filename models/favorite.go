package models

import "RenomachiBack/db"

type Favorite struct {
	ID  int
	UID string
	BID string
}

// お気に入りを追加
func (favorite *Favorite) AddFavorite() error {
	query := "INSERT INTO favorites (user_id, building_id) VALUES (?, ?)"
	result, err := db.DB.Exec(query, favorite.UID, favorite.BID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	favorite.ID = int(id)
	return nil
}

// お気に入りを削除
func (favorite *Favorite) DeleteFavorite() error {
	query := "DELETE FROM favorites WHERE id = ?"
	_, err := db.DB.Exec(query, favorite.ID)
	return err
}

// お気に入り一覧を表示
func GetFavorites() ([]Favorite, error) {
	query := "SELECT * FROM favorites"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var favorites []Favorite
	for rows.Next() {
		var favorite Favorite
		if err := rows.Scan(&favorite.ID, &favorite.UID, &favorite.BID); err != nil {
			return nil, err
		}
		favorites = append(favorites, favorite)
	}
	return favorites, nil
}

// UIDのお気に入り一覧を表示
func GetFavoritesByUserID(user_id string) ([]Favorite, error) {
	query := "SELECT * FROM favorites WHERE user_id = ?"
	rows, err := db.DB.Query(query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var favorites []Favorite
	for rows.Next() {
		var favorite Favorite
		if err := rows.Scan(&favorite.ID, &favorite.UID, &favorite.BID); err != nil {
			return nil, err
		}
		favorites = append(favorites, favorite)
	}
	return favorites, nil
}

// BIDのお気に入り数を表示
func GetFavoritesByBuildingID(building_id string) (int, error) {
	query := "SELECT COUNT(*) AS record_count FROM favorites WHERE building_id = ?;"
	var count int
	err := db.DB.QueryRow(query, building_id).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}
