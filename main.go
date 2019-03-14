package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/stock-check", GetStockPrice).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetStockPrice(w http.ResponseWriter, r *http.Request) {}

//https://www.alphavantage.co/query?apikey=demo&function=TIME_SERIES_DAILY_ADJUSTED&symbol=MSFT
//C227WD9W3LUVKVV9
