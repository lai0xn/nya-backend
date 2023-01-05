package data

type Response struct {
	Data []Anime `json:"data"`
}

type SearchData struct {
	Data Anime `json:"data"`
}

type Anime struct {
	Name       string   `json:"title"`
	MalId      int      `json:"mal_id"`
	MalLink    string   `json:"url"`
	Episodes   int      `json:"episodes"`
	Images     Images   `json:"images"`
	Status     string   `json:"status"`
	Season     string   `json:"season"`
	Year       int      `json:"year"`
	Synopsis   string   `json:"synopsis"`
	Source     string   `json:"source"`
	Genres     []Genre  `json:"genres"`
	Score      float64  `json:"score"`
	Ranked     int      `json:"rank"`
	Popularity int      `json:"popularity"`
	Studios    []Studio `json:"studios"`
	Type       string
}

type Studio struct {
	Nmae  string `json:"name"`
	MalID int    `json:"mal_id"`
	Url   string `json:"url"`
}

type Genre struct {
	MalID int    `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	Url   string `json:"url"`
}

type Images struct {
	Image JPG `json:"jpg"`
}
type JPG struct {
	ImageURL   string `json:"image_url"`
	ImageURL_S string `json:"small_image_url"`
	ImageURL_L string `json:"large_image_url"`
}

type DownloadLink struct {
	ServerName string `json:"server"`
	Url        string `json:"download_url"`
}

type Episode struct {
	Poster    string `json:"poster"`
	Episode   string `json:"episode"`
	AnimeName string `json:"anime"`
}
