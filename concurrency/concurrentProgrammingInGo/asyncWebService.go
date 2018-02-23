package concurrentProgrammingInGo

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

func AsyncWebService() {
	runtime.GOMAXPROCS(4)
	tickerSymbols := []string{
		"ipcc",
		"aapl",
		"kmpr",
		"so",
	}
	start := time.Now()
	numComplete := 0
	for _, symbol := range tickerSymbols {
		go func(symbol string) {
			resp, error := http.Get("http://dev.markitondemand.com/MODApis/Api/V2/Quote?symbol=" + symbol)
			if error != nil {
				fmt.Println(error)
				return
			}
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			quote := new(quoteResponse)
			xml.Unmarshal(body, &quote)
			fmt.Printf("%s: %.2f\n", quote.Name, quote.LastPrice)
			numComplete++
		}(symbol)
	}
	for numComplete < len(tickerSymbols) {
		time.Sleep(10 * time.Millisecond)
	}
	elapse := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapse)
}

type quoteResponse struct {
	Status           string
	Name             string
	LastPrice        float32
	Change           float32
	ChangePercent    float32
	TimeStamp        string
	MSDate           float32
	MarketCap        int
	Volume           int
	ChangeYTD        float32
	ChangePercentYTD float32
	High             float32
	Low              float32
	Open             float32
}
