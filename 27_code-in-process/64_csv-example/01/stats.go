package csvexample

import (
	"fmt"

	"google.golang.org/appengine/urlfetch"

	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
	"io/ioutil"
)

func getData(ctx context.Context, symbol string) {
	url := fmt.Sprintf(
		"http://real-chart.finance.yahoo.com/table.csv"+
			"?s=%s"+
			"&d=6"+
			"&e=1"+
			"&f=2015"+
			"&g=m"+
			"&a=6"+
			"&b=1"+
			"&c=1986"+
			"&ignore=.csv",
		symbol,
	)
	client := urlfetch.Client(ctx)

	result, err := client.Get(url)
	if err != nil {
		return
	}
	defer result.Body.Close()

	bs, _ := ioutil.ReadAll(result.Body)
	log.Infof(ctx, "RESULT: %v", string(bs))
}
