package csvexample

import (
	"encoding/csv"
	"fmt"
	"math"
	"strconv"

	"google.golang.org/appengine/urlfetch"

	"golang.org/x/net/context"
)

func correlation(xs, ys []float64) float64 {
	return covariance(xs, ys) / (standardDeviation(xs) * standardDeviation(ys))
}

func covariance(x, y []float64) float64 {
	if len(x) != len(y) {
		panic("Vector lengths must be the same")
	}

	n := len(x)
	sum, xsum, xmean, ysum, ymean := 0.0, 0.0, 0.0, 0.0, 0.0

	for i := 0; i < n; i++ {
		xsum += x[i]
		ysum += y[i]
	}
	xmean = xsum / float64(n)
	ymean = ysum / float64(n)

	for i := 0; i < n; i++ {
		sum += (x[i] - xmean) * (y[i] - ymean)
	}

	return sum / float64(n-1)
}

func getData(ctx context.Context, symbol string) ([]float64, error) {
	url := fmt.Sprintf(
		"http://real-chart.finance.yahoo.com/table.csv"+
			"?s=%s"+
			"&d=6"+
			"&e=1"+
			"&f=2015"+
			"&g=m"+
			"&a=6"+
			"&b=1"+
			"&c=2014"+
			"&ignore=.csv",
		symbol,
	)
	client := urlfetch.Client(ctx)

	result, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	// parse the csv file
	reader := csv.NewReader(result.Body)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// convert the rows to floats
	const (
		dateColumn  = 0
		closeColumn = 4
	)
	var data []float64
	for i, row := range records {
		if i == 0 || len(row) < 5 {
			continue
		}
		val, _ := strconv.ParseFloat(row[closeColumn], 64)
		data = append(data, val)
	}
	return relativize(data), nil
}

func standardDeviation(xs []float64) float64 {
	return math.Sqrt(variance(xs))
}

func variance(vector []float64) float64 {
	n := 0.0
	mean := 0.0
	S := 0.0
	delta := 0.0

	for _, v := range vector {
		n++
		delta = v - mean
		mean = mean + (delta / n)
		S += delta * (v - mean)
	}

	return S / (n - 1)
}

func relativize(data []float64) []float64 {
	nv := make([]float64, len(data)-1)
	for i := 1; i < len(data); i++ {
		nv[i-1] = (data[i] - data[i-1]) / data[i-1]
	}
	return nv
}
