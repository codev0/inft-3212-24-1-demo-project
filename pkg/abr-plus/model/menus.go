package model

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/codev0/inft3212-6/pkg/abr-plus/validator"
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

func (m MenuModel) GetAll(title string, from, to int, filters Filters) ([]*Menu, Metadata, error) {

	// Retrieve all menu items from the database.
	query := fmt.Sprintf(
		`
		SELECT count(*) OVER(), id, created_at, updated_at, title, description, nutrition_value
		FROM menus
		WHERE (LOWER(title) = LOWER($1) OR $1 = '')
		AND (nutrition_value >= $2 OR $2 = 0)
		AND (nutrition_value <= $3 OR $3 = 0)
		ORDER BY %s %s, id ASC
		LIMIT $4 OFFSET $5
		`,
		filters.sortColumn(), filters.sortDirection())

	// Create a context with a 3-second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Organize our four placeholder parameter values in a slice.
	args := []interface{}{title, from, to, filters.limit(), filters.offset()}

	// log.Println(query, title, from, to, filters.limit(), filters.offset())
	// Use QueryContext to execute the query. This returns a sql.Rows result set containing
	// the result.
	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	// Importantly, defer a call to rows.Close() to ensure that the result set is closed
	// before GetAll returns.
	defer func() {
		if err := rows.Close(); err != nil {
			m.ErrorLog.Println(err)
		}
	}()

	// Declare a totalRecords variable
	totalRecords := 0

	var menus []*Menu
	for rows.Next() {
		var menu Menu
		err := rows.Scan(&totalRecords, &menu.Id, &menu.CreatedAt, &menu.UpdatedAt, &menu.Title, &menu.Description, &menu.NutritionValue)
		if err != nil {
			return nil, Metadata{}, err
		}

		// Add the Movie struct to the slice
		menus = append(menus, &menu)
	}

	// When the rows.Next() loop has finished, call rows.Err() to retrieve any error
	// that was encountered during the iteration.
	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	// Generate a Metadata struct, passing in the total record count and pagination parameters
	// from the client.
	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	// If everything went OK, then return the slice of the movies and metadata.
	return menus, metadata, nil
}

func (m MenuModel) Insert(menu *Menu) error {
	// Insert a new menu item into the database.
	query := `
		INSERT INTO menus (title, description, nutrition_value) 
		VALUES ($1, $2, $3) 
		RETURNING id, created_at, updated_at
		`
	args := []interface{}{menu.Title, menu.Description, menu.NutritionValue}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&menu.Id, &menu.CreatedAt, &menu.UpdatedAt)
}

func (m MenuModel) Get(id int) (*Menu, error) {
	// Return an error if the ID is less than 1.
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	// Retrieve a specific menu item based on its ID.
	query := `
		SELECT id, created_at, updated_at, title, description, nutrition_value
		FROM menus
		WHERE id = $1
		`
	var menu Menu
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&menu.Id, &menu.CreatedAt, &menu.UpdatedAt, &menu.Title, &menu.Description, &menu.NutritionValue)
	if err != nil {
		return nil, fmt.Errorf("cannot retrive menu with id: %v, %w", id, err)
	}
	return &menu, nil
}

func (m MenuModel) Update(menu *Menu) error {
	// Update a specific menu item in the database.
	query := `
		UPDATE menus
		SET title = $1, description = $2, nutrition_value = $3, updated_at = CURRENT_TIMESTAMP
		WHERE id = $4 AND updated_at = $5
		RETURNING updated_at
		`
	args := []interface{}{menu.Title, menu.Description, menu.NutritionValue, menu.Id, menu.UpdatedAt}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&menu.UpdatedAt)
}

func (m MenuModel) Delete(id int) error {
	// Return an error if the ID is less than 1.
	if id < 1 {
		return ErrRecordNotFound
	}

	// Delete a specific menu item from the database.
	query := `
		DELETE FROM menus
		WHERE id = $1
		`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, id)
	return err
}

func ValidateMenu(v *validator.Validator, menu *Menu) {
	// Check if the title field is empty.
	v.Check(menu.Title != "", "title", "must be provided")
	// Check if the title field is not more than 100 characters.
	v.Check(len(menu.Title) <= 100, "title", "must not be more than 100 bytes long")
	// Check if the description field is not more than 1000 characters.
	v.Check(len(menu.Description) <= 1000, "description", "must not be more than 1000 bytes long")
	// Check if the nutrition value is not more than 10000.
	v.Check(menu.NutritionValue <= 10000, "nutritionValue", "must not be more than 10000")
}
