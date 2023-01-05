package data

import (
	"fmt"

	"github.com/gocolly/colly"
)

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
