package taskexample

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/delay"
	"google.golang.org/appengine/log"
)

var delayedPuppy *delay.Function

func init() {
	delayedPuppy = delay.Func("delayedPuppy", runLater)
}

func runLater(ctx context.Context) {
	log.Infof(ctx, "delayedPuppy ran")
}
