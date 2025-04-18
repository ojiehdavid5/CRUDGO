package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Item struct {
	ID       int    `json:"id"`
	ItemName string `json:"item_name"`
}

type ItemResponse struct {
	Message string `json:"message"`
}

var items []Item
var nextID = 1

func main() {
	fmt.Println("Hello, World!")
	port := ":8080"

	http.HandleFunc("/", itemHandler)
	http.HandleFunc("/create", itemCreate)
	// http.HandleFunc("/update", itemUpdate)
	// http.HandleFunc("/all", itemAll)
	// http.HandleFunc("/delete", itemDelete)
	http.ListenAndServe(port, nil)
}
func itemHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(items)

}
func itemCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return

	}
	var newItem Item

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newItem.ID = nextID
	nextID++
	items = append(items, newItem)

	response := ItemResponse{
		Message: fmt.Sprintf("%s (ID: %d) created", newItem.ItemName, newItem.ID),
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	fmt.Println(response)

	fmt.Println("Item Created")

}

// func itemUpdate(w http.ResponseWriter, r *http.Request) {
// fmt.Println("Items updated");

// }
// func itemAll(w http.ResponseWriter, r *http.Request) {
// fmt.Println("All items called")

// }
// func itemDelete(w http.ResponseWriter, r *http.Request) {
// fmt.Println("Item Deleted")
