package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	paypay "paypay-sdk-go"
)

func main() {
	c, _ := paypay.NewClient("api-key", "api-secret", 00000000000, false, nil)

	r := mux.NewRouter()
	r.HandleFunc("/payment/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		resp, err := c.GetPaymentDetail(vars["id"])
		if err != nil {
			log.Printf("ERROR: %v", err)
			return
		}
		payment, _ := json.Marshal(resp)
		w.Write(payment)
	}).Methods("GET")

	r.HandleFunc("/payment/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		resp, err := c.CancelPayment(vars["id"])
		if err != nil {
			log.Printf("ERROR: %v", err)
		}
		res, _ := json.Marshal(resp)
		w.Write(res)
	}).Methods("DELETE")

	fmt.Println("Listen to :8080")
	http.ListenAndServe(":8080", r)
}