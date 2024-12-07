package plex

type PlexResponseData interface {
	isPlexResponse()
}

type PlexLibrary struct {
	Name string
}

// /library/sections
type PlexLibrarySelectionAll struct {
	MediaContainer struct {
		Size      int    `json:"size"`
		AllowSync bool   `json:"allowSync"`
		Title1    string `json:"title1"`
		Directory []struct {
			AllowSync        bool   `json:"allowSync"`
			Art              string `json:"art"`
			Composite        string `json:"composite"`
			Filters          bool   `json:"filters"`
			Refreshing       bool   `json:"refreshing"`
			Thumb            string `json:"thumb"`
			Key              string `json:"key"`
			Type             string `json:"type"`
			Title            string `json:"title"`
			Agent            string `json:"agent"`
			Scanner          string `json:"scanner"`
			Language         string `json:"language"`
			UUID             string `json:"uuid"`
			UpdatedAt        int    `json:"updatedAt"`
			CreatedAt        int    `json:"createdAt"`
			ScannedAt        int    `json:"scannedAt"`
			Content          bool   `json:"content"`
			Directory        bool   `json:"directory"`
			ContentChangedAt int    `json:"contentChangedAt"`
			Hidden           int    `json:"hidden"`
			Location         []struct {
				ID   int    `json:"id"`
				Path string `json:"path"`
			} `json:"Location"`
		} `json:"Directory"`
	} `json:"MediaContainer"`
}

// /library/sections/1
type PlexLibrarySelection struct {
	MediaContainer struct {
		Size             int    `json:"size"`
		AllowSync        bool   `json:"allowSync"`
		Art              string `json:"art"`
		Content          string `json:"content"`
		Identifier       string `json:"identifier"`
		LibrarySectionID int    `json:"librarySectionID"`
		MediaTagPrefix   string `json:"mediaTagPrefix"`
		MediaTagVersion  int    `json:"mediaTagVersion"`
		Thumb            string `json:"thumb"`
		Title1           string `json:"title1"`
		ViewGroup        string `json:"viewGroup"`
		Directory        []struct {
			Key       string `json:"key"`
			Title     string `json:"title"`
			Secondary bool   `json:"secondary,omitempty"`
			Prompt    string `json:"prompt,omitempty"`
			Search    bool   `json:"search,omitempty"`
		} `json:"Directory"`
	} `json:"MediaContainer"`
}

// /accounts/1
type PlexAccount struct {
	MediaContainer struct {
		Size       int    `json:"size"`
		Identifier string `json:"identifier"`
		Account    []struct {
			ID                      int    `json:"id"`
			Key                     string `json:"key"`
			Name                    string `json:"name"`
			DefaultAudioLanguage    string `json:"defaultAudioLanguage"`
			AutoSelectAudio         bool   `json:"autoSelectAudio"`
			DefaultSubtitleLanguage string `json:"defaultSubtitleLanguage"`
			SubtitleMode            int    `json:"subtitleMode"`
			Thumb                   string `json:"thumb"`
		} `json:"Account"`
	} `json:"MediaContainer"`
}

// /devices
type Devices struct {
	MediaContainer struct {
		Size       int    `json:"size"`
		Identifier string `json:"identifier"`
		Device     []struct {
			ID               int    `json:"id"`
			Name             string `json:"name"`
			Platform         string `json:"platform"`
			ClientIdentifier string `json:"clientIdentifier"`
			CreatedAt        int    `json:"createdAt"`
		} `json:"Device"`
	} `json:"MediaContainer"`
}

// /devices/1
type Device struct {
	MediaContainer struct {
		Size       int    `json:"size"`
		Identifier string `json:"identifier"`
		Device     []struct {
			ID               int    `json:"id"`
			Name             string `json:"name"`
			Platform         string `json:"platform"`
			ClientIdentifier string `json:"clientIdentifier"`
			CreatedAt        int    `json:"createdAt"`
		} `json:"Device"`
	} `json:"MediaContainer"`
}

// /status/sessions
type StatusSessions struct {
	MediaContainer struct {
		Size     int `json:"size"`
		Metadata []struct {
			AddedAt               int     `json:"addedAt"`
			Art                   string  `json:"art"`
			AudienceRating        float64 `json:"audienceRating"`
			AudienceRatingImage   string  `json:"audienceRatingImage"`
			ChapterSource         string  `json:"chapterSource"`
			ContentRating         string  `json:"contentRating"`
			Duration              int     `json:"duration"`
			GUID                  string  `json:"guid"`
			Key                   string  `json:"key"`
			LastViewedAt          int     `json:"lastViewedAt"`
			LibrarySectionID      string  `json:"librarySectionID"`
			LibrarySectionKey     string  `json:"librarySectionKey"`
			LibrarySectionTitle   string  `json:"librarySectionTitle"`
			OriginallyAvailableAt string  `json:"originallyAvailableAt"`
			PrimaryExtraKey       string  `json:"primaryExtraKey"`
			Rating                float64 `json:"rating"`
			RatingImage           string  `json:"ratingImage"`
			RatingKey             string  `json:"ratingKey"`
			SessionKey            string  `json:"sessionKey"`
			Slug                  string  `json:"slug"`
			Studio                string  `json:"studio"`
			Summary               string  `json:"summary"`
			Tagline               string  `json:"tagline"`
			Thumb                 string  `json:"thumb"`
			Title                 string  `json:"title"`
			TitleSort             string  `json:"titleSort"`
			Type                  string  `json:"type"`
			UpdatedAt             int     `json:"updatedAt"`
			ViewOffset            int     `json:"viewOffset"`
			Year                  int     `json:"year"`
			Media                 []struct {
				AspectRatio      string `json:"aspectRatio"`
				AudioChannels    int    `json:"audioChannels"`
				AudioCodec       string `json:"audioCodec"`
				Bitrate          int    `json:"bitrate"`
				Container        string `json:"container"`
				Duration         int    `json:"duration"`
				HasVoiceActivity string `json:"hasVoiceActivity"`
				Height           int    `json:"height"`
				ID               string `json:"id"`
				VideoCodec       string `json:"videoCodec"`
				VideoFrameRate   string `json:"videoFrameRate"`
				VideoProfile     string `json:"videoProfile"`
				VideoResolution  string `json:"videoResolution"`
				Width            int    `json:"width"`
				Selected         bool   `json:"selected"`
				Part             []struct {
					Container    string `json:"container"`
					Duration     int    `json:"duration"`
					File         string `json:"file"`
					ID           string `json:"id"`
					Indexes      string `json:"indexes"`
					Key          string `json:"key"`
					Size         int64  `json:"size"`
					VideoProfile string `json:"videoProfile"`
					Decision     string `json:"decision"`
					Selected     bool   `json:"selected"`
					Stream       []struct {
						BitDepth             int     `json:"bitDepth,omitempty"`
						Bitrate              int     `json:"bitrate"`
						ChromaLocation       string  `json:"chromaLocation,omitempty"`
						ChromaSubsampling    string  `json:"chromaSubsampling,omitempty"`
						Codec                string  `json:"codec"`
						CodedHeight          int     `json:"codedHeight,omitempty"`
						CodedWidth           int     `json:"codedWidth,omitempty"`
						ColorPrimaries       string  `json:"colorPrimaries,omitempty"`
						ColorRange           string  `json:"colorRange,omitempty"`
						ColorSpace           string  `json:"colorSpace,omitempty"`
						ColorTrc             string  `json:"colorTrc,omitempty"`
						Default              bool    `json:"default,omitempty"`
						DisplayTitle         string  `json:"displayTitle"`
						ExtendedDisplayTitle string  `json:"extendedDisplayTitle"`
						FrameRate            float64 `json:"frameRate,omitempty"`
						Height               int     `json:"height,omitempty"`
						ID                   string  `json:"id"`
						Index                int     `json:"index"`
						Level                int     `json:"level,omitempty"`
						Profile              string  `json:"profile,omitempty"`
						RefFrames            int     `json:"refFrames,omitempty"`
						StreamType           int     `json:"streamType"`
						Width                int     `json:"width,omitempty"`
						Location             string  `json:"location,omitempty"`
						AudioChannelLayout   string  `json:"audioChannelLayout,omitempty"`
						Channels             int     `json:"channels,omitempty"`
						Language             string  `json:"language,omitempty"`
						LanguageCode         string  `json:"languageCode,omitempty"`
						LanguageTag          string  `json:"languageTag,omitempty"`
						SamplingRate         int     `json:"samplingRate,omitempty"`
						Selected             bool    `json:"selected,omitempty"`
						CanAutoSync          string  `json:"canAutoSync,omitempty"`
						HeaderCompression    bool    `json:"headerCompression,omitempty"`
						Decision             string  `json:"decision,omitempty"`
					} `json:"Stream"`
				} `json:"Part"`
			} `json:"Media"`
			UltraBlurColors []struct {
				BottomLeft  string `json:"bottomLeft"`
				BottomRight string `json:"bottomRight"`
				TopLeft     string `json:"topLeft"`
				TopRight    string `json:"topRight"`
			} `json:"UltraBlurColors"`
			Genre []struct {
				Count  string `json:"count"`
				Filter string `json:"filter"`
				ID     string `json:"id"`
				Tag    string `json:"tag"`
			} `json:"Genre"`
			Country []struct {
				Count  string `json:"count"`
				Filter string `json:"filter"`
				ID     string `json:"id"`
				Tag    string `json:"tag"`
			} `json:"Country"`
			Rating0 []struct {
				Count string `json:"count"`
				Image string `json:"image"`
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"Rating"`
			Director []struct {
				Filter string `json:"filter"`
				ID     string `json:"id"`
				Tag    string `json:"tag"`
				TagKey string `json:"tagKey"`
				Thumb  string `json:"thumb"`
			} `json:"Director"`
			Writer []struct {
				Count  string `json:"count"`
				Filter string `json:"filter"`
				ID     string `json:"id"`
				Tag    string `json:"tag"`
				TagKey string `json:"tagKey"`
				Thumb  string `json:"thumb"`
			} `json:"Writer"`
			Role []struct {
				Count  string `json:"count,omitempty"`
				Filter string `json:"filter"`
				ID     string `json:"id"`
				Role   string `json:"role"`
				Tag    string `json:"tag"`
				TagKey string `json:"tagKey"`
				Thumb  string `json:"thumb,omitempty"`
			} `json:"Role"`
			Producer []struct {
				Count  string `json:"count,omitempty"`
				Filter string `json:"filter"`
				ID     string `json:"id"`
				Tag    string `json:"tag"`
				TagKey string `json:"tagKey"`
				Thumb  string `json:"thumb"`
			} `json:"Producer"`
			Label []struct {
				Count  string `json:"count"`
				Filter string `json:"filter"`
				ID     string `json:"id"`
				Tag    string `json:"tag"`
			} `json:"Label"`
			Field []struct {
				Locked bool   `json:"locked"`
				Name   string `json:"name"`
			} `json:"Field"`
			User struct {
				ID    string `json:"id"`
				Thumb string `json:"thumb"`
				Title string `json:"title"`
			} `json:"User"`
			Player struct {
				Address             string `json:"address"`
				MachineIdentifier   string `json:"machineIdentifier"`
				Model               string `json:"model"`
				Platform            string `json:"platform"`
				PlatformVersion     string `json:"platformVersion"`
				Product             string `json:"product"`
				Profile             string `json:"profile"`
				RemotePublicAddress string `json:"remotePublicAddress"`
				State               string `json:"state"`
				Title               string `json:"title"`
				Version             string `json:"version"`
				Local               bool   `json:"local"`
				Relayed             bool   `json:"relayed"`
				Secure              bool   `json:"secure"`
				UserID              int    `json:"userID"`
			} `json:"Player"`
			Session struct {
				ID        string `json:"id"`
				Bandwidth int    `json:"bandwidth"`
				Location  string `json:"location"`
			} `json:"Session"`
		} `json:"Metadata"`
	} `json:"MediaContainer"`
}

// sent from the Plex webhook on event
type PlexEventBody struct {
	Event   string `json:"event"`
	User    bool   `json:"user"`
	Owner   bool   `json:"owner"`
	Account struct {
		ID    int    `json:"id"`
		Thumb string `json:"thumb"`
		Title string `json:"title"`
	} `json:"Account"`
	Server struct {
		Title string `json:"title"`
		UUID  string `json:"uuid"`
	} `json:"Server"`
	Player struct {
		Local         bool   `json:"local"`
		PublicAddress string `json:"publicAddress"`
		Title         string `json:"title"`
		UUID          string `json:"uuid"`
	} `json:"Player"`
	Metadata struct {
		LibrarySectionType    string  `json:"librarySectionType"`
		RatingKey             string  `json:"ratingKey"`
		Key                   string  `json:"key"`
		GUID                  string  `json:"guid"`
		Slug                  string  `json:"slug"`
		Studio                string  `json:"studio"`
		Type                  string  `json:"type"`
		Title                 string  `json:"title"`
		TitleSort             string  `json:"titleSort"`
		LibrarySectionTitle   string  `json:"librarySectionTitle"`
		LibrarySectionID      int     `json:"librarySectionID"`
		LibrarySectionKey     string  `json:"librarySectionKey"`
		ContentRating         string  `json:"contentRating"`
		Summary               string  `json:"summary"`
		Rating                float64 `json:"rating"`
		AudienceRating        float64 `json:"audienceRating"`
		ViewOffset            int     `json:"viewOffset"`
		LastViewedAt          int     `json:"lastViewedAt"`
		Year                  int     `json:"year"`
		Tagline               string  `json:"tagline"`
		Thumb                 string  `json:"thumb"`
		Art                   string  `json:"art"`
		Duration              int     `json:"duration"`
		OriginallyAvailableAt string  `json:"originallyAvailableAt"`
		AddedAt               int     `json:"addedAt"`
		UpdatedAt             int     `json:"updatedAt"`
		AudienceRatingImage   string  `json:"audienceRatingImage"`
		ChapterSource         string  `json:"chapterSource"`
		PrimaryExtraKey       string  `json:"primaryExtraKey"`
		RatingImage           string  `json:"ratingImage"`
		Image                 []struct {
			Alt  string `json:"alt"`
			Type string `json:"type"`
			URL  string `json:"url"`
		} `json:"Image"`
		UltraBlurColors struct {
			TopLeft     string `json:"topLeft"`
			TopRight    string `json:"topRight"`
			BottomRight string `json:"bottomRight"`
			BottomLeft  string `json:"bottomLeft"`
		} `json:"UltraBlurColors"`
		Genre []struct {
			ID     int    `json:"id"`
			Filter string `json:"filter"`
			Tag    string `json:"tag"`
			Count  int    `json:"count"`
		} `json:"Genre"`
		Country []struct {
			ID     int    `json:"id"`
			Filter string `json:"filter"`
			Tag    string `json:"tag"`
			Count  int    `json:"count"`
		} `json:"Country"`
		GUID0 []struct {
			ID string `json:"id"`
		} `json:"Guid"`
		Rating0 []struct {
			Image string  `json:"image"`
			Value float64 `json:"value"`
			Type  string  `json:"type"`
			Count int     `json:"count"`
		} `json:"Rating"`
		Director []struct {
			ID     int    `json:"id"`
			Filter string `json:"filter"`
			Tag    string `json:"tag"`
			TagKey string `json:"tagKey"`
			Thumb  string `json:"thumb"`
		} `json:"Director"`
		Writer []struct {
			ID     int    `json:"id"`
			Filter string `json:"filter"`
			Tag    string `json:"tag"`
			TagKey string `json:"tagKey"`
			Count  int    `json:"count"`
			Thumb  string `json:"thumb"`
		} `json:"Writer"`
		Role []struct {
			ID     int    `json:"id"`
			Filter string `json:"filter"`
			Tag    string `json:"tag"`
			TagKey string `json:"tagKey"`
			Count  int    `json:"count,omitempty"`
			Role   string `json:"role"`
			Thumb  string `json:"thumb,omitempty"`
		} `json:"Role"`
		Producer []struct {
			ID     int    `json:"id"`
			Filter string `json:"filter"`
			Tag    string `json:"tag"`
			TagKey string `json:"tagKey"`
			Count  int    `json:"count,omitempty"`
			Thumb  string `json:"thumb"`
		} `json:"Producer"`
		Label []struct {
			ID     int    `json:"id"`
			Filter string `json:"filter"`
			Tag    string `json:"tag"`
			Count  int    `json:"count"`
		} `json:"Label"`
		Field []struct {
			Locked bool   `json:"locked"`
			Name   string `json:"name"`
		} `json:"Field"`
	} `json:"Metadata"`
}

func (PlexLibrarySelection) isPlexResponse()    {}
func (PlexLibrarySelectionAll) isPlexResponse() {}
func (PlexAccount) isPlexResponse()             {}
func (Device) isPlexResponse()                  {}
func (StatusSessions) isPlexResponse()          {}
