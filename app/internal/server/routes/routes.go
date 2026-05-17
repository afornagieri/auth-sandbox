package routes

import (
	"net/http"

	"github.com/afornagieri/auth-sandbox.git/app/internal/controller"
)

func RegisterRoutes(mux *http.ServeMux, controllers *controller.TemplateController) {
	mux.HandleFunc("/", controllers.Login)
	mux.HandleFunc("/login", controllers.Login)
	mux.HandleFunc("/error", controllers.Error)
}
