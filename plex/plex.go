package plex

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type PlexClientCreate struct {
	Token           string
	Host            string
	Port            string
	WebhooksEnabled bool
}

type PlexClient struct {
	Token           string
	Host            string
	Port            string
	URL             string
	WebhooksEnabled bool
}

type PlexRequest struct {
	method   string
	endpoint string
}

var libraryNameCache = make(map[int]*PlexLibrary)

func New(p PlexClientCreate) *PlexClient {
	return &PlexClient{
		Host:            p.Host,
		Token:           p.Token,
		Port:            p.Port,
		URL:             fmt.Sprintf("http://%s:%s", p.Host, p.Port),
		WebhooksEnabled: p.WebhooksEnabled,
	}
}

func (p *PlexClient) Fetch(r PlexRequest) (*http.Response, error) {
	url := fmt.Sprintf("%s%s?X-Plex-Token=%s", p.URL, r.endpoint, p.Token)
	req, err := http.NewRequest(r.method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func unMarshalPlexResponse[T PlexResponseData](res *http.Response) (*T, error) {
	var responseData T
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		panic(err)
	}
	return &responseData, nil
}

func (p *PlexClient) GetLibrarySectionById(i int) (*PlexLibrarySelection, error) {
	req := PlexRequest{
		method:   "GET",
		endpoint: fmt.Sprintf("/library/sections/%d", i),
	}
	res, err := p.Fetch(req)
	if err != nil {
		panic(err)
	}
	data, err := unMarshalPlexResponse[PlexLibrarySelection](res)
	if err != nil {
		panic(err)
	}
	return data, nil

}

func (p *PlexClient) GetLibrarySectionsAll() (*PlexLibrarySelectionAll, error) {
	req := PlexRequest{
		method:   "GET",
		endpoint: "/library/sections",
	}
	res, err := p.Fetch(req)
	if err != nil {
		panic(err)
	}
	data, err := unMarshalPlexResponse[PlexLibrarySelectionAll](res)
	if err != nil {
		panic(err)
	}
	return data, nil

}

func (p *PlexClient) GetAccount() (*PlexAccount, error) {
	req := PlexRequest{
		method:   "GET",
		endpoint: "/accounts/1",
	}
	res, err := p.Fetch(req)
	if err != nil {
		panic(err)
	}
	data, err := unMarshalPlexResponse[PlexAccount](res)
	if err != nil {
		panic(err)
	}
	return data, nil

}

func (p *PlexClient) GetStatusSessions() (*StatusSessions, error) {
	req := PlexRequest{
		method:   "GET",
		endpoint: "/status/sessions",
	}
	res, err := p.Fetch(req)
	if err != nil {
		panic(err)
	}
	data, err := unMarshalPlexResponse[StatusSessions](res)
	if err != nil {
		panic(err)
	}
	return data, nil

}

func (p PlexClient) getLibraryIdsBeingStreamed() []int {
	data, err := p.GetStatusSessions()
	if err != nil {
		panic(err)
	}
	var libraries []int
	for _, stream := range data.MediaContainer.Metadata {
		id, err := strconv.Atoi(stream.LibrarySectionID)
		if err != nil {
			panic(err)
		}
		libraries = append(libraries, id)

	}
	return libraries
}

func (p PlexClient) GetLibrariesBeingStreamed() ([]*PlexLibrary, error) {
	var libraries []*PlexLibrary
	for _, id := range p.getLibraryIdsBeingStreamed() {
		lib, err := p.GetLibraryByID(id)
		if err != nil {
			return nil, err
		}
		libraries = append(libraries, lib)

	}
	return libraries, nil

}

func (p PlexClient) GetAllLibraries() ([]*PlexLibrary, error) {
	var libs []*PlexLibrary
	data, err := p.GetLibrarySectionsAll()
	if err != nil {
		return nil, err
	}
	for _, d := range data.MediaContainer.Directory {
		lib := PlexLibrary{
			Name: d.Title,
		}
		libs = append(libs, &lib)
	}
	return libs, nil

}

func (p PlexClient) getLibraryName(i int) (string, error) {
	data, err := p.GetLibrarySectionById(i)
	if err != nil {
		return "", err
	}
	return data.MediaContainer.Directory[0].Title, nil
}

func (p PlexClient) GetLibraryByID(i int) (*PlexLibrary, error) {
	var library *PlexLibrary
	library, ok := libraryNameCache[i]
	if !ok {
		libraryName, err := p.getLibraryName(i)
		if err != nil {
			return library, err
		}
		library = &PlexLibrary{
			Name: libraryName,
		}
		libraryNameCache[i] = library
	}
	return library, nil

}
