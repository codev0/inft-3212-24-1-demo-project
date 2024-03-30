package main

import (
	"errors"
	"net/http"
	"strings"

	"github.com/codev0/inft3212-6/pkg/abr-plus/model"
	"github.com/codev0/inft3212-6/pkg/abr-plus/validator"
)

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add the "Vary: Authorization" header to the response. This indicates to any caches
		// that the response may vary based on the value of the Authorization header in the request.
		w.Header().Set("Vary", "Authorization")

		// Retrieve the value of the Authorization header from teh request. This will return the
		// empty string "" if there is no such header found.
		authorizationHeader := r.Header.Get("Authorization")

		// If there is no Authorization header found, use the contextSetUser() helper to add
		// an AnonymousUser to the request context. Then we call the next handler in the chain
		// and return without executing any of the code below.
		if authorizationHeader == "" {
			r = app.contextSetUser(r, model.AnonymousUser)
			next.ServeHTTP(w, r)
			return
		}

		// Otherwise, we expect the value of the Authorization header to be in the format
		// "Bearer <token>". We try to split this into its constituent parts, and if the header
		// isn't in the expected format we return a 401 Unauthorized response using the
		// invalidAuthenticationTokenResponse helper.
		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			app.invalidAuthenticationTokenResponse(w, r)
			return
		}

		// Extract the actual authentication toekn from the header parts
		token := headerParts[1]

		// Validate the token to make sure it is in a sensible format.
		v := validator.New()

		// If the token isn't valid, use the invalidAuthenticationtokenResponse
		// helper to send a response, rather than the failedValidatedResponse helper.
		if model.ValidateTokenPlaintext(v, token); !v.Valid() {
			app.invalidAuthenticationTokenResponse(w, r)
			return
		}

		// Retrieve the details of the user associated with the authentication token.
		// call invalidAuthenticationTokenResponse if no matching record was found.
		user, err := app.models.Users.GetForToken(model.ScopeAuthentication, token)
		if err != nil {
			switch {
			case errors.Is(err, model.ErrRecordNotFound):
				app.invalidAuthenticationTokenResponse(w, r)
			default:
				app.serverErrorResponse(w, r, err)
			}
			return
		}

		// Call the contextSetUser healer to add the user information to the request context.
		r = app.contextSetUser(r, user)

		// Call next handler in chain
		next.ServeHTTP(w, r)
	})
}

// requireAuthenticatedUser checks that the user is not anonymous (i.e., they are authenticated).
func (app *application) requireAuthenticatedUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Use the contextGetUser helper to retrieve the user information from the request context.
		user := app.contextGetUser(r)

		// If the user is anonymous, then call authenticationRequiredResponse to inform the client
		// that they should be authenticated before trying again.
		if user.IsAnonymous() {
			app.authenticationRequiredResponse(w, r)
			return
		}

		next.ServeHTTP(w, r)
	}
}

// requiredActivatedUser checks that the user is both authenticated and activated.
func (app *application) requireActivatedUser(next http.HandlerFunc) http.HandlerFunc {
	// Rather than returning this http.HandlerFunc we assign it to the variable fn.
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := app.contextGetUser(r)

		// Check that a user is activated
		if !user.Activated {
			app.inactiveAccountResponse(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})

	// Wrap fn with the requireAuthenticatedUser middleware before returning it.
	return app.requireAuthenticatedUser(fn)
}

func (app *application) requirePermissions(code string, next http.HandlerFunc) http.HandlerFunc {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the user from the request context.
		user := app.contextGetUser(r)

		// Get the slice of permission for the user
		permissions, err := app.models.Permissions.GetAllForUser(user.ID)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		// Check if the slice includes the required permission. If it doesn't, then return a 403
		// Forbidden response.
		if !permissions.Include(code) {
			app.notPermittedResponse(w, r)
			return
		}

		// Otherwise, they have the required permission so we call the next handler in the chain.
		next.ServeHTTP(w, r)
	})

	// Wrap this with the requireActivatedUser middleware before returning
	return app.requireActivatedUser(fn)
}
