package service

import (
	"encoding/json"
	"os"
)

type Menu struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

func GetMenu() ([]Menu, error) {

	file, err := os.ReadFile("data/menu.json")
	if err != nil {
		return nil, err
	}

	var menus []Menu

	err = json.Unmarshal(file, &menus)
	if err != nil {
		return nil, err
	}

	return menus, nil
}
