package handlers

import (
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/suda/go-swagger-example/restapi/operations/tweets"
)

// Tweets handlers
type Tweets struct {
}

// FindTweets will list all tweets
func (t *Tweets) FindTweets(params tweets.TweetsListParams) middleware.Responder {
	return middleware.NotImplemented("I'll implement listing tweets later ;)")
}
