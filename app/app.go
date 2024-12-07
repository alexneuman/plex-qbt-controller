package app

import (
	"plex-qbt-controller/plex"
	"plex-qbt-controller/qbittorrent"
)

type App struct {
	QbittorrentClients []*qbittorrent.QBittorrentClient
	Plex               *plex.PlexClient
}

func (a *App) AddQBittorrentClient(q ...*qbittorrent.QBittorrentClient) {
	a.QbittorrentClients = append(a.QbittorrentClients, q...)
}

func (a *App) SetPlexClient(p *plex.PlexClient) {
	a.Plex = p
}

func (a *App) Run() {
	// allLibs, err := a.plex.GetAllLibraries()
	// if err != nil {
	// 	panic(err)
	// }
	for {
		if len(a.QbittorrentClients) == 0 {
			panic("there are no qbittorrent clients")
		}
		if a.Plex == nil {
			panic("there is no plex client")
		}
		streamLibs, err := a.Plex.GetLibrariesBeingStreamed()
		if err != nil {
			panic(err)
		}
		// no streams
		if len(streamLibs) == 0 {
			continue
		}
		for _, streamLib := range streamLibs {
			for _, q := range a.QbittorrentClients {
				for _, lib := range q.PlexLibrariesToPause {
					if lib == streamLib.Name {
						go q.PauseAllTorrents()
					}
				}

			}
		}
	}
}
