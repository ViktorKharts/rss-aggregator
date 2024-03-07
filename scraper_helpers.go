package main

import (
	"context"
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
