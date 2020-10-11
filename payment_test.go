package paypay

import "testing"

var GetPaymentDetailJSON = []byte(`
{
  "resultInfo": {
	"code": "SUCCESS",
	"message": "Success",
	"codeId": "08100001"
  },
  "data": {
	"paymentId": "12341234123412341234",
	"status": "COMPLETED",
	"acceptedAt": 1601391628,
	"refunds": {
		"data": []
	},
	"merchantPaymentId": "3b5b82d675834d418c1fbf27445ac5xij",
	"amount": {
		"amount": 1,
		"currency": "JPY"
	},
	"requestedAt": 1601391628,
	"storeId": "",
	"terminalId": "",
	"orderDescription": "",
	"orderItems": [
		{
			"name": "test",
			"category": "category_test",
			"quantity": 1,
			"productId": "1",
			"unitPrice": {
				"amount": 1,
				"currency": "JPY"
			}
		}
	],
	"metadata": {}
	}
}
`)

var CancelPaymentJOSN = []byte(`
{
	"resultInfo": {
		"code": "REQUEST_ACCEPTED",
		"message": "Request accepted",
		"codeId": "08100001"
	}
}
`)

func TestGetPaymentDetail(t *testing.T) {
	mock, transport := NewMockClient(200, GetPaymentDetailJSON)
	c, err := NewClient("api-key", "api-secret", 0000000000000000, false, mock)
	if err != nil {
		t.Errorf("err should be nil, but %s", err)
	}

	payment, err := c.GetPaymentDetail("3b5b82d675834d418c1fbf27445ac5xij")
	if transport.URL != "https://stg-api.sandbox.paypay.ne.jp/v2/codes/payments/3b5b82d675834d418c1fbf27445ac5xij" {
		t.Errorf("URL is wrong: %s", transport.URL)
	}
	if transport.Method != "GET" {
		t.Errorf("Method should be Get, but %s", transport.Method)
	}
	if err != nil {
		t.Errorf("err should be nil, but %v", err)
	} else if payment == nil {
		t.Error("payment should not be nil")
	} else if payment.Data.PaymentID != "12341234123412341234" {
		t.Errorf("PaymentID should be 12341234123412341234, but %s", payment.Data.PaymentID)
	} else if payment.ResultInfo.Code != "SUCCESS" {
		t.Errorf("ResultInfo Code should be SUCCESS, but %s", payment.ResultInfo.Code)
	}
}

func TestCancelPayment(t *testing.T) {
	mock, transport := NewMockClient(202, CancelPaymentJOSN)
	c, err := NewClient("api-key", "api-secret", 0000000000000000, false, mock)
	if err != nil {
		t.Error(err)
	}

	cancel, err := c.CancelPayment("3b5b82d675834d418c1fbf27445ac5xij")
	if transport.URL != "https://stg-api.sandbox.paypay.ne.jp/v2/payments/3b5b82d675834d418c1fbf27445ac5xij" {
		t.Errorf("URL is wrong: %s", transport.URL)
	}
	if transport.Method != "DELETE" {
		t.Errorf("Method should be DELETE, but %s", transport.Method)
	}
	if err != nil {
		t.Errorf("err should be nil, but %v", err)
	} else if cancel == nil {
		t.Error("cancel should not be nil")
	} else if cancel.ResultInfo.Code != "REQUEST_ACCEPTED" {
		t.Errorf("ResultInfo Code should be REQUEST_ACCEPTED, but %s", cancel.ResultInfo.Code)
	}
}
