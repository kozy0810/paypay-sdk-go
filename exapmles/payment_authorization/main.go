package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	paypay "paypay-sdk-go"
	"strconv"
	"time"
)

func main() {
	c, _ := paypay.NewClient("api-key", "api-secret", 00000000000, false, nil)

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}).Methods("GET")

	r.HandleFunc("/capture", func(w http.ResponseWriter, r *http.Request) {
		epoch := strconv.FormatInt(time.Now().Unix(), 10)
		e, _ := strconv.Atoi(epoch)
		req := paypay.CaptureRequest{
			MerchantPaymentID: "3b5b82d675834d418c1fbf27445ac6cb", //加盟店発番のユニークな決済トランザクションID
			MerchantCaptureID: "3b5b82d675834d418c1fbf27445ac6cb", //加盟店発番のユニークなキャプチャトランザクションID
			Amount: paypay.Amount{
				Amount:   1,
				Currency: "JPY",
			},
			RequestedAt:      e, // エポックタイムスタンプ（秒単位）
			OrderDescription: "test",
		}
		resp, err := c.Capture(req)
		if err != nil {
			log.Printf("ERROR: %v", err)
			return
		}
		res, _ := json.Marshal(resp)
		w.Write(res)
	}).Methods("POST")

	r.HandleFunc("/revert", func(w http.ResponseWriter, r *http.Request) {
		epoch := strconv.FormatInt(time.Now().Unix(), 10)
		e, _ := strconv.Atoi(epoch)
		req := paypay.RevertRequest{
			MerchantRevertID: "3b5b82d675834d418c1fbf27445ac6cb",
			PaymentID:        "03450018268289114112",
			RequestedAt:      e,
		}

		resp, err := c.Revert(req)
		if err != nil {
			log.Printf("ERROR: %v", err)
			return
		}
		res, _ := json.Marshal(resp)
		w.Write(res)
	}).Methods("POST")

	fmt.Println("Listen to :8080")
	http.ListenAndServe(":8080", r)
}
