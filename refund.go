package paypay

import (
	"encoding/json"
	"fmt"
	"log"
)

type (
	GetRefundData struct {
		Status           string `json:"status,omitempty"`
		AcceptedAt       int    `json:"acceptedAt,omitempty"`
		MerchantRefundID string `json:"merchantRefundId,omitempty"`
		PaymentID        string `json:"paymentId,omitempty"`
		Amount           Amount `json:"amount,omitempty"`
		RequestedAt      int    `json:"requestedAt,omitempty"`
		Reason           string `json:"reason,omitempty"`
		AssumeMerchant   string `json:"assumeMerchant,omitempty"`
	}

	GetRefundDetailResponse struct {
		ResultInfo ResultInfo    `json:"resultInfo"`
		Data       GetRefundData `json:"data,omitempty"`
	}
)

func (c *Client) GetRefund(merchantPaymentID string) (*GetRefundDetailResponse, error) {
	resp, err := c.doRequest("GET", fmt.Sprintf("/v2/refunds/%s", merchantPaymentID), map[string]string{}, nil)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}
	var getRefundDetailResponse GetRefundDetailResponse
	if err := json.Unmarshal(resp, &getRefundDetailResponse); err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}
	return &getRefundDetailResponse, nil
}

type (
	RefundRequest struct {
		MerchantRefundID string `json:"merchantRefundId"`
		PaymentID        string `json:"paymentId"`
		Amount           Amount `json:"amount"`
		RequestedAt      int    `json:"requestedAt"`
		Reason           string `json:"reason"`
	}

	RefundData struct {
		Status           string `json:"status,omitempty"`
		AcceptedAt       int    `json:"acceptedAt,omitempty"`
		MerchantRefundID string `json:"merchantRefundId,omitempty"`
		PaymentID        string `json:"paymentId,omitempty"`
		Amount           Amount `json:"amount,omitempty"`
		RequestedAt      int    `json:"requestedAt,omitempty"`
		Reason           string `json:"reason,omitempty"`
		AssumeMerchant   string `json:"assumeMerchant,omitempty"`
	}

	RefundResponse struct {
		ResultInfo ResultInfo         `json:"resultInfo"`
		Data       RefundData `json:"data,omitempty"`
	}
)

func (c *Client) Refund(refundRequest RefundRequest) (*RefundResponse, error) {
	req, err := json.Marshal(refundRequest)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}
	resp, err := c.doRequest("POST", "/v2/refunds", map[string]string{}, req)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}

	var refundResponse RefundResponse
	if err := json.Unmarshal(resp, &refundResponse); err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}
	return &refundResponse, nil
}
