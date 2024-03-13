package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/codev0/inft3212-6/pkg/abr-plus/model"
	"github.com/gorilla/mux"
)

func (app *application) respondWithError(w http.ResponseWriter, code int, message string) {
	app.respondWithJSON(w, code, map[string]string{"error": message})
}

func (app *application) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (app *application) createMenuHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title          string `json:"title"`
		Description    string `json:"description"`
		NutritionValue uint   `json:"nutritionValue"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		log.Println(err)
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	menu := &model.Menu{
		Title:          input.Title,
		Description:    input.Description,
		NutritionValue: input.NutritionValue,
	}

	err = app.models.Menus.Insert(menu)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusCreated, menu)
}

func (app *application) getMenuHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["menuId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid menu ID")
		return
	}

	menu, err := app.models.Menus.Get(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Menu with ID %d not found\n", id)
		}
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	app.respondWithJSON(w, http.StatusOK, menu)
}

func (app *application) updateMenuHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["menuId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid menu ID")
		return
	}

	menu, err := app.models.Menus.Get(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	var input struct {
		Title          *string `json:"title"`
		Description    *string `json:"description"`
		NutritionValue *uint   `json:"nutritionValue"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if input.Title != nil {
		menu.Title = *input.Title
	}

	if input.Description != nil {
		menu.Description = *input.Description
	}

	if input.NutritionValue != nil {
		menu.NutritionValue = *input.NutritionValue
	}

	err = app.models.Menus.Update(menu)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, menu)
}

func (app *application) deleteMenuHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["menuId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid menu ID")
		return
	}

	err = app.models.Menus.Delete(id)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		return err
	}

	return nil
}
