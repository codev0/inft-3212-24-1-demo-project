package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/codev0/inft3212-6/pkg/abr-plus/model"
)

func (app *application) createMenuHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title          string `json:"title"`
		Description    string `json:"description"`
		NutritionValue uint   `json:"nutritionValue"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		log.Println(err)
		app.errorResponse(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}

	menu := &model.Menu{
		Title:          input.Title,
		Description:    input.Description,
		NutritionValue: input.NutritionValue,
	}

	err = app.models.Menus.Insert(menu)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusCreated, envelope{"menu": menu}, nil)
}

func (app *application) getMenusList(w http.ResponseWriter, r *http.Request) {
	// TODO: implement filtering
	menus, err := app.models.Menus.GetAll()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"menus": menus}, nil)
}

func (app *application) getMenuHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	menu, err := app.models.Menus.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"menu": menu}, nil)
}

func (app *application) updateMenuHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	menu, err := app.models.Menus.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Title          *string `json:"title"`
		Description    *string `json:"description"`
		NutritionValue *uint   `json:"nutritionValue"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
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
		switch {
		case errors.Is(err, model.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"menu": menu}, nil)
}

func (app *application) deleteMenuHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Menus.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"message": "success"}, nil)
}
