package main

import (
	"fmt"
	"net/http"
	"strconv"
	"log"
)

type dollars float64

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f",d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such file: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	itemPrice := r.URL.Query().Get("price")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such file: %q\n", item)
		return
	}
	value, _ := strconv.ParseFloat(itemPrice, 64)
	if dollars(value) != price {
		db[item] = dollars(value)
	}
	fmt.Fprintf(w, "%s\n", db[item])
}

func change(s map[string] string, key, value string) {
	s[key] = value
}

func main() {

	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	//mux.Handle("/list", http.HandlerFunc(db.list))
	//mux.Handle("/price", http.HandlerFunc(db.price))
	//HandleFunc 包装好的HandlerFunc
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/update", db.update)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
