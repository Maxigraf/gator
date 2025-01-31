package command

import (
	"context"
	"fmt"

	"github.com/maxigraf/gator/internal/feed"
)

func handlerAggregate(s *State, cmd Command) error {
	result, err := feed.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")

	if err != nil {
		return fmt.Errorf("failed to get feed: %v", err)
	}

	fmt.Printf("%+v\n", result)

	return nil
}
