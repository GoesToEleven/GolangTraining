package main

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"fmt"
)

func main() {
	
	resp, _ := http.Get("http://dev.markitondemand.com/Api/v2/Quote?symbol=googl")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	
	quote := new(QuoteResponse)
	xml.Unmarshal(body, &quote)
	
	fmt.Printf("%s: %.2f", quote.Name, quote.LastPrice)
	
}


type QuoteResponse struct {
	Status string
	Name string
	LastPrice float32
	Change float32
	ChangePercent float32
	TimeStamp string
	MSDate float32
	MarketCap int
	Volume int
	ChangeYTD float32
	ChangePercentYTD float32
	High float32
	Low float32
	Open float32
}