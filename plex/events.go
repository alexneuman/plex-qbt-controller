package plex

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type PlexEvent int

const (
	PAUSED PlexEvent = iota + 1
	RESUMED
)

var PlexEventMap = map[string]PlexEvent{
	"PAUSED":  PAUSED,
	"RESUMED": RESUMED,
}

func ParsePlexEventDataFromPayload(b []byte) []byte {
	r := regexp.MustCompile(`({.event.:.*})(\\r)`)
	match := r.FindAllString(string(b), -1)[0]
	match = strings.ReplaceAll(match, "\\r", "")
	match = strings.ReplaceAll(match, "\\", "")
	// bob := string(match)
	// fmt.Println(bob)
	matchBytes := []byte(match)
	return matchBytes

}

func unMarshalPlexEventJSON(b []byte) *PlexEventBody {
	var plexEventBody PlexEventBody

	err := json.Unmarshal(b, &plexEventBody)
	if err != nil {
		panic(err)
	}
	return &plexEventBody

}

func ParsePlexEventBodyFromRequest(r *http.Request) (*PlexEventBody, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return &PlexEventBody{}, err

	}
	parsedJson := ParsePlexEventDataFromPayload(body)
	plexEventBody := unMarshalPlexEventJSON(parsedJson)

	return plexEventBody, nil

}
