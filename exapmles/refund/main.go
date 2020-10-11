package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
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

	r.HandleFunc("/refund/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		resp, err := c.GetRefund(vars["id"])
		if err != nil {
			log.Printf("ERROR: %v", err)
			return
		}
		b, _ := json.Marshal(resp)
		w.Write(b)
	}).Methods("GET")

	r.HandleFunc("/refund", func(w http.ResponseWriter, r *http.Request) {
		epoch := strconv.FormatInt(time.Now().Unix(), 10)
		e, _ := strconv.Atoi(epoch)
		reqBody, _ := ioutil.ReadAll(r.Body)
		req := paypay.RefundRequest{
			MerchantRefundID: "", //加盟店発番のユニークな取消トランザクションID
			PaymentID:        "", //PayPay発番の決済トランザクションID
			Amount: paypay.Amount{
				Amount:   1,
				Currency: "JPY",
			},
			RequestedAt: e,
			Reason:      "test",
		}
		if err := json.Unmarshal(reqBody, &req); err != nil {
			log.Printf("ERROR: %v", err)
			return
		}
		resp, err := c.Refund(req)
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
