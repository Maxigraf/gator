package command

import (
	"context"
	"fmt"
	"strconv"

	"github.com/maxigraf/gator/internal/database"
)

func handlerBrowse(s *State, cmd Command, user database.User) error {
	var err error

	limit := 2

	if len(cmd.Args) > 0 {
		limit, err = strconv.Atoi(cmd.Args[0])

		if err != nil {
			limit = 2
		}
	}

	params := database.GetPostsForUserParams{
		ID:    user.ID,
		Limit: int32(limit),
	}

	posts, err := s.Database.GetPostsForUser(context.Background(), params)

	if err != nil {
		return fmt.Errorf("failed to fetch posts: %v", err)
	}

	fmt.Printf("Latest %d posts:\n", limit)

	for _, post := range posts {
		fmt.Printf("[%s]\n", post.PublishedAt)
		fmt.Printf("%s\n", post.Title)
		if post.Description.Valid {
			fmt.Printf("##############################\n%s\n##############################\n", post.Description.String)
		}
		fmt.Println()
	}

	return nil
}
