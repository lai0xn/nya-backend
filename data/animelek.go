package data

import (
	"fmt"

	"github.com/gocolly/colly"
)

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
