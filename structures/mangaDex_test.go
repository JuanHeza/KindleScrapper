package sources

import (
	"testing"
)

func TestMangadex_init(t *testing.T) {
	t.Setenv("MANGADEX_USER", "jk_heza")
	t.Setenv("MANGADEX_PASS", "zocoloco")
	manga, err := Mangadex{}.init()
	if manga.user == "" || manga.pass == "" {
		t.Fatalf(`manga.user = %s, manga.pass = %s and error %v`, manga.user, manga.pass, err)
	}
	t.Log(manga)
}

func TestMangadex_getMangas(t *testing.T) {
	t.Setenv("MANGADEX_USER", "jk_heza")
	t.Setenv("MANGADEX_PASS", "zocoloco")
	manga, err := Mangadex{}.init()
	if err != nil {
		t.Fatalf("Unexpected Error %v", err)
	}
	ml, err := manga.getMangas()
	if err != nil {
		t.Fatalf("Unexpected Error %v", err)
	}
	t.Log(ml)
}
