package context

import (
	"context"
	"github.com/google/go-github/v43/github"
	"strings"
)

type AuthorizedContext struct {
	client  *github.Client
	context context.Context
}

func NewAuthorizedContext(login string, token string) *AuthorizedContext {
	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(login),
		Password: strings.TrimSpace(token),
	}

	return &AuthorizedContext{
		client:  github.NewClient(tp.Client()),
		context: context.Background(),
	}
}
