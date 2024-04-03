package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available", "system_info": map[string]string{
			"environment": app.config.env,
			"version":     version},
	}
	// Add a 10 second delay to demonstrate the server returning a response after shutting down
	// time.Sleep(10 * time.Second)
	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
