package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type quote struct {
	Title  string
	Author string
}

// var quotes []quote = make([]quote,0)

var quotes []quote = []quote{
	{"learn to lead", "Saividya campus"},
	{"you can totally do this", "unknown"},
	{"you are great", "unkown"},
}

func main() {
	rand.Seed(time.Now().Unix())

	api := echo.New()
	api.GET("/quotes", getQuotes)
	api.GET("/quotes/random", getRandomQuotes)

	port := os.Getenv("PORT")
	if port == "" {
		port = "32445"
	}
	api.Start(":" + port) // use export PORT=***** before running it on undefined port number
	// also can use set PORT=****
}

func getQuotes(c echo.Context) error {
	return c.JSON(http.StatusOK, quotes)
}

func getRandomQuotes(c echo.Context) error {
	index := rand.Intn(len(quotes))

	return c.JSON(http.StatusOK, quotes[index])
}
