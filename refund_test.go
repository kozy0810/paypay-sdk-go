package paypay

import "testing"

var GetRefundJSON = []byte(`
{
	"resultInfo": {
		"code": "SUCCESS",
		"message": "Success",
		"codeId": "08100001"
	},
	"data": {
		"status": "REFUNDED",
		"acceptedAt": 1601565829,
		"merchantRefundId": "3b5b82d675834d418c1fbf27445ac6cb",
		"paymentId": "03450018268289114112",
		"amount": {
			"amount": 1,
			"currency": "JPY"
		},
		"requestedAt": 1601565828,
		"assumeMerchant": "ï¿½"
	}
}
`)

var RefundJSON = []byte(`
{
	"resultInfo": {
		"code": "SUCCESS",
		"message": "Success",
		"codeId": "08100001"
	},
	"data": {
		"status": "CREATED",
		"acceptedAt": 1602344804,
		"merchantRefundId": "3a5b82d675834d418c1fbf27544bcxyz",
		"paymentId": "03461188464614121472",
		"amount": {
			"amount": 1,
			"currency": "JPY"
		},
		"requestedAt": 1602344803,
		"reason": "test",
		"assumeMerchant": "ï¿½"
	}
}
`)

func TestGetRefund(t *testing.T) {
	mock, transport := NewMockClient(200, GetRefundJSON)
	c, err := NewClient("api-key", "api-secret", 0000000000000000, false, mock)
	if err != nil {
		t.Errorf("err should be nil, but %s", err)
	}

	revert, err := c.GetRefund("3b5b82d675834d418c1fbf27445ac6cb")
	if transport.URL != "https://stg-api.sandbox.paypay.ne.jp/v2/refunds/3b5b82d675834d418c1fbf27445ac6cb" {
		t.Errorf("URL is wrong: %s", transport.URL)
	}
	if transport.Method != "GET" {
		t.Errorf("Method should be GET, but %s", transport.Method)
	}
	if err != nil {
		t.Errorf("err should be nil, but %s", err)
	} else if revert == nil {
		t.Errorf("revert should be nil, but %s", err)
	} else if revert.ResultInfo.Code != "SUCCESS" {
		t.Errorf("ResultInfo Code should be SUCCESS, but %s", revert.ResultInfo.Code)
	} else if revert.Data.Status != "REFUNDED" {
		t.Errorf("Status should be SUCCESS, but %s", revert.Data.Status)
	}
}

func TestRefund(t *testing.T) {
	mock, transport := NewMockClient(200, RefundJSON)
	c, err := NewClient("api-key", "api-secret", 0000000000000000, false, mock)
	if err != nil {
		t.Errorf("err should be nil, but %s", err)
	}

	req := RefundRequest{
		MerchantRefundID: "3a5b82d675834d418c1fbf27544bcxyz",
		PaymentID: "03461188464614121472",
		Amount: Amount{
			Amount: 1,
			Currency: "JPY",
		},
		RequestedAt: 1602344803,
		Reason: "test",
	}

	refund, err := c.Refund(req)
	if transport.URL != "https://stg-api.sandbox.paypay.ne.jp/v2/refunds" {
		t.Errorf("URL is wrong: %s", transport.URL)
	}
	if transport.Method != "POST" {
		t.Errorf("Method should be POST, but %s", transport.Method)
	}
	if err != nil {
		t.Errorf("err should be nil, but %v", err)
	} else if refund == nil {
		t.Error("refund should not be nil")
	} else if refund.ResultInfo.Code != "SUCCESS" {
		t.Errorf("ResultInfo Code should be SUCCESS, but %s", refund.ResultInfo.Code)
	} else if refund.Data.MerchantRefundID != "3a5b82d675834d418c1fbf27544bcxyz" {
		t.Errorf("MerchantRefundID should be 3a5b82d675834d418c1fbf27544bcxyz, but %s", refund.Data.MerchantRefundID)
	} else if refund.Data.PaymentID != "03461188464614121472" {
		t.Errorf("PaymentID should be 03461188464614121472, but %s", refund.Data.PaymentID)
	}
}