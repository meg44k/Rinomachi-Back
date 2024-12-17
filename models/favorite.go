package models

import (
	"RenomachiBack/db"
	"time"
)

type Favorite struct {
	ID        int       `json:"id"`
	UID       string    `json:"uid"`
	BID       string    `json:"bid"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
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

	fav_count, _ := GetFavoritesByBuildingID(favorite.BID)
	query = "UPDATE buildings SET favorites = ? WHERE building_id = ?"
	_, err = db.DB.Exec(query, fav_count, favorite.BID)
	if err != nil {
		return err
	}

	return nil
}

// お気に入りを削除
func DeleteFavorite(user_id string, building_id string) error {
	query := "DELETE FROM favorites WHERE user_id = ? AND building_id = ?"
	_, err := db.DB.Exec(query, user_id, building_id)
	return err
}

// UIDのお気に入り一覧を表示
func GetFavorites(user_id string) ([]Favorite, error) {
	query := "SELECT * FROM favorites WHERE user_id = ?"
	rows, err := db.DB.Query(query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var favorites []Favorite
	for rows.Next() {
		var favorite Favorite
		if err := rows.Scan(&favorite.ID, &favorite.UID, &favorite.BID, &favorite.CreatedAt, &favorite.UpdatedAt); err != nil {
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
