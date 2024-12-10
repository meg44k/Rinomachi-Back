package models

import "RenomachiBack/db"

type Building struct {
	ID             int     `json:"id"`
	BID            string  `json:"bid"`
	Address        string  `json:"address"`
	Structure      string  `json:"structure"`
	Floors         int     `json:"floors"`
	Age            int     `json:"age"`
	Area           float64 `json:"area"`
	Contract       string  `json:"contract"`
	Description    string  `json:"discription"`
	IsAvailable    bool    `json:"isAvailable"`
	Price          int     `json:"price"`
	Favorites      int     `json:"favorites"`
	Transportation string  `json:"transportation"`
}

// ユーザを追加
func (building *Building) AddBuilding() error {
	query := "INSERT INTO buildings (building_id, address, structure, floors, age, area, contract, description, is_available, price, favorites, transportation) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := db.DB.Exec(query,
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

// ユーザを編集
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

// ユーザを削除
func (building *Building) DeleteBuilding() error {
	query := "DELETE FROM buildings WHERE id = ?"
	_, err := db.DB.Exec(query, building.ID)
	return err
}

// ユーザ一覧を表示
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
			&building.Transportation); err != nil {
			return nil, err
		}
		buildings = append(buildings, building)
	}
	return buildings, nil
}
