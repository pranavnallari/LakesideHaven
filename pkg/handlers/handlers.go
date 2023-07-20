package handlers

import (
	"net/http"

	"github.com/pranavnallari/LakesideHaven/pkg/config"
	"github.com/pranavnallari/LakesideHaven/pkg/models"
	"github.com/pranavnallari/LakesideHaven/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	// send data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{})
}
