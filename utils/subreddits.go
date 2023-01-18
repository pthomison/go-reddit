package utils

import (
	"context"

	"github.com/pthomison/errcheck"
	"github.com/pthomison/go-reddit/reddit"
)

func (c *Client) SlurpPopularSubreddits(limit int) map[string]*reddit.Subreddit {

	r := make(map[string]*reddit.Subreddit)
	after := ""
	maxCountPerRequest := 100

	for i := 0; i < limit; i += maxCountPerRequest {

		subreddits := c.PopularSubreddits(&reddit.ListSubredditOptions{
			ListOptions: reddit.ListOptions{
				Limit: RequestLimit(i, limit),
				After: after,
			},
			Sort: "activity",
		})

		after = subreddits[len(subreddits)-1].FullID

		for _, p := range subreddits {
			r[p.ID] = p
		}
	}

	return r
}

func (c *Client) PopularSubreddits(opts *reddit.ListSubredditOptions) []*reddit.Subreddit {
	subreddits, response, err := c.Subreddit.Popular(context.Background(), opts)

	errcheck.Check(err)
	reddit.CheckResponse(response.Response)

	return subreddits
}
