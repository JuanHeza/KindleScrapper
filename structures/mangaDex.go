package sources

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"

	m "github.com/darylhjd/mangodex"
)

type Mangadex struct {
	user    string
	pass    string
	client  *m.DexClient
	context context.Context
}

type MangadexCover struct {
	Description string  `json:"description,omitempty"`
	Volume      string  `json:"volume,omitempty"`
	FileName    string  `json:"fileName,omitempty"`
	Locale      string  `json:"locale,omitempty"`
	CreatedAt   string  `json:"createdAt,omitempty"`
	UpdatesAt   string  `json:"updatedAt,omitempty"`
	Version     float32 `json:"version,omitempty"`
}

// init set an call the login and creates a Mangadex object
func (obj Mangadex) init() (Mangadex, error) {
	c := m.NewDexClient()

	// Login using your username and password.
	err := c.Auth.Login(os.Getenv("MANGADEX_USER"), os.Getenv("MANGADEX_PASS"))
	if err != nil {
		return Mangadex{}, err
	} else {
		obj.user = os.Getenv("MANGADEX_USER")
		obj.pass = os.Getenv("MANGADEX_PASS")
		obj.client = c
		return obj, nil
	}
}

func (mClient *Mangadex) getMangas() (*m.MangaList, error) {
	v := url.Values{}
	v.Set("limit", "10")
	v.Set("offset", "0")
	v.Set("includes[]", "cover_art")
	v.Set("availableTranslatedLanguage[]", "es")
	v.Add("availableTranslatedLanguage[]", "en")
	v.Add("availableTranslatedLanguage[]", "fr")
	ml, err := mClient.client.Manga.GetMangaList(v)
	if err != nil {
		return &m.MangaList{}, err
	}
	for _, manga := range ml.Data {
		for _, cover := range manga.Relationships {
			if cover.Type == "cover_art" {
				parsed := MangadexCover{}
				coverInfo, ok := cover.Attributes.(*json.RawMessage)
				if !ok {
					return &m.MangaList{}, errors.New("cant be parsed")

				}
				var dat map[string]interface{}
				json.Unmarshal(*coverInfo, &parsed)
				json.Unmarshal(*coverInfo, &dat)
				fmt.Printf("https://uploads.mangadex.org/covers/%v/%v.512.jpg\n", manga.ID, parsed.FileName)
			}
			// Found!
		}
	}
	return ml, nil
}
