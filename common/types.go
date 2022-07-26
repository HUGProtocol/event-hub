package common

import "time"

type Account struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	TwitterId      string    `json:"twitter_id"`
	ContentAddress string    `json:"content_address"`
	AssetAddress   string    `json:"asset_address"`
	ProfileUrl     string    `json:"profile_url"`
	Publisher      bool      `json:"publisher"`
	CreateTime     time.Time `json:"create_time"`
	UpdateTime     time.Time `json:"update_time"`
}

type Event struct {
	EventId       string        `json:"event_id"`
	TwitterId     string        `json:"twitter_id"`
	Publisher     string        `json:"publisher"`
	CreateTime    time.Time     `json:"create_time"`
	PublicMetrics PublicMetrics `json:"public_metrics"`
}

type PublicMetrics struct {
	RetweetCount int `json:"retweet_count"`
	ReplyCount   int `json:"reply_count"`
	LikeCount    int `json:"like_count"`
	QuoteCount   int `json:"quote_count"`
}

type Reply struct {
	EventId string  `json:"event_id"`
	Account Account `json:"account"`
	Text    string  `json:"text"`
}
