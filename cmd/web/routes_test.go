package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi"
	"github.com/malikheena/heena2-main/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig
	mux := routes(&app)
	switch v := mux.(type) {
	case *chi.Mux:
	default:
		t.Error(fmt.Sprintf("tpye is not *chi.Mux,type is %T", v))
	}

}
