package app

import (
	"log"
	"net/http"

	"github.com/OpenTabletopHub/TabletopWiki/internal/wiki"
)

type AppServer struct {
	addr string
}

func NewAppServer(addr string) *AppServer {
	return &AppServer{addr: addr}
}

func (s *AppServer) Run() error {

	// Register the wiki handler
	wikiHandler := wiki.NewHandler()
	wikiRouter := wikiHandler.Register()

	router := http.NewServeMux()
	router.Handle("/wiki/", http.StripPrefix("/wiki", wikiRouter))

	log.Println("Starting server on port", s.addr)
	return http.ListenAndServe(s.addr, router)
}
