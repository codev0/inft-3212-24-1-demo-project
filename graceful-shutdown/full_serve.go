package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func fullServe() error {
	port := 8083
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, from Full Sever!")
	})

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Create a shutdownError channel. We will use this to receive any errors returned
	// by the graceful Shutdown() function.
	shutdownError := make(chan error)
	go func() {
		// Intercept the signals, as before.
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit
		// Update the log entry to say "shutting down server" instead of "caught signal".
		fmt.Println("shutting down server", map[string]string{"signal": s.String()})
		// Create a context with a 5-second timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// Call Shutdown() on our server, passing in the context we just made.
		// Shutdown() will return nil if the graceful shutdown was successful, or an
		// error (which may happen because of a problem closing the listeners, or
		// because the shutdown didn't complete before the 5-second context deadline is // hit). We relay this return value to the shutdownError channel.
		shutdownError <- srv.Shutdown(ctx)
	}()

	fmt.Println("starting server", map[string]string{"addr": srv.Addr, "port": fmt.Sprint(port)})

	// Calling Shutdown() on our server will cause ListenAndServe() to immediately
	// return a http.ErrServerClosed error. So if we see this error, it is actually a
	// good thing and an indication that the graceful shutdown has started. So we check
	// specifically for this, only returning the error if it is NOT http.ErrServerClosed.
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	// Otherwise, we wait to receive the return value from Shutdown() on the
	// shutdownError channel. If return value is an error, we know that there was a
	// problem with the graceful shutdown and we return the error.
	err = <-shutdownError
	if err != nil {
		return err
	}
	// At this point we know that the graceful shutdown completed successfully and we
	// log a "stopped server" message.
	fmt.Println("starting server", map[string]string{"addr": srv.Addr, "port": fmt.Sprint(port)})
	return nil
}
