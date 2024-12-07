package api

import (
	"net/http"
	app "plex-qbt-controller/app"
	"plex-qbt-controller/plex"
)

type Handler struct {
	App *app.App
}

var handler *Handler

func NewHandler(a *app.App) *Handler {
	return &Handler{
		App: a,
	}
}

func SetHandler(h *Handler) {
	handler = h
}

func (h *Handler) PlexEvents(w http.ResponseWriter, r *http.Request) {
	data, err := plex.ParsePlexEventBodyFromRequest(r)
	if err != nil {
		panic(err)
	}
	event, ok := plex.PlexEventMap[data.Event]
	library := plex.PlexLibrary{
		Name: data.Metadata.LibrarySectionTitle,
	}
	for _, qbt := range h.App.QbittorrentClients {
		for _, libName := range qbt.PlexLibrariesToPause {
			if libName == library.Name {
				switch event {
				case plex.PAUSED:
					streaming, err := h.App.Plex.GetLibrariesBeingStreamed()
					if err != nil {
						panic(err)
					}
					if len(streaming) <= 0 {
						qbt.ResumeAllTorrents()
					}
				case plex.RESUMED:
					qbt.PauseAllTorrents()
				default:
					panic("Could not find event")
				}
			}
		}
		if !ok {
			panic("Event is invalid")
		}
	}
}
