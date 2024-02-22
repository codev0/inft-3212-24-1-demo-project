package model

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Menu struct {
	Id             string `json:"id"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	NutritionValue uint   `json:"nutritionValue"`
}

type MenuModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m MenuModel) Insert(menu *Menu) error {
	// Insert a new menu item into the database.
	query := `
		INSERT INTO menu (title, description, nutrition_value) 
		VALUES ($1, $2, $3) 
		RETURNING id, created_at, updated_at
		`
	args := []interface{}{menu.Title, menu.Description, menu.NutritionValue}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&menu.Id, &menu.CreatedAt, &menu.UpdatedAt)
}

func (m MenuModel) Get(id int) (*Menu, error) {
	// Retrieve a specific menu item based on its ID.
	query := `
		SELECT id, created_at, updated_at, title, description, nutrition_value
		FROM menu
		WHERE id = $1
		`
	var menu Menu
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&menu.Id, &menu.CreatedAt, &menu.UpdatedAt, &menu.Title, &menu.Description, &menu.NutritionValue)
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

func (m MenuModel) Update(menu *Menu) error {
	// Update a specific menu item in the database.
	query := `
		UPDATE menu
		SET title = $1, description = $2, nutrition_value = $3
		WHERE id = $4
		RETURNING updated_at
		`
	args := []interface{}{menu.Title, menu.Description, menu.NutritionValue, menu.Id}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&menu.UpdatedAt)
}

func (m MenuModel) Delete(id int) error {
	// Delete a specific menu item from the database.
	query := `
		DELETE FROM menu
		WHERE id = $1
		`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, id)
	return err
}
