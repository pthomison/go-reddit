package redditutils

import (
	"context"

	"github.com/pthomison/errcheck"
	"github.com/pthomison/go-reddit/reddit"
)

func (c *Client) GetUser(username string) *reddit.User {
	user, response, err := c.User.Get(context.Background(), username)

	errcheck.Check(err)
	reddit.CheckResponse(response.Response)

	return user
}
