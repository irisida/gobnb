package handlers

import (
	"fmt"
	"net/http"

	"github.com/irisida/gobnb/pkg/config"
	"github.com/irisida/gobnb/pkg/models"
	"github.com/irisida/gobnb/pkg/render"
)

// Repository type definition
type Repository struct {
	App *config.AppConfig
}

// Repo instance used by the handlers
var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler.
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is handler function for page to be shown at the /about route.
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Booking(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "booking.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Americana(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "americana.page.tmpl", &models.TemplateData{})
}

func (m *Repository) SwingingLondon(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "swinging-london.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Parisian(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "parisian.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "availability.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// post request handlers

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	//render.RenderTemplate(w, "availability.page.tmpl", &models.TemplateData{})
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date: %s  end date: %s", start, end)))

}
