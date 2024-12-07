package api

import (
	// "plex-qbt-controller/handlers"
	"net/http"
)

func NewRouter(h Handler) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", h.PlexEvents)

}
