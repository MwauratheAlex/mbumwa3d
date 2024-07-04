package handlers

import (
	"fmt"
	"net/http"
)

func GetActiveOrders(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Active")
	return nil
}
func GetAvailableOrders(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Available")

	return nil
}
func GetCompletedOrders(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Completed")

	return nil
}
