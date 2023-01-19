package redditutils

import (
	"context"

	"github.com/pthomison/errcheck"
	"github.com/pthomison/go-reddit/reddit"
)

func (c *Client) SlurpTopPosts(subreddit string, timeframe string, limit int) map[string]*reddit.Post {

	r := make(map[string]*reddit.Post)
	after := ""

	for i := 0; i < limit; i += MAX_LIMIT_PER_REQUEST {

		posts := c.GetTopPosts(subreddit, &reddit.ListPostOptions{
			ListOptions: reddit.ListOptions{
				Limit: RequestLimit(i, limit),
				After: after,
			},
			Time: timeframe,
		})

		after = posts[len(posts)-1].FullID

		for _, p := range posts {
			r[p.ID] = p
		}
	}

	return r
}

func (c *Client) GetTopPosts(subreddit string, opts *reddit.ListPostOptions) []*reddit.Post {

	posts, response, err := c.Subreddit.TopPosts(
		context.Background(),
		subreddit,
		opts)

	errcheck.Check(err)
	reddit.CheckResponse(response.Response)

	return posts
}
