package app

import (
	"fmt"
	"os"
	"plex-qbt-controller/plex"
	"plex-qbt-controller/qbittorrent"
	"sync"
	"time"
)

type App struct {
	QbittorrentClients []*qbittorrent.QBittorrentClient
	Plex               *plex.PlexClient
	Quit               chan os.Signal
	AppChan            chan bool
	Wait               sync.WaitGroup
}

func (a *App) AddQBittorrentClient(q ...*qbittorrent.QBittorrentClient) {
	a.QbittorrentClients = append(a.QbittorrentClients, q...)
}

func (a *App) SetPlexClient(p *plex.PlexClient) {
	a.Plex = p
}

func (a *App) Run() {
	a.Wait.Add(1)
	if len(a.QbittorrentClients) == 0 {
		panic("there are no qbittorrent clients")
	}
	if a.Plex == nil {
		panic("there is no plex client")
	}
	// if webhooks, use API, else, use polling
	if a.Plex.WebhooksEnabled {

	} else {
		a.Wait.Add(1)
		fmt.Println("no webhooks enabled. Using polling...")
		go a.runPolling()

	}
}

func (a *App) runPolling() {
	defer a.Wait.Done()
	for {
		select {
		case <-a.AppChan:
			fmt.Println("AppChan")
		case <-a.Quit:
			fmt.Println("Quit signal received... ")
			a.Wait.Done()
			return
		case <-time.After(time.Second * 1):
			fmt.Println("Polling...")
			plexStreams, err := a.Plex.GetLibrariesBeingStreamed()
			if err != nil {
				panic(err)
			}
			for _, ps := range plexStreams {
				for _, qbt := range a.QbittorrentClients {
					for _, libName := range qbt.PlexLibrariesToPause {
						if ps.Name == libName {
							qbt.PauseAllTorrents()
						}
					}
				}
			}

		}
	}
}
