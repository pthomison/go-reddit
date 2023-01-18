package redditutils

import (
	"github.com/pthomison/awsutils"
	"github.com/pthomison/errcheck"
	"github.com/pthomison/go-reddit/reddit"
)

func CreateClient(creds reddit.Credentials) Client {
	client, err := reddit.NewClient(creds)
	errcheck.Check(err)
	return Client{client}
}

func RequestCredentials(username, pw_loc, id_loc, secret_loc string) reddit.Credentials {
	region := "us-east-2"

	password, err := awsutils.AWSGetParameter(pw_loc, region)
	errcheck.Check(err)

	id, err := awsutils.AWSGetParameter(id_loc, region)
	errcheck.Check(err)

	secret, err := awsutils.AWSGetParameter(secret_loc, region)
	errcheck.Check(err)

	return reddit.Credentials{
		ID:       id,
		Secret:   secret,
		Username: username,
		Password: password,
	}
}
