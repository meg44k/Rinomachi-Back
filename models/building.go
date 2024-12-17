package models

import (
	"RenomachiBack/db"
	"RenomachiBack/utils"
	"database/sql"
	"time"
)

type Building struct {
	ID             int       `json:"id"`
	BID            string    `json:"bid"`
	Address        string    `json:"address"`
	Structure      string    `json:"structure"`
	Floors         int       `json:"floors"`
	Age            int       `json:"age"`
	Area           float64   `json:"area"`
	Contract       string    `json:"contract"`
	Description    string    `json:"discription"`
	IsAvailable    bool      `json:"isAvailable"`
	Price          int       `json:"price"`
	Favorites      int       `json:"favorites"`
	Transportation string    `json:"transportation"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// 建物を追加
func (building *Building) AddBuilding() error {
	BID := utils.GenerateBuildingID()

	query := "INSERT INTO buildings (building_id, address, structure, floors, age, area, contract, description, is_available, price, favorites, transportation) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := db.DB.Exec(query,
		BID,
		building.Address,
		building.Structure,
		building.Floors,
		building.Age,
		building.Area,
		building.Contract,
		building.Description,
		building.IsAvailable,
		building.Price,
		building.Favorites,
		building.Transportation)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	building.ID = int(id)
	return nil
}

// 建物を編集
func (building *Building) UpdateBuilding() error {
	query := "UPDATE buildings SET building_id = ?, address = ?, structure = ?, floors = ?, age = ?, area = ?, contract = ?, description = ?, is_available = ?, price = ?, favorites = ?, transportation = ? WHERE id = ?"
	_, err := db.DB.Exec(query,
		building.BID,
		building.Address,
		building.Structure,
		building.Floors,
		building.Age,
		building.Area,
		building.Contract,
		building.Description,
		building.IsAvailable,
		building.Price,
		building.Favorites,
		building.Transportation,
		building.ID)
	return err
}

// 建物を削除
func DeleteBuilding(building_id string) error {
	query := "DELETE FROM buildings WHERE building_id = ?"
	_, err := db.DB.Exec(query, building_id)
	return err
}

// 建物一覧を表示
func GetBuildings() ([]Building, error) {
	query := "SELECT * FROM buildings"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var buildings []Building
	for rows.Next() {
		var building Building
		if err := rows.Scan(
			&building.ID,
			&building.BID,
			&building.Address,
			&building.Structure,
			&building.Floors,
			&building.Age,
			&building.Area,
			&building.Contract,
			&building.Description,
			&building.IsAvailable,
			&building.Price,
			&building.Favorites,
			&building.Transportation,
			&building.CreatedAt,
			&building.UpdatedAt); err != nil {
			return nil, err
		}
		buildings = append(buildings, building)
	}
	return buildings, nil
}

func GetBuilding(building_id string) (*Building, error) {
	query := "SELECT * FROM buildings WHERE building_id = ?"
	row := db.DB.QueryRow(query, building_id)

	var building Building
	err := row.Scan(
		&building.ID,
		&building.BID,
		&building.Address,
		&building.Structure,
		&building.Floors,
		&building.Age,
		&building.Area,
		&building.Contract,
		&building.Description,
		&building.IsAvailable,
		&building.Price,
		&building.Favorites,
		&building.Transportation,
		&building.CreatedAt,
		&building.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // 該当する建物がない場合
		}
		return nil, err // その他のエラー
	}

	return &building, nil
}
