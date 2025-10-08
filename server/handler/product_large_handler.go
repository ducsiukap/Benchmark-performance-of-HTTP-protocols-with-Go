package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// request product small
func ProductLargeHandler(w http.ResponseWriter, r *http.Request) {

	// log http version
	log.Printf("HTTP version:%s", r.Proto)

	// w.Header().Add("Content-Type", "application/json")

	switch r.Method {
	// GET
	case http.MethodGet:
		getProdLarge(w, r)

	// POST
	case http.MethodPost:
		postProdLarge(w, r)

	//Other method
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s is not allowed!", r.Method)
	}
}

// @GET
func getProdLarge(w http.ResponseWriter, r *http.Request) {

	query := fmt.Sprintf("SELECT id, name, price FROM products_large")
	// execute query
	rows, err := conn.Query(query)
	if err != nil {
		http.Error(w, "Error when fetching data!", http.StatusInternalServerError)
		return
	}
	// đóng row
	defer rows.Close()

	// chuyen thanh product
	var products []Product
	for rows.Next() {
		var p Product
		// quét hàng hiện tại vào p
		err := rows.Scan(&p.id, &p.name, &p.price)
		if err != nil {
			continue
		}
		//
		products = append(products, p)
	}

	// send data to client
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprintf(w, "{\"products\": [")
	for i, p := range products {
		if i > 0 {
			fmt.Fprintf(w, ", ")
		}
		fmt.Fprintf(w, "{\"id\": %d, \"name\": %s, \"price\": %.2f}", p.id, p.name, p.price)
	}
	fmt.Fprintf(w, "]}")
}

// @POST
func postProdLarge(w http.ResponseWriter, r *http.Request) {

	var products []Product

	// decode request
	err := json.NewDecoder(r.Body).Decode(&products)
	if err != nil {
		http.Error(w, "json parse error: "+err.Error(), http.StatusBadRequest)
		return
	}

	// prepare statement
	pstm, err := conn.Prepare("INSERT INTO Products_large(name, price) values (?, ?)")
	if err != nil {
		http.Error(w, "DB pstm error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer pstm.Close()

	// update
	for _, product := range products {
		// execute pstm
		_, err = pstm.Exec(product.name, product.price)
		if err != nil {
			http.Error(w, "DB insert error: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	//response
	fmt.Fprintf(w, "Insert %d products into PRODUCTS_SMALL successfully!", len(products))
}
