package qbittorrent

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type NewBittorrentCreateParams struct {
	Username             string
	Password             string
	Port                 string
	Host                 string
	PlexLibrariesToPause []string
}

type QBittorrentClient struct {
	Username             string
	Password             string
	Host                 string
	Port                 string
	PlexLibrariesToPause []string
	// Cookie   string
	// Cookie http.Cookie
	Cookies []*http.Cookie
	url     string
}

func (q *QBittorrentClient) defaultRequest() *QBittorrentClientRequestParams {
	return &QBittorrentClientRequestParams{
		method:            "GET",
		endpoint:          "/torrents/pause",
		QBittorrentClient: q,
		Cookies:           q.Cookies,
		ContentType:       "application/x-www-form-urlencoded",
	}
}

func (q *QBittorrentClientRequestParams) New() (*http.Response, error) {
	url := q.QBittorrentClient.url + q.endpoint + q.Querystring
	req, err := http.NewRequest(q.method, url, bytes.NewBufferString(q.Payload))
	req.Header.Set("Content-Type", q.ContentType)
	if err != nil {
		panic(err)
	}

	for _, cookie := range q.Cookies {
		req.AddCookie(cookie)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return res, nil

}

func New(p NewBittorrentCreateParams) (*QBittorrentClient, error) {
	// url := "http://localhost:8080/api/v2/auth/login"
	// username := "admin"
	// password := "admin123"
	url := fmt.Sprintf("http://%s:%s/api/v2", p.Host, p.Port)
	bodyStr := []byte(fmt.Sprintf("username=%s&password=%s", p.Host, p.Port))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/login", url), bytes.NewBuffer(bodyStr))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	cookies := res.Cookies()
	if len(cookies) == 0 {
		panic("Cookie not found")
	}
	return &QBittorrentClient{
		Username: p.Username,
		Password: p.Password,
		Host:     p.Host,
		Port:     p.Port,
		// Cookie:   cookieStr,
		Cookies:              cookies,
		url:                  url,
		PlexLibrariesToPause: p.PlexLibrariesToPause,
	}, nil

}

func (q *QBittorrentClient) Login() (e error) {
	url := fmt.Sprintf("http://%s:%s/api/v2", q.Host, q.Port)
	bodyStr := []byte(fmt.Sprintf("username=%s&password=%s", q.Host, q.Port))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/login", url), bytes.NewBuffer(bodyStr))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return nil
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	cookies := res.Cookies()
	if len(cookies) == 0 {
		return errors.New("could not retrieve cookies")
	}
	q.Cookies = cookies
	return nil
}

func (q *QBittorrentClient) GetTotalBit(torrents ...Torrent) int {
	var total int
	for _, t := range torrents {
		total += int(t.SeedingTime)
	}
	return total

}

func (q *QBittorrentClient) GetAllTorrents(filters ...TorrentState) []*Torrent {
	var queryString string

	if len(filters) > 0 {
		queryString += "?"
		for _, f := range filters {
			filter := TorrentStateMap[f]
			queryString += fmt.Sprintf("filter=%s&", filter)

		}

	}
	r := QBittorrentClientRequestParams{
		method:            "GET",
		endpoint:          "/torrents/info",
		QBittorrentClient: q,
		Cookies:           q.Cookies,
		ContentType:       "application/json",
		Querystring:       queryString,
	}
	res, err := r.New()
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var torrentInfoResponse []*Torrent
	err = json.Unmarshal(body, &torrentInfoResponse)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(body))

	return torrentInfoResponse

}

// TODO CHANGE t.State to correct value
func (q *QBittorrentClient) ActiveTorrents() []*Torrent {
	allTorrents := q.GetAllTorrents()
	var active []*Torrent
	for _, t := range allTorrents {
		if t.State == "active" {
			active = append(active, t)
		}
	}
	return active
}

// TODO CHANGE t.State to correct value
func (q *QBittorrentClient) PausedTorrents() []*Torrent {
	allTorrents := q.GetAllTorrents()
	var active []*Torrent
	for _, t := range allTorrents {
		if t.State == "paused" {
			active = append(active, t)
		}
	}
	return active
}

func (q *QBittorrentClient) ResumeAllTorrents() {
	t := Torrent{
		InfohashV1: "all",
	}
	q.ResumeTorrents(&t)

}

func (q *QBittorrentClient) ResumeTorrents(t ...*Torrent) error {
	r := q.defaultRequest()
	var hashes []string
	for _, tor := range t {
		hashes = append(hashes, tor.InfohashV1)
	}
	hashesStr := strings.Join(hashes, "|")
	r.Payload = fmt.Sprintf("hashes=%s", hashesStr)
	r.method = "POST"
	r.endpoint = "/torrents/resume"
	_, err := r.New()
	if err != nil {
		return err
	}
	// fmt.Println(res.StatusCode, res.Request.URL, res.Request.Body)
	return nil

}

func (q *QBittorrentClient) PauseAllTorrents() {
	t := Torrent{
		InfohashV1: "all",
	}
	q.PauseTorrents(&t)

}

func (q *QBittorrentClient) PauseTorrents(t ...*Torrent) {
	r := q.defaultRequest()
	var hashes []string
	for _, tor := range t {
		hashes = append(hashes, tor.InfohashV1)
	}
	hashesStr := strings.Join(hashes, "|")
	r.Querystring = fmt.Sprintf("hashes=%s", hashesStr)
	r.method = "POST"
	r.endpoint = "/torrents/pause"

	res, err := r.New()
	if err != nil {
		panic(err)
	}
	fmt.Println(res.StatusCode, res.Request.URL, res.Request.URL, hashesStr)
}
