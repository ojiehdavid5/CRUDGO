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
	http.HandleFunc("/update", itemUpdate)
	http.HandleFunc("/all", itemAll)
	http.HandleFunc("/delete", itemDelete)
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

func itemUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Items updated")

	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var UpdateItem Item
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&UpdateItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	for i, item := range items {

		if item.ID == UpdateItem.ID {
			items[i].ItemName = UpdateItem.ItemName
			response := ItemResponse{
				Message: fmt.Sprintf("%s (ID: %d) updated", UpdateItem.ItemName, UpdateItem.ID),
			}
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(response); err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			fmt.Println(response)
			return
		}

	}
}
func itemAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	if err := json.NewEncoder(w).Encode(items); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println("All items called")

}

func itemDelete(w http.ResponseWriter, r *http.Request) {

	var itemDelete Item;
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&itemDelete); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, existingItem := range items {
		if existingItem.ID == itemDelete.ID {
			items = append(items[:i], items[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	response:=ItemResponse{
		Message: fmt.Sprintf("Item with ID %d has been deleted", itemDelete.ID),}


	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	fmt.Println(response)

fmt.Println("Item Deleted")

}
