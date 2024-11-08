package rss

import (
	"errors"
	"github.com/mmcdole/gofeed"
)

// Post Публикация, получаемая из RSS.
type Post struct {
	ID      string `json:"guid,omitempty"`        // номер записи
	Title   string `json:"title,omitempty"`       // заголовок публикации
	Content string `json:"description,omitempty"` // содержание публикации
	PubTime int64  `json:"pubDate,omitempty"`     // время публикации
	Link    string `json:"link,omitempty"`        // ссылка на источник
}

func getRSS(url string) ([]Post, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, err
	}

	if len(feed.Items) == 0 {
		return nil, errors.New("feed is empty")
	}

	var posts []Post
	for _, item := range feed.Items {
		posts = append(posts, Post{
			ID:      item.GUID,
			Title:   item.Title,
			Content: item.Description,
			PubTime: item.PublishedParsed.Unix(),
			Link:    item.Link,
		})
	}
	return posts, nil
}
