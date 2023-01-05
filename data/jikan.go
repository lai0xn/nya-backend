package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

	url := "https://api.jikan.moe/v4/top/anime?limit=12"

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

func (AnimeWrapper) SearchAnimeByID(id string) (SearchData, error) {

	url := fmt.Sprintf("https://api.jikan.moe/v4/anime/%s/full", id)
	fmt.Println(url)
	responese, err := http.Get(url)
	if err != nil {
		return SearchData{}, errors.New("Something Went Wrong")
	}
	responseData, err := ioutil.ReadAll(responese.Body)
	if err != nil {
		return SearchData{}, errors.New("Something Went Wrong")
	}

	var responseObject SearchData
	json.Unmarshal(responseData, &responseObject)
	return responseObject, nil

}

func (AnimeWrapper) RandomAnime() (SearchData, error) {

	url := "https://api.jikan.moe/v4/random/anime"
	fmt.Println(url)
	responese, err := http.Get(url)
	if err != nil {
		return SearchData{}, errors.New("Something Went Wrong")
	}
	responseData, err := ioutil.ReadAll(responese.Body)
	if err != nil {
		return SearchData{}, errors.New("Something Went Wrong")
	}

	var responseObject SearchData
	json.Unmarshal(responseData, &responseObject)
	return responseObject, nil

}
