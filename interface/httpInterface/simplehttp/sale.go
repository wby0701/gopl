package main

import (
	"fmt"
	"log"
	"net/http"
)


type dollars float64

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f",d)
}

type database map[string]dollars

//func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request){
//	for item , price := range db {
//		fmt.Fprintf(w, "%s, %s\n", item, price)
//	}
//}

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}

	case "/price":
		item := r.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)

	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s\n", r.URL)
	}
}

func main() {
	db := database{"shoes" : 50,
		"sock" : 5,
		"111" : 6}
	log.Fatal(http.ListenAndServe("127.0.0.1:8000",db))
}