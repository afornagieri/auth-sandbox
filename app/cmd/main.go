package main

import (
	"log"
	"net/http"

	"github.com/afornagieri/auth-sandbox.git/app/internal/controller"
	"github.com/afornagieri/auth-sandbox.git/app/internal/server"
	"github.com/afornagieri/auth-sandbox.git/app/internal/server/config"
	"github.com/afornagieri/auth-sandbox.git/app/internal/server/routes"
)

func main() {
	cfg := config.LoadConfig()
	mux := http.NewServeMux()
	templateController, err := controller.NewTemplateController()
	if err != nil {
		log.Fatal("Erro ao carregar templates: ", err)
	}

	routes.RegisterRoutes(mux, templateController)
	svr := server.NewServer(cfg.Port, mux)
	log.Printf("Servidor iniciando na porta %s...\n", cfg.Port)
	log.Fatal(svr.Start())
}
