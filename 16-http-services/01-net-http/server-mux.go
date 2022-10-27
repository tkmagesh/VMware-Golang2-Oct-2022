package main

import (
	"fmt"
	"net/http"
	"time"
)

//profile middleware
func profile(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler(w, r)
		elapsed := time.Now().Sub(start)
		fmt.Printf("Time taken : %v\n", elapsed)
	}
}

//log middleware
func log(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s - %s\n", r.Method, r.URL)
		handler(w, r)
	}
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Welcome to REST Server"))
}

func productsHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		/*
			1. check if the user wants all the products or 1 product
			2. get the data from the database
			3. serialize the response into json
			4. send the response
		*/
		res.Write([]byte("All the products will be served"))
	case "POST":
		res.Write([]byte("The given new product will be added"))
	}
}

func customersHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("All the customer related requests will be processed"))
}

func main() {
	http.HandleFunc("/", profile(log(indexHandler)))
	http.HandleFunc("/products", profile(log(productsHandler)))
	http.HandleFunc("/customers", profile(log(customersHandler)))
	http.ListenAndServe(":8080", nil)
}

/*
	GET - http://localhost:8080/products
	GET - http://localhost:8080/products/P-101
	POST - http://localhost:8080/products

	GET - http://localhost:8080/customers
	POST - http://localhost:8080/customers
*/
