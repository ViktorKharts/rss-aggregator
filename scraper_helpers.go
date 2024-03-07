package main

import "context"

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
