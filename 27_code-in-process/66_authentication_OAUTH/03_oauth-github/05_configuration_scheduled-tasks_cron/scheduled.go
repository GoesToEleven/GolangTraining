package githubexample

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {
	http.HandleFunc("/scheduled", handleScheduleExample)
}

func handleScheduleExample(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	log.Infof(ctx, "in the scheduler")

	if req.Header.Get("X-Appengine-Cron") != "true" {
		http.Error(res, "Forbidden", http.StatusForbidden)
		return
	}

	log.Infof(ctx, "I was Scheduled!!!")
}
