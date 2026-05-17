package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

type TemplateController struct {
	templates *template.Template
}

func NewTemplateController() (*TemplateController, error) {
	// templatesPath := "../templates/*.html"
	templatesPath := "app/templates/*.html"

	fmt.Printf("Buscando templates em: %s\n", templatesPath)

	temp, err := template.ParseGlob(templatesPath)
	if err != nil {
		fmt.Printf("Erro ao parsear templates: %v\n", err)
		return nil, err
	}

	fmt.Println("Templates carregados com sucesso!")

	return &TemplateController{
		templates: temp,
	}, nil
}

func (temp *TemplateController) Login(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	data := map[string]string{
		"Title":   "Login",
		"Message": "Login",
	}

	err := temp.templates.ExecuteTemplate(res, "login.html", data)
	if err != nil {
		fmt.Printf("Erro ao executar template: %v\n", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (temp *TemplateController) Error(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	data := map[string]string{
		"Title":   "Error",
		"Message": "Error",
	}

	err := temp.templates.ExecuteTemplate(res, "error.html", data)
	if err != nil {
		fmt.Printf("Erro ao executar template: %v\n", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
