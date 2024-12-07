package app

import (
	"fmt"
	"plex-qbt-controller/plex"
	"plex-qbt-controller/qbittorrent"
	"time"
)

type App struct {
	QbittorrentClients []*qbittorrent.QBittorrentClient
	Plex               *plex.PlexClient
	Quit               chan bool
	AppChan            chan bool
}

func (a *App) AddQBittorrentClient(q ...*qbittorrent.QBittorrentClient) {
	a.QbittorrentClients = append(a.QbittorrentClients, q...)
}

func (a *App) SetPlexClient(p *plex.PlexClient) {
	a.Plex = p
}

func (a *App) Run() {
	if len(a.QbittorrentClients) == 0 {
		panic("there are no qbittorrent clients")
	}
	if a.Plex == nil {
		panic("there is no plex client")
	}
	// if webhooks, use API, else, use polling
	if a.Plex.WebhooksEnabled {

	} else {
		go a.RunPolling()
	}
}

func (a *App) RunPolling() {
	for {
		select {
		case <-a.AppChan:
			fmt.Println("AppChan")
		case <-a.Quit:
			fmt.Println("Quit signal received... ")
			return
		case <-time.After(time.Second * 5):
			fmt.Println("Polling...")
			// for _, qbt := range a.QbittorrentClients {
			// 	for _, libName := range qbt.PlexLibrariesToPause {
			// 		plexStreams, err := a.Plex.GetLibrariesBeingStreamed()
			// 		if err != nil {
			// 			panic(err)
			// 		}
			// 		for _, ps := range plexStreams {
			// 			if ps.Name == libName {
			// 				qbt.PauseAllTorrents()
			// 			}
			// 		}
			// 	}
			// }
		}
	}
}
