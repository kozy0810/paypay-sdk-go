package paypay

import (
	"encoding/json"
	"fmt"
	"log"
)

type (
	CreateQRRequest struct {
		MerchantPaymentID   string      `json:"merchantPaymentId"`
		Amount              Amount      `json:"amount"`
		OrderDescription    string      `json:"orderDescription"`
		OrderItems          []OrderItem `json:"orderItems"`
		Metadata            Metadata    `json:"metadata,omitempty"`
		CodeType            string      `json:"codeType"`
		StoreInfo           string      `json:"storeInfo"`
		StoreID             string      `json:"storeId"`
		TerminalID          string      `json:"terminalId"`
		RequestedAt         int         `json:"requestedAt"`
		RedirectURL         string      `json:"redirectUrl"`
		RedirectType        string      `json:"redirectType"`
		UserAgent           string      `json:"userAgent"`
		IsAuthorization     bool        `json:"isAuthorization"`
		AuthorizationExpiry interface{} `json:"authorizationExpiry"`
	}

	CreateQRData struct {
		CodeID              string      `json:"codeId,omitempty"`
		URL                 string      `json:"url,omitempty"`
		Deeplink            string      `json:"deeplink,omitempty"`
		ExpiryDate          int         `json:"expiryDate,omitempty"`
		MerchantPaymentID   string      `json:"merchantPaymentId,omitempty"`
		Amount              Amount      `json:"amount,omitempty"`
		OrderDescription    string      `json:"orderDescription,omitempty"`
		OrderItems          []OrderItem `json:"orderItems,omitempty"`
		Metadata            Metadata    `json:"metadata,omitempty"`
		CodeType            string      `json:"codeType,omitempty"`
		StoreInfo           string      `json:"storeInfo,omitempty"`
		StoreID             string      `json:"storeId,omitempty"`
		TerminalID          string      `json:"terminalId,omitempty"`
		RequestedAt         int         `json:"requestedAt,omitempty"`
		RedirectURL         string      `json:"redirectUrl,omitempty"`
		RedirectType        string      `json:"redirectType,omitempty"`
		IsAuthorization     bool        `json:"isAuthorization,omitempty"`
		AuthorizationExpiry interface{} `json:"authorizationExpiry,omitempty"`
	}

	CreateQRResponse struct {
		ResultInfo ResultInfo   `json:"resultInfo"`
		Data       CreateQRData `json:"data,omitempty"`
	}
)

func (c *Client) CreateQR(createQRRequest CreateQRRequest) (*CreateQRResponse, error) {
	req, err := json.Marshal(createQRRequest)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}
	resp, err := c.doRequest("POST", "/v2/codes", nil, req)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}
	fmt.Println(string(resp))

	var createQRResponse CreateQRResponse
	if err := json.Unmarshal(resp, &createQRResponse); err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}
	return &createQRResponse, nil
}

type (
	DeleteQRData struct {
		PaymentID           string      `json:"paymentId,omitempty"`
		Status              string      `json:"status,omitempty"`
		AcceptedAt          int         `json:"acceptedAt,omitempty"`
		Refunds             Refunds     `json:"refunds,omitempty"`
		Captures            Captures    `json:"captures,omitempty"`
		Revert              Revert      `json:"revert,omitempty"`
		MerchantPaymentID   string      `json:"merchantPaymentId,omitempty"`
		UserAuthorizationID Amount      `json:"amount,omitempty"`
		RequestedAt         int         `json:"requestedAt,omitempty"`
		ExpiresAt           interface{} `json:"expiresAt,omitempty"`
		CanceledAt          interface{} `json:"canceledAt,omitempty"`
		TerminalID          string      `json:"terminalId,omitempty"`
		OrderReceiptNumber  string      `json:"orderReceiptNumber,omitempty"`
		OrderDescription    string      `json:"orderDescription,omitempty"`
		OrderItems          []OrderItem `json:"orderItems,omitempty"`
		Metadata            Metadata    `json:"metadata,omitempty"`
	}

	DeleteQRResponse struct {
		ResultInfo ResultInfo   `json:"resultInfo"`
		Data       DeleteQRData `json:"data,omitempty"`
	}
)

func (c *Client) DeleteQR(codeID string) (*DeleteQRResponse, error) {
	if codeID == "" {
		log.Printf("codeID is empty")
		return nil, nil
	}
	resp, err := c.doRequest("DELETE", fmt.Sprintf("/v2/codes/%s", codeID), map[string]string{}, nil)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}
	var deleteQRResponse DeleteQRResponse
	if err := json.Unmarshal(resp, &deleteQRResponse); err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}
	return &deleteQRResponse, nil
}
