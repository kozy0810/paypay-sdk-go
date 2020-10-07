package paypay

import (
	"encoding/json"
	"log"
)

type (
	CaptureRequest struct {
		MerchantPaymentID string `json:"merchantPaymentId"`
		Amount            Amount `json:"amount"`
		MerchantCaptureID string `json:"merchantCaptureId"`
		RequestedAt       int    `json:"requestedAt"`
		OrderDescription  string `json:"orderDescription"`
	}

	CaptureData struct {
		PaymentID           string      `json:"paymentId"`
		Status              string      `json:"status"`
		AcceptedAt          int         `json:"acceptedAt"`
		Refunds             Refunds     `json:"refunds"`
		Captures            Captures    `json:"captures"`
		MerchantPaymentID   string      `json:"merchantPaymentId"`
		UserAuthorizationID string      `json:"userAuthorizationId"`
		Amount              Amount      `json:"amount"`
		RequestedAt         int         `json:"requestedAt"`
		ExpiresAt           interface{} `json:"expiresAt"`
		StoreID             string      `json:"storeId"`
		TerminalID          string      `json:"terminalId"`
		OrderReceiptNumber  string      `json:"orderReceiptNumber"`
		OrderDescription    string      `json:"orderDescription"`
		OrderItems          []OrderItem `json:"orderItems"`
		Metadata            Metadata    `json:"metadata"`
		AssumeMerchant      string      `json:"assumeMerchant"`
	}

	CaptureResponse struct {
		ResultInfo ResultInfo  `json:"resultInfo"`
		Data       CaptureData `json:"data,omitempty"`
	}
)

func (c *Client) Capture(captureRequest CaptureRequest) (*CaptureResponse, error) {
	req, err := json.Marshal(captureRequest)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}
	resp, err := c.doRequest("POST", "/v2/payments/capture", map[string]string{}, req)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}

	var captureResponse CaptureResponse
	if err := json.Unmarshal(resp, &captureResponse); err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}
	return &captureResponse, nil
}

type (
	RevertRequest struct {
		MerchantRevertID string `json:"merchantRevertId"`
		PaymentID        string `json:"paymentId"`
		RequestedAt      int    `json:"requestedAt"`
		Reason           string `json:"reason"`
	}

	RevertData struct {
		Status      string `json:"status"`
		AcceptedAt  int    `json:"acceptedAt"`
		PaymentID   string `json:"paymentId"`
		RequestedAt int    `json:"requestedAt"`
		Reason      string `json:"reason"`
	}

	RevertResponse struct {
		ResultInfo ResultInfo `json:"resultInfo"`
		Data       RevertData `json:"data"`
	}
)

func (c *Client) Revert(revertRequest RevertRequest) (*RevertResponse, error) {
	req, err := json.Marshal(revertRequest)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest("POST", "/v2/payments/preauthorize/revert", map[string]string{}, req)
	if err != nil {
		return nil, err
	}

	var revertResponse RevertResponse
	if err := json.Unmarshal(resp, &revertResponse); err != nil {
		return nil, err
	}
	return &revertResponse, nil
}
