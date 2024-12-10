package models

import (
	"RenomachiBack/db"
	"time"
)

type History struct {
	ID   int       `json:"id"`
	UID  string    `json:"uid"`
	BID  string    `json:"bid"`
	Time time.Time `json:"time"`
}

// 履歴を追加
func (history *History) AddHistory() error {
	query := "INSERT INTO histories (user_id, building_id) VALUES (?, ?)"
	result, err := db.DB.Exec(query, history.UID, history.BID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	history.ID = int(id)
	return nil
}

// 履歴を削除
func DeleteHistory(user_id string, building_id string) error {
	query := "DELETE FROM histories WHERE user_id = ? AND building_id"
	_, err := db.DB.Exec(query, user_id, building_id)
	return err
}

// 履歴一覧を表示
func GetHistories(user_id string) ([]History, error) {
	query := "SELECT * FROM histories WHERE user_id = ?"
	rows, err := db.DB.Query(query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var histories []History
	for rows.Next() {
		var history History
		if err := rows.Scan(&history.ID, &history.UID, &history.BID, &history.Time); err != nil {
			return nil, err
		}
		histories = append(histories, history)
	}
	return histories, nil
}

// UIDのお気に入り一覧を表示
func GetHistoriesByUserID(user_id string) ([]History, error) {
	query := "SELECT * FROM histories WHERE user_id = ?"
	rows, err := db.DB.Query(query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var histories []History
	for rows.Next() {
		var history History
		if err := rows.Scan(&history.ID, &history.UID, &history.BID, &history.Time); err != nil {
			return nil, err
		}
		histories = append(histories, history)
	}
	return histories, nil
}
