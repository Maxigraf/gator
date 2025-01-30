package feed

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, http.NoBody)

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("User-Agent", "gator")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	defer res.Body.Close()

	dataXml, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("failed to read body: %v", err)
	}

	result := &RSSFeed{}

	err = xml.Unmarshal(dataXml, result)

	if err != nil {
		return nil, fmt.Errorf("failed to deserialize feed: %v", err)
	}

	result.Channel.Title = html.UnescapeString(result.Channel.Title)
	result.Channel.Description = html.UnescapeString(result.Channel.Description)

	for i, item := range result.Channel.Item {
		result.Channel.Item[i].Title = html.UnescapeString(item.Title)
		result.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}

	return result, nil
}
