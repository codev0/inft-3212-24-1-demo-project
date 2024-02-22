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

var restaurants = []Restaurant{
	{
		Id:      "1",
		Title:   "The Cinnamon Club",
		Address: "The Old Westminster Library, 30-32 Great Smith St, Westminster, London SW1P 3BU",
		Cousine: "Indian",
	},
	{
		Id:      "2",
		Title:   "Nobu",
		Address: "19 Old Park Ln, Mayfair, London W1K 1LB",
		Cousine: "Japanese",
	},
	{
		Id:      "3",
		Title:   "Gordon Ramsay",
		Address: "68 Royal Hospital Rd, Chelsea, London SW3 4HP",
		Cousine: "French",
	},
	{
		Id:      "4",
		Title:   "The Ledbury",
		Address: "127 Ledbury Rd, Notting Hill, London W11 2AQ",
		Cousine: "Modern European",
	},
	{
		Id:      "5",
		Title:   "Wahaca",
		Address: "66 Chandos Pl, Covent Garden, London WC2N 4HG",
		Cousine: "Mexican",
	},
}

func GetRestaurants() []Restaurant {
	return restaurants
}

func GetRestaurant(id string) (*Restaurant, error) {
	for _, r := range restaurants {
		if r.Id == id {
			return &r, nil
		}
	}
	return nil, errors.New("Restaurant not found")
}
