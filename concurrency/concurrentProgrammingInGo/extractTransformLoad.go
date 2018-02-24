package concurrentProgrammingInGo

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	productListLocation = "./resources/productList.txt"
)

func extractOrders(ch chan *order) {
	f, _ := os.Open("./resources/orders.txt")
	defer f.Close()

	reader := csv.NewReader(f)
	for record, err := reader.Read(); err == nil; record, err = reader.Read() {
		ord := new(order)
		ord.customerNumber, _ = strconv.Atoi(record[0])
		ord.partNumber = record[1]
		ord.quantity, _ = strconv.Atoi(record[2])
		ch <- ord
	}
	close(ch)
}

func transformOrders(extractCh, transCh chan *order) {
	f, _ := os.Open(productListLocation)
	defer f.Close()
	r := csv.NewReader(f)
	records, _ := r.ReadAll()
	productList := make(map[string]*product)
	for _, record := range records {
		prod := new(product)
		prod.partNumber = record[0]
		prod.unitCost, _ = strconv.ParseFloat(record[1], 64)
		prod.unitPrice, _ = strconv.ParseFloat(record[2], 64)
		productList[prod.partNumber] = prod
	}
	wg := new(sync.WaitGroup)
	for o := range extractCh {
		wg.Add(1)
		go func(o *order) {
			time.Sleep(3 * time.Millisecond) // simulate webservice call
			o.unitCost = productList[o.partNumber].unitCost
			o.unitPrice = productList[o.partNumber].unitPrice
			transCh <- o
			wg.Done()
		}(o)
	}
	wg.Wait()
	close(transCh)
}
func loadOrders(transCh chan *order, doneCh chan bool) {
	f, _ := os.Create("./resources/orderDestination.txt")
	defer f.Close()
	fmt.Fprintf(f, "%20s%15s%12s%12s%15s%15s\n", "Part Number", "Quantity", "Unit Cost",
		"Unit Price", "Total Cost", "Total Price")
	wg := new(sync.WaitGroup)
	for o := range transCh {
		wg.Add(1)
		go func(o *order) {
			time.Sleep(1 * time.Millisecond) //delay simulation
			fmt.Fprintf(f, "%20s %15d %12.2f %12.2f %15.2f %15.2f\n", o.partNumber, o.quantity, o.unitCost, o.unitPrice,
				o.unitCost*float64(o.quantity),
				o.unitPrice*float64(o.quantity))
			wg.Done()
		}(o)
	}
	wg.Wait()
	doneCh <- true
}

type product struct {
	partNumber string
	unitCost   float64
	unitPrice  float64
}
type order struct {
	customerNumber int
	partNumber     string
	quantity       int

	unitCost  float64
	unitPrice float64
}
