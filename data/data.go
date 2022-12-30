package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gocolly/colly"
)

var blkom_base_url string = "https://animeblkom.net"

type AnimeWrapper struct{}

func (AnimeWrapper) SearchAnime(query string) (Response, error) {
	query = strings.ReplaceAll(query, " ", "%20")
	url := "https://api.jikan.moe/v4/anime?q=" + query

	responese, err := http.Get(url)
	if err != nil {
		return Response{}, errors.New("Something Went Wrong")
	}
	responseData, err := ioutil.ReadAll(responese.Body)
	if err != nil {
		return Response{}, errors.New("Something Went Wrong")
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	return responseObject, nil

}

func GetFirstLink(query string) string {
	c := colly.NewCollector(
		colly.Async(true),
	)
	url := blkom_base_url + "/search?query=" + query
	var watchlinks []string
	c.OnHTML(".name", func(h *colly.HTMLElement) {
		link := h.ChildAttr("a", "href")
		page_link := blkom_base_url + link + "/"
		watchlinks = append(watchlinks, page_link)

	})

	c.Visit(url)
	c.Wait()
	return watchlinks[0]
}

func (AnimeWrapper) AnimeWatchLink(query string, episode string) string {
	c := colly.NewCollector(
		colly.Async(true),
	)
	var watchlink string
	var url string = GetFirstLink(query) + episode
	fmt.Println(url)
	c.OnHTML(".video", func(h *colly.HTMLElement) {
		watchlink = h.Attr("src")
		fmt.Println(h.Attr("src"))
	})
	c.Visit(url)

	c.Wait()
	return watchlink

}

func (AnimeWrapper) DownloadLinks(query string, episode string) []DownloadLink {
	c := colly.NewCollector(
		colly.Async(true),
	)
	url := fmt.Sprintf("https://animelek.me/episode/%s-%s-الحلقة", query, episode)
	fmt.Println(url)
	var downloadLinks []DownloadLink
	c.OnHTML("#downloads", func(h *colly.HTMLElement) {
		h.ForEach(".watch", func(_ int, h *colly.HTMLElement) {
			link := h.ChildAttr("a", "href")
			downloadLinks = append(downloadLinks, DownloadLink{Url: link, ServerName: h.ChildText("a")})

		})
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting" + r.URL.Path)
	})
	c.Visit(url)
	c.Wait()
	return downloadLinks
}

func (AnimeWrapper) GetUpcomingAnimes() (Response, error) {

	url := "https://api.jikan.moe/v4/seasons/upcoming?limit=12"

	responese, err := http.Get(url)
	if err != nil {
		return Response{}, errors.New("Something Went Wrong")
	}
	responseData, err := ioutil.ReadAll(responese.Body)
	if err != nil {
		return Response{}, errors.New("Something Went Wrong")
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	return responseObject, nil

}

func (AnimeWrapper) GetTopAnimes() (Response, error) {

	url := "https://api.jikan.moe/v4/seasons/top?limit=12"

	responese, err := http.Get(url)
	if err != nil {
		return Response{}, errors.New("Something Went Wrong")
	}
	responseData, err := ioutil.ReadAll(responese.Body)
	if err != nil {
		return Response{}, errors.New("Something Went Wrong")
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	return responseObject, nil

}
