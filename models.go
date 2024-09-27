package main

import (
	"time"

	"github.com/Vkanhan/go-aggregator/internal/database"
	"github.com/google/uuid"
)

// User represents a user with ID, timestamps, name, and API key.
type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func databaseUserToUse(dbUser database.User) *User {
	return &User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

// Feed represents a feed with ID, timestamps, name, URL, and associated user ID.
type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAT time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAT: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		URL:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func databaseFeedstoReturn(dbFeeds []database.Feed) []Feed {
	feeds := make([]Feed, 0, len(dbFeeds))

	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}

// FeedFollow represents a relationship between a user and a feed they follow.
type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(dBFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dBFeedFollow.ID,
		CreatedAt: dBFeedFollow.CreatedAt,
		UpdatedAt: dBFeedFollow.UpdatedAt,
		UserID:    dBFeedFollow.UserID,
		FeedID:    dBFeedFollow.FeedID,
	}

}

func databaseFeedFollowsToFeedFollows(dBFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := make([]FeedFollow, 0, len(dBFeedFollows))

	for _, dBFeedFollow := range dBFeedFollows {
		feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(dBFeedFollow))
	}
	return feedFollows
}
