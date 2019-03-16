package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

func main() {
	db := map[int]string{}
	i := 0
	e := echo.New()

	//request looks like this: urlshortener.com/shorten?url=www.url_to_shorten.com
	e.GET("/shorten", func(c echo.Context) error {
		url := c.QueryParam("url")
		//add http:// to url if it is missing
		if strings.HasPrefix(url, "http://") != true && strings.HasPrefix(url, "https://") != true {
			url = "http://" + url
		}

		//checking if url is reachable
		_, err := http.Get(url)
		if err != nil {
			responseJSON := &shortenResponse{
				Success:      false,
				ShortenedURL: "",
				Error:        "URL seems unreachable",
			}
			println(err.Error())
			return c.JSONPretty(http.StatusOK, responseJSON, "  ")

		}

		db[i] = url
		i = i + 1 //increasing index by one
		defer println("generated shortcut for url: " + url + " on index: " + strconv.Itoa(i))
		responseJSON := &shortenResponse{
			Success:      true,
			ShortenedURL: "localhost:4321/get/" + strconv.Itoa(i-1),
			Error:        "",
		}
		return c.JSONPretty(http.StatusOK, responseJSON, "  ")
	})
	e.GET("/show", func(c echo.Context) error {

		dbData, err := json.Marshal(db)
		if err != nil {
			return c.NoContent(http.StatusOK)
		}
		return c.Blob(http.StatusOK, "application/json", dbData)

	})
	e.GET("/get/:index", func(c echo.Context) error {
		index, err := strconv.Atoi(c.Param("index"))

		//if index is not an integer or there is no corresponding url to index
		if err != nil || len(db[index]) == 0 {
			responseJSON := &getResponse{
				Exists: false,
				URL:    "",
			}
			return c.JSONPretty(http.StatusOK, responseJSON, "  ")
		}

		url := db[index]
		responseJSON := &getResponse{
			Exists: true,
			URL:    url,
		}
		return c.JSONPretty(http.StatusOK, responseJSON, "  ")
	})
	e.Logger.Fatal(e.Start(":4321"))
}
