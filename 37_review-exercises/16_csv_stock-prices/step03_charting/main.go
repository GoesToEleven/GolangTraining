package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type record struct {
	Date string
	Open float64
}

func makeRecord(row []string) record {
	open, _ := strconv.ParseFloat(row[1], 64)
	return record{
		Date: row[0],
		Open: open,
	}
}

func main() {
	f, err := os.Open("table.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rdr := csv.NewReader(f)
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(`<!DOCTYPE html>
<head>
<script src="http://code.jquery.com/jquery-1.11.3.min.js"></script>
<script src="http://code.highcharts.com/highcharts.js"></script>
<script src="http://code.highcharts.com/modules/exporting.js"></script>
</head>
<body>
<div id="container" style="min-width: 310px; height: 400px; margin: 0 auto"></div>



  <table>
    <thead>
      <tr>
        <th>Date</th>
        <th>Open</th>
      </tr>
    </thead>
    <tbody>
    `)

	openValues := []string{}

	for i, row := range rows {
		if i == 0 {
			continue
		}
		record := makeRecord(row)
		fmt.Println(`
      <tr>
        <td>` + record.Date + `</td>
        <td>` + fmt.Sprintf("%.2f", record.Open) + `</td>
      </tr>
      `)

		openValues = append(openValues, fmt.Sprintf("%.2f", record.Open))

	}

	fmt.Println(`
    </tbody>
  </table>

  <script>
  $(function () {
      $('#container').highcharts({
          title: {
              text: 'Monthly Average Temperature',
              x: -20 //center
          },
          subtitle: {
              text: 'Source: WorldClimate.com',
              x: -20
          },
          xAxis: {
              categories: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun',
                  'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']
          },
          yAxis: {
              title: {
                  text: 'Temperature (°C)'
              },
              plotLines: [{
                  value: 0,
                  width: 1,
                  color: '#808080'
              }]
          },
          tooltip: {
              valueSuffix: '°C'
          },
          legend: {
              layout: 'vertical',
              align: 'right',
              verticalAlign: 'middle',
              borderWidth: 0
          },
          series: [{
              name: 'Tokyo',
              data: [

` + strings.Join(openValues, ",") + `

              ]
          }]
      });
  });
  </script>
</body>
    `)
}
