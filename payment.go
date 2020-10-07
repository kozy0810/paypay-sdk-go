package paypay

import (
	"encoding/json"
	"errors"
	"fmt"
)

type (
	PaymentDetail struct {
		PaymentID           string      `json:"paymentId,omitempty"`
		Status              string      `json:"status,omitempty"`
		AcceptedAt          int         `json:"acceptedAt,omitempty"`
		Refunds             Refunds     `json:"refunds,omitempty"`
		Captures            Captures    `json:"captures,omitempty"`
		Revert              Revert      `json:"revert,omitempty"`
		MerchantPaymentID   string      `json:"merchantPaymentId,omitempty"`
		UserAuthorizationID string      `json:"userAuthorizationId,omitempty"`
		Amount              Amount      `json:"amount,omitempty"`
		RequestedAt         int         `json:"requestedAt,omitempty"`
		ExpiresAt           interface{} `json:"expiresAt,omitempty"`
		CanceledAt          interface{} `json:"canceledAt,omitempty"`
		StoreID             string      `json:"storeId,omitempty"`
		TerminalID          string      `json:"terminalId,omitempty"`
		OrderReceiptNumber  string      `json:"orderReceiptNumber,omitempty"`
		OrderDescription    string      `json:"orderDescription,omitempty"`
		OrderItems          []OrderItem `json:"orderItems,omitempty"`
		Metadata            Metadata    `json:"metadata,omitempty"`
	}

	PaymentDetailResponse struct {
		ResultInfo ResultInfo    `json:"resultInfo"`
		Data       PaymentDetail `json:"data,omitempty"`
	}

	ResponseDeletePayment struct {
		ResultInfo ResultInfo `json:"resultInfo"`
	}
)

func (c *Client) GetPaymentDetail(merchantPaymentID string) (*PaymentDetailResponse, error) {
	if merchantPaymentID == "" {
		return nil, errors.New("merchantPaymentID must be empty")
	}
	resp, err := c.doRequest("GET", fmt.Sprintf("/v2/codes/payments/%s", merchantPaymentID), map[string]string{}, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(resp))

	var paymentDetail PaymentDetailResponse
	if err := json.Unmarshal(resp, &paymentDetail); err != nil {
		return nil, err
	}
	return &paymentDetail, nil
}

func (c *Client) CancelPayment(merchantPaymentID string) (*ResponseDeletePayment, error) {
	resp, err := c.doRequest("DELETE", fmt.Sprintf("/v2/payments/%v", merchantPaymentID), map[string]string{}, nil)
	if err != nil {
		return nil, err
	}
	var responseDeletePayment ResponseDeletePayment
	if err := json.Unmarshal(resp, &responseDeletePayment); err != nil {
		return nil, err
	}
	return &responseDeletePayment, nil
}
