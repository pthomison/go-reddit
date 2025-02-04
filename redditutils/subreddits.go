package redditutils

import (
	"context"

	"github.com/pthomison/errcheck"
	"github.com/pthomison/go-reddit/reddit"
)

func (c *Client) SlurpPopularSubreddits(limit int) map[string]*reddit.Subreddit {

	r := make(map[string]*reddit.Subreddit)
	after := ""

	for i := 0; i < limit; i += MAX_LIMIT_PER_REQUEST {

		subreddits := c.GetPopularSubreddits(&reddit.ListSubredditOptions{
			ListOptions: reddit.ListOptions{
				Limit: RequestLimit(i, limit),
				After: after,
			},
			Sort: "activity",
		})

		if len(subreddits) == 0 {
			return r
		}

		after = subreddits[len(subreddits)-1].FullID

		for _, p := range subreddits {
			r[p.ID] = p
		}
	}

	return r
}

func (c *Client) GetPopularSubreddits(opts *reddit.ListSubredditOptions) []*reddit.Subreddit {
	subreddits, response, err := c.Subreddit.Popular(context.Background(), opts)

	errcheck.Check(err)
	reddit.CheckResponse(response.Response)

	return subreddits
}
