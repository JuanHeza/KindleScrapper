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
type MangadexChapter struct {
	Description string  `json:"description,omitempty"`
	Volume      string  `json:"volume,omitempty"`
	FileName    string  `json:"fileName,omitempty"`
	Locale      string  `json:"locale,omitempty"`
	CreatedAt   string  `json:"createdAt,omitempty"`
	UpdatesAt   string  `json:"updatedAt,omitempty"`
	Version     float32 `json:"version,omitempty"`
}
type MangadexPage struct {
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

func (mClient *Mangadex) getchapters(id string) {
	v := url.Values{}
	v.Set("limit", "10")
	v.Set("offset", "0")
	v.Set("translatedLanguage[]", "es")
	v.Add("translatedLanguage[]", "en")
	v.Add("translatedLanguage[]", "fr")
	v.Set("order[createdAt]", "asc")
	v.Add("order[updatedAt]", "asc")
	v.Add("order[publishAt]", "asc")
	v.Add("order[readableAt]", "asc")
	v.Add("order[volume]", "asc")
	v.Add("order[chapter]", "asc")
	v.Set("contentRating[]", "safe")
	v.Add("contentRating[]", "suggestive")
	v.Add("contentRating[]", "erotica")
	//v.Add("contentRating[]", "pornographic")
	//v.Set("includeFutureUpdates", "1")
	cl, err := mClient.client.Chapter.GetMangaChapters(id, v)
	if err != nil {
		return //&m.MangaList{}, err
	}
	for _, chapter := range cl.Data {
		fmt.Printf("%v - %v - %v \n", chapter.Attributes.Title, chapter.Attributes.Volume, chapter.Attributes.Chapter)
	}
}

func (mClient *Mangadex) getPages() {
}
