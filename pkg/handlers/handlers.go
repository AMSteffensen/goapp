package handlers

import (
	"net"
	"net/http"

	"github.com/AMSteffensen/goapp/pkg/config"
	"github.com/AMSteffensen/goapp/pkg/models"
	"github.com/AMSteffensen/goapp/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the homepage handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		remoteIP = r.RemoteAddr
	}
	ip := net.ParseIP(remoteIP)
	if ip != nil && ip.IsLoopback() && ip.To4() == nil {
		remoteIP = "127.0.0.1"
	}
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	stringMap["test"] = "Hello, again."

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
