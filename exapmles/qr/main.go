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
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}).Methods("GET")

	r.HandleFunc("/qr", func(w http.ResponseWriter, r *http.Request) {
		req := paypay.CreateQRRequest{
			//MerchantPaymentID: "1b5b82d675834d418c1fbf27544bcxyz", //加盟店発番のユニークな決済トランザクションID
			MerchantPaymentID: "2b5b82d675834d418c1fbf27544bcxyz", //加盟店発番のユニークな決済トランザクションID

			CodeType:          "ORDER_QR",
			Amount: paypay.Amount{
				Amount:   10,
				Currency: "JPY",
			},
			OrderItems: []paypay.OrderItem{
				paypay.OrderItem{
					Name:      "product1",
					Category:  "cake",
					Quantity:  1,
					ProductID: "1",
					UnitPrice: paypay.UnitPrice{
						Amount:   10,
						Currency: "JPY",
					},
				},
			},
			RedirectURL:  "http://localhost:8000/",
			RedirectType: "WEB_LINK",
		}

		resp, err := c.CreateQR(req)
		if err != nil {
			log.Printf("ERROR: %v", err)
			return
		}
		res, _ := json.Marshal(resp)
		w.Write(res)
	}).Methods("POST")

	r.HandleFunc("/qr/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		resp, err := c.DeleteQR(vars["id"])
		if err != nil {
			log.Printf("ERROR: %v", err)
			return
		}
		res, _ := json.Marshal(resp)
		w.Write(res)
	}).Methods("DELETE")

	fmt.Println("Listen to :8081")
	http.ListenAndServe(":8081", r)
}
