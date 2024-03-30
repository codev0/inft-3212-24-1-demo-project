package model

import (
	"database/sql"
	"errors"
	"log"
)

type Restaurant struct {
	Id          string `json:"id"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Coordinates string `json:"coordinates"`
	Cousine     string `json:"cousine"`
}

type RestaurantModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m RestaurantModel) GetAll() ([]*Restaurant, error) {
	// TODO: implement this method
	return nil, errors.New("not implemented")
}

func (m RestaurantModel) Insert(menu *Restaurant) error {
	// TODO: implement this method
	return errors.New("not implemented")
}

func (m RestaurantModel) Get(id int) (*Restaurant, error) {
	// TODO: implement this method
	return nil, errors.New("not implemented")
}

func (m RestaurantModel) Update(menu *Restaurant) error {
	// TODO: implement this method
	return errors.New("not implemented")
}

func (m RestaurantModel) Delete(id int) error {
	// TODO: implement this method
	return errors.New("not implemented")
}
