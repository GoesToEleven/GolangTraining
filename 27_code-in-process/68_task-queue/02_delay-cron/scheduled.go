package taskexample

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {
	http.HandleFunc("/schedule-example", handleScheduleExample)
}

func handleScheduleExample(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("X-Appengine-Cron") != "true" {
		http.Error(res, "Forbidden", http.StatusForbidden)
		return
	}
	ctx := appengine.NewContext(req)
	log.Infof(ctx, "I was Scheduled!!!")
	delayedPuppy.Call(ctx)
}
