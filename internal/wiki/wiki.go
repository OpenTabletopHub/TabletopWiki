package wiki

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Register() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("GET /", getfrontPage)

	r.HandleFunc("GET /p/{id}/", getPage)

	return r
}

func getfrontPage(w http.ResponseWriter, r *http.Request) {
	// Render the front page
	fmt.Fprintln(w, "Welcome to the OTT Wiki!")
}
func getPage(w http.ResponseWriter, r *http.Request) {
	// Extract the page ID from the URL
	pageID := r.URL.Path[len("/wiki/"):] // Assuming the URL is like /wiki/{id}/

	// Render the requested page
	fmt.Fprintf(w, "You are viewing page: %s", pageID)
}
