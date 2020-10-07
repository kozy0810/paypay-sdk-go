package paypay

import (
	"fmt"
	"net/http"
)

const (
	SandBoxBaseURL    = "https://stg-api.sandbox.paypay.ne.jp"
	ProductionBaseURL = "https://api.paypay.ne.jp"
)

type (
	ResultInfo struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		CodeID  string `json:"codeId"`
	}

	Amount struct {
		Amount   int    `json:"amount"`
		Currency string `json:"currency"`
	}

	UnitPrice struct {
		Amount   int    `json:"amount"`
		Currency string `json:"currency"`
	}

	Refund struct {
		Status           string `json:"status,omitempty"`
		AcceptedAt       int    `json:"acceptedAt,omitempty"`
		MerchantRefundID string `json:"merchantRefundId,omitempty"`
		PaymentID        string `json:"paymentId,omitempty"`
		Amount           Amount `json:"amount,omitempty"`
		RequestedAt      int    `json:"requestedAt,omitempty"`
		Reason           string `json:"reason,omitempty"`
	}

	Refunds struct {
		Data []Refund `json:"data,omitempty"`
	}

	Capture struct {
		AcceptedAt        int         `json:"acceptedAt,omitempty"`
		MerchantCaptureID string      `json:"merchantCaptureId,omitempty"`
		Amount            Amount      `json:"amount,omitempty"`
		OrderDescription  string      `json:"orderDescription,omitempty"`
		RequestedAt       int         `json:"requestedAt,omitempty"`
		ExpiresAt         interface{} `json:"expiresAt,omitempty"`
		Status            string      `json:"status,omitempty"`
	}

	Captures struct {
		Data []Capture `json:"data,omitempty"`
	}

	Revert struct {
		AcceptedAt       int    `json:"acceptedAt,omitempty"`
		MerchantRevertID string `json:"merchantRevertId,omitempty"`
		RequestedAt      int    `json:"requestedAt,omitempty"`
		Reason           string `json:"reason,omitempty"`
	}

	OrderItem struct {
		Name      string    `json:"name,omitempty"`
		Category  string    `json:"category,omitempty"`
		Quantity  int       `json:"quantity,omitempty"`
		ProductID string    `json:"productId,omitempty"`
		UnitPrice UnitPrice `json:"unitPrice,omitempty"`
	}

	Metadata struct{}

	Link struct {
		Href        string `json:"href"`
		Rel         string `json:"rel,omitempty"`
		Method      string `json:"method,omitempty"`
		Description string `json:"description,omitempty"`
		Enctype     string `json:"enctype,omitempty"`
	}

	ErrorResponseDetail struct {
		Field string `json:"field"`
		Issue string `json:"issue"`
		Links []Link `json:"link"`
	}

	ErrorResponse struct {
		Response        *http.Response        `json:"-"`
		Name            string                `json:"name"`
		DebugID         string                `json:"debug_id"`
		Message         string                `json:"message"`
		InformationLink string                `json:"information_link"`
		Details         []ErrorResponseDetail `json:"details"`
	}
)

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %s, %+v", r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message, r.Details)
}
