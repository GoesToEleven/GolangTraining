package csvexample

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"google.golang.org/appengine"
)

func init() {
	http.HandleFunc("/", handleInput)
	http.HandleFunc("/madoff", handleOutput)
}

func handleInput(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, `<!DOCTYPE html>
<html>
  <head></head>
  <body>
    <form method="GET" action="/madoff">
      <label>Symbol #1:
        <input type="text" name="symbol1">
      </label>
      <label>Symbol #2:
        <input type="text" name="symbol2">
      </label>
      <input type="submit">
    </form>
  </body>
</html>`)
}

func handleOutput(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	symbol1, symbol2 := req.FormValue("symbol1"), req.FormValue("symbol2")

	data1, err := getData(ctx, symbol1)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	data2, err := getData(ctx, symbol2)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	data1json, _ := json.Marshal(data1)
	data2json, _ := json.Marshal(data2)

	c := correlation(data1, data2)

	io.WriteString(res, `<!DOCTYPE html>
<html>
  <head>
    <script src="//cdnjs.cloudflare.com/ajax/libs/Chart.js/1.0.2/Chart.js"></script>
  </head>
  <body>
    <canvas id="myChart" width="400" height="400"></canvas>


    SYMBOL 1: `+symbol1+`<br>
    SYMBOL 2: `+symbol2+`<br>
    CORRELATION: <strong>`+fmt.Sprintf("%.4f", c)+`</strong>

    <script>
var ctx = document.getElementById("myChart").getContext("2d");
var myNewChart = new Chart(ctx).Line({
  labels: [
  "January","February"
  ],
  datasets: [
    {
      label: "`+symbol1+`",
      fillColor: "rgba(220,220,220,0.2)",
      strokeColor: "rgba(220,220,220,1)",
      pointColor: "rgba(220,220,220,1)",
      pointStrokeColor: "#fff",
      pointHighlightFill: "#fff",
      pointHighlightStroke: "rgba(220,220,220,1)",
      data: `+string(data1json)+`
    },
    {
      label: "`+symbol2+`",
      fillColor: "rgba(151,187,205,0.2)",
      strokeColor: "rgba(151,187,205,1)",
      pointColor: "rgba(151,187,205,1)",
      pointStrokeColor: "#fff",
      pointHighlightFill: "#fff",
      pointHighlightStroke: "rgba(151,187,205,1)",
      data: `+string(data2json)+`
    }
  ]
}, null);
    </script>
  </body>
</html>
`)

}
