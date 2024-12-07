package main

import (
	"fmt"
	"os"
	"os/signal"
	"plex-qbt-controller/api"
	"plex-qbt-controller/app"
	"plex-qbt-controller/plex"
	"plex-qbt-controller/qbittorrent"
	"sync"
	"syscall"
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
	app := app.App{
		Quit: make(chan os.Signal, 1),
		Wait: sync.WaitGroup{},
	}
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
	// time.Sleep(time.Second * 10)
	// appQuit <- true
	signal.Notify(app.Quit, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Waiting until ctrl+c")
	app.Wait.Wait()

}
