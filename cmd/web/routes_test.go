package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/samiulru/bookings/internal/config"
	"testing"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//do noting, test passed
	default:
		t.Errorf("Type is not *chi.Mux, but is %s", v)
	}

}
