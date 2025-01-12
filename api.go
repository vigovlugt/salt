package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type PageResponse struct {
	Data struct {
		Content    string `json:"content"`
		Pagination struct {
			CurrentPage int `json:"currentPage"`
			LastPage    int `json:"lastPage"`
			Items       int `json:"items"`
		} `json:"pagination"`
	} `json:"data"`
}

func GetPage(page int) (PageResponse, error) {
	requestUrl := fmt.Sprintf("https://nl.pepper.com/nieuw?page=%d&ajax=true&layout=horizontal", page)

	jar, err := cookiejar.New(nil)
	if err != nil {
		return PageResponse{}, err
	}

	cookieUrl, err := url.Parse("https://nl.pepper.com")
	if err != nil {
		return PageResponse{}, err
	}

	jar.SetCookies(
		cookieUrl,
		[]*http.Cookie{
			{
				Name:  "hide_expired",
				Value: "1",
			},
		},
	)

	client := http.Client{
		Jar: jar,
	}

	res, err := client.Get(requestUrl)
	if err != nil {
		return PageResponse{}, err
	}
	defer res.Body.Close()

	var pageResponse PageResponse
	err = json.NewDecoder(res.Body).Decode(&pageResponse)
	if err != nil {
		return PageResponse{}, err
	}

	return pageResponse, nil
}
