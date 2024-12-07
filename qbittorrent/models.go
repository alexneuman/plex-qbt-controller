package qbittorrent

import (
	"net/http"
)

type TorrentState int

const (
	ALL = iota
	RUNNING
	SEEDING
	RESUMED
	DOWNLOADING
	COMPLETED
	PAUSED
	ACTIVE
	STALLED_UPLOADING
)

var TorrentStateMap = map[TorrentState]string{
	ALL:               "all",
	RUNNING:           "running",
	SEEDING:           "seeding",
	RESUMED:           "resumed",
	DOWNLOADING:       "downloading",
	COMPLETED:         "completed",
	PAUSED:            "paused",
	ACTIVE:            "active",
	STALLED_UPLOADING: "stalled_uploading",
}

type QBittorrentClientRequestParams struct {
	method      string
	endpoint    string
	Querystring string
	*QBittorrentClient
	Cookies     []*http.Cookie
	ContentType string
	Payload     string
}

type Torrent struct {
	AddedOn                  int64   `json:"added_on"`
	AmountLeft               int64   `json:"amount_left"`
	AutoTMM                  bool    `json:"auto_tmm"`
	Availability             float64 `json:"availability"`
	Category                 string  `json:"category"`
	Completed                int64   `json:"completed"`
	CompletionOn             int64   `json:"completion_on"`
	ContentPath              string  `json:"content_path"`
	DlLimit                  int64   `json:"dl_limit"`
	DlSpeed                  int64   `json:"dlspeed"`
	DownloadPath             string  `json:"download_path"`
	Downloaded               int64   `json:"downloaded"`
	DownloadedSession        int64   `json:"downloaded_session"`
	Eta                      int64   `json:"eta"`
	FirstLastPiecePriority   bool    `json:"f_l_piece_prio"`
	ForceStart               bool    `json:"force_start"`
	Hash                     string  `json:"hash"`
	InactiveSeedingTimeLimit int64   `json:"inactive_seeding_time_limit"`
	InfohashV1               string  `json:"infohash_v1"`
	InfohashV2               string  `json:"infohash_v2"`
	LastActivity             int64   `json:"last_activity"`
	MagnetURI                string  `json:"magnet_uri"`
	MaxInactiveSeedingTime   int64   `json:"max_inactive_seeding_time"`
	MaxRatio                 float64 `json:"max_ratio"`
	MaxSeedingTime           int64   `json:"max_seeding_time"`
	Name                     string  `json:"name"`
	NumComplete              int     `json:"num_complete"`
	NumIncomplete            int     `json:"num_incomplete"`
	NumLeechs                int     `json:"num_leechs"`
	NumSeeds                 int     `json:"num_seeds"`
	Priority                 int     `json:"priority"`
	Progress                 float64 `json:"progress"`
	Ratio                    float64 `json:"ratio"`
	RatioLimit               float64 `json:"ratio_limit"`
	SavePath                 string  `json:"save_path"`
	SeedingTime              int64   `json:"seeding_time"`
	SeedingTimeLimit         int64   `json:"seeding_time_limit"`
	SeenComplete             int64   `json:"seen_complete"`
	SequentialDownload       bool    `json:"seq_dl"`
	Size                     int64   `json:"size"`
	State                    string  `json:"state"`
	SuperSeeding             bool    `json:"super_seeding"`
	Tags                     string  `json:"tags"`
	TimeActive               int64   `json:"time_active"`
	TotalSize                int64   `json:"total_size"`
	Tracker                  string  `json:"tracker"`
	TrackersCount            int     `json:"trackers_count"`
	UploadLimit              int64   `json:"up_limit"`
	Uploaded                 int64   `json:"uploaded"`
	UploadedSession          int64   `json:"uploaded_session"`
	UploadSpeed              int64   `json:"upspeed"`
}
