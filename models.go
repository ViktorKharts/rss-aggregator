package main

import (
	"time"

	"github.com/ViktorKharts/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey	  string    `json:"api_key"`
}

type Feed struct {
	ID            uuid.UUID `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Name          string `json:"name"`
	Url           string `json:"url"`
	UserID        uuid.UUID `json:"user_id"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
}

type Post struct {
	ID		uuid.UUID `json:"id"`
	CreatedAt	time.Time `json:"created_at"`
	UpdatedAt	time.Time `json:"updated_at"`
	Title		string `json:"title"`
	Url		string `json:"url"`
	Description	*string	`json:"description"`	
	PublishedAt	*time.Time `json:"published_at"`
	FeedID		uuid.UUID `json:"feed_id"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:	   user.ApiKey,
	}
}

func databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:	       feed.ID, 
		CreatedAt:     feed.CreatedAt, 
		UpdatedAt:     feed.UpdatedAt, 
		Name:          feed.Name, 
		Url:           feed.Url, 
		UserID:        feed.UserID, 
		LastFetchedAt: &feed.LastFetchedAt.Time, 
	}
}

func databasePostToPost(post database.Post) Post {
	return Post{
		ID: post.ID,
		Title: post.Title,	
		Url: post.Url, 	
		Description: &post.Description.String,		
		FeedID: post.FeedID,		
	}
}

