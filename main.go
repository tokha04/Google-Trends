package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title            string        `xml:"title"`
	GoogleTrendsList []GoogleTrend `xml:"item"`
}
type GoogleTrend struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Traffic   string `xml:"approx_traffic"`
	NewsItems []News `xml:"news_item"`
}

type News struct {
	Headline     string `xml:"news_item_title"`
	HeadlineLink string `xml:"news_item_url"`
}

func main() {
	var r RSS

	data := readGoogleTrends()

	err := xml.Unmarshal(data, &r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("\nBelow are all the Google Search Trends For Today in Turkey!")
	fmt.Println("-------------------------------------------------")

	for i := range r.Channel.GoogleTrendsList {
		rank := i + 1
		fmt.Println("#", rank)
		fmt.Println("Search Term:", r.Channel.GoogleTrendsList[i].Title)
		fmt.Println("Link to the trend:", r.Channel.GoogleTrendsList[i].Link)
		fmt.Println("Headline:", r.Channel.GoogleTrendsList[i].NewsItems[0].Headline)
		fmt.Println("Link to the article:", r.Channel.GoogleTrendsList[i].NewsItems[0].HeadlineLink)
		fmt.Println("-------------------------------------------------")
	}
}

func readGoogleTrends() []byte {
	resp := getGoogleTrends()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return data
}

func getGoogleTrends() *http.Response {
	resp, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=TR")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return resp
}
