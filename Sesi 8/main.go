package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const PORT = ":8080"

/* Example API in Postman
   origin = http://localhost:8080/api/users
   [POST]
   [GET]
   [PUT]
*/

type Product struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Price uint   `json:"price"`
}

type InsertProductRequestDto struct {
	Title string `json:"title"`
	Price uint   `json:"price"`
}

var products []Product = []Product{
	{
		ID:    "P-1",
		Title: "Kemeja Hitam",
		Price: 300000,
	},
	{
		ID:    "P-2",
		Title: "Sepatu",
		Price: 300000,
	},
	{
		ID:    "P-3",
		Title: "Celana Jeans",
		Price: 300000,
	},
	{
		ID:    "P-4",
		Title: "Kaos",
		Price: 300000,
	},
	{
		ID:    "P-5",
		Title: "Botol",
		Price: 300000,
	},
}

func getProducts(w http.ResponseWriter) {
	b, err := json.Marshal(products)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func insertProduct(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	var dto InsertProductRequestDto

	var newProduct := Product{

	}

	fmt.Printf("%+v\n", dto)

	err = json.Unmarshal(body, &dto)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println(string(body))

	// w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("success"))
}

func main() {

	// var p Product = Product{
	// 	ID: "P-1",
	// 	Title: "Kemjea",
	// 	Price: 200,
	// }

	// b, err := json.Marshal(p)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println(string(b))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			getProducts(w)
			return
		}

		if r.Method == "POST" {
			insertProduct(w, r)
			return
		}

		return

		// fmt.Printf("request => %v\n", r)

		// responseMessage := "sukses"

		// w.Write([]byte(responseMessage))
	})

	fmt.Printf("Listening on PORT %s\n", PORT)
	http.ListenAndServe(PORT, nil)

}
