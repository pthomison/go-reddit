package redditutils

import (
	"context"

	"github.com/pthomison/errcheck"
	"github.com/pthomison/go-reddit/reddit"
)

func GetUser(client *reddit.Client, username string) *reddit.User {
	user, response, err := client.User.Get(context.Background(), username)

	errcheck.Check(err)
	reddit.CheckResponse(response.Response)

	return user
}
