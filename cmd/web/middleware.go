package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf adds CSRF protection on all post requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteDefaultMode,
	})
	return csrfHandler
}

// SesionLoad loads and saves the ssion on every resest
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
