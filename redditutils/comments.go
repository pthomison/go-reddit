package redditutils

import (
	"context"

	"github.com/pthomison/errcheck"
	"github.com/pthomison/go-reddit/reddit"
)

func (c *Client) SlurpUserComments(username string, limit int) []*reddit.Comment {

	r := []*reddit.Comment{}
	after := ""

	for i := 0; i < limit; i += MAX_LIMIT_PER_REQUEST {

		comments := c.GetUserComments(username, &reddit.ListUserOverviewOptions{
			ListOptions: reddit.ListOptions{
				Limit: RequestLimit(i, limit),
				After: after,
			},
		})

		if len(comments) == 0 {
			return r
		}

		after = comments[len(comments)-1].FullID

		r = append(r, comments...)
	}

	return r
}

func (c *Client) GetUserComments(username string, opts *reddit.ListUserOverviewOptions) []*reddit.Comment {
	comments, response, err := c.User.CommentsOf(context.Background(), username, opts)
	errcheck.Check(err)
	reddit.CheckResponse(response.Response)

	return comments
}
