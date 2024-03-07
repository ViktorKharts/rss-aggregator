package main

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
	"time"

	"github.com/ViktorKharts/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (c *apiConfig) getNextFeedToFetch(ctx context.Context) ([]Feed, error) {
	dbFeeds, err := c.DB.GetFeedsByLastFetchedAt(ctx)
	if err != nil {
		return []Feed{}, err 
	} 

	feeds := []Feed{}
	for _, f := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(f))
	}

	return feeds, nil
}

func (c *apiConfig) markFeedFetched(ctx context.Context, feedId uuid.UUID) error {
	err := c.DB.UpdateFeedLastFetchedAt(ctx, database.UpdateFeedLastFetchedAtParams{
		ID: feedId,
		UpdatedAt: time.Now(),
	})
	return err 
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchDataFromFeed(url string) (interface{}, error) {	
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rssFeed := RSSFeed{}
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return nil, err
	}

	return rssFeed, nil  
}
