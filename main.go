package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/savaki/jq"
)

/*
type Data struct {
	//MSFT  data=[110.56, 111.25, 115.78], average=112.50
	Name    string `json:"firstname,omitempty"`
	data    array  `json:"id,omitempty"`
	Average string `json:"id,omitempty"`
}
*/

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetStockPrice).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetStockPrice(w http.ResponseWriter, r *http.Request) {
	//url := "https://www.alphavantage.co/query?apikey=C227WD9W3LUVKVV9&function=TIME_SERIES_DAILY_ADJUSTED&symbol=MSFT"
	fmt.Println("SYMBOL:", os.Getenv("SYMBOL"))
	fmt.Println("NDAYS:", os.Getenv("NDAYS"))
	fmt.Println("APIKEY:", os.Getenv("APIKEY"))

	apiKey := os.Getenv("APIKEY")
	symbol := os.Getenv("SYMBOL")
	//ndays := os.Getenv("NDAYS")

	url := "https://www.alphavantage.co/query?apikey=%s&function=TIME_SERIES_DAILY_ADJUSTED&symbol=%s"
	resp, _ := http.Get(fmt.Sprintf(url, apiKey, symbol))
	bytes, _ := ioutil.ReadAll(resp.Body)

	//fmt.Println("HTML:\n\n", string(bytes))

	resp.Body.Close()

	//myData := "MSFT"

	op, _ := jq.Parse(".")
	value, _ := op.Apply(bytes)
	//fmt.Println(string(value))
	json.NewEncoder(w).Encode(string(value))

}

//https://www.alphavantage.co/query?apikey=demo&function=TIME_SERIES_DAILY_ADJUSTED&symbol=MSFT
//C227WD9W3LUVKVV9
