package main

import (
	"plex-qbt-controller/api"
	"plex-qbt-controller/app"
	"plex-qbt-controller/plex"
	"plex-qbt-controller/qbittorrent"
)

func main() {
	qbp := qbittorrent.NewBittorrentCreateParams{
		Username:             "admin",
		Password:             "admin123",
		Host:                 "localhost",
		PlexLibrariesToPause: []string{"All Movies"},
	}
	plexParams := plex.PlexClientCreate{
		Host:  "localhost",
		Token: "swfkQ5sXuS-w-meqHksV",
		Port:  "32400",
	}
	plex := plex.New(plexParams)
	_, err := plex.GetLibrarySectionById(1)
	if err != nil {
		panic(err)
	}

	qbc, err := qbittorrent.New(qbp)
	if err != nil {
		panic(err)
	}
	app := app.App{}
	handler := api.NewHandler(&app)
	api.SetHandler(handler)

	app.AddQBittorrentClient(qbc)
	// plex.GetLibraryByID(1)
	// plex.GetLibraryByID(1)
	app.SetPlexClient(plex)
	// torrents := qbc.GetAllTorrents(qbittorrent.ALL)
	// handler := api.Handler.NewHandler(&app)
	// handlers.SetHandler(handler)

	// qbc.ResumeTorrents(torrents...)
	// fmt.Println("bob")
	// plex.GetAllLibraries()
	app.Run()

}
