package sources

type Content struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	URL    string  `json:"artist"`
	Origin float64 `json:"price"`
}

type Manga struct {
	Serie  string
	Autor  string
	Number int
	URL    string
	pages  []string
	next   string
	prev   string
}

type Relato struct {
	Title       string
	Serie       string
	Description string
	Paragraph   []string
	URL         string
	Autor       string
	next        string
	prev        string
}

type Tags struct {
	Id     string
	Tipo   string
	Nombre string
	URL    string
}
