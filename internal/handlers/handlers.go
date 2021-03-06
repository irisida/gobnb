package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/irisida/gobnb/internal/config"
	"github.com/irisida/gobnb/internal/forms"
	"github.com/irisida/gobnb/internal/models"
	"github.com/irisida/gobnb/internal/render"
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
	var emptyBooking models.Booking
	data := make(map[string]interface{})
	data["booking"] = emptyBooking

	render.RenderTemplate(w, r, "booking.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostBooking handles the posting of a booking form
func (m *Repository) PostBooking(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	booking := models.Booking{
		Firstname: r.Form.Get("firstname"),
		Lastname:  r.Form.Get("lastname"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	form.Required("firstname", "lastname", "email")
	form.MinLength("firstname", 3, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["booking"] = booking

		render.RenderTemplate(w, r, "booking.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})

		return
	}
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
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
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

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// json responses
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}
