package paypay

import "testing"

var CreateQRJSON = []byte(`
{
	"resultInfo": {
		"code": "SUCCESS",
		"message": "Success",
		"codeId": "08100001"
	},
	"data": {
		"codeId": "04-373vfjL04aNyu9jh",
		"url": "https://qr-stg.sandbox.paypay.ne.jp/28180104373vfjL04aNyu9jh",
		"deeplink": "paypay://payment?link_key=https%3A%2F%2Fqr-stg.sandbox.paypay.ne.jp%2F28180104373vfjL04aNyu9jh",
		"expiryDate": 1602226535,
		"merchantPaymentId": "3b5b82d675834d418c1fbf27445ac5xij",
		"amount": {
			"amount": 10,
			"currency": "JPY"
		},
		"orderItems": [
			{
				"name": "product1",
				"category": "cake",
				"quantity": 1,
				"productId": "1",
				"unitPrice": {
					"amount": 0,
					"currency": ""
				}
			}
		],
		"metadata": {},
		"codeType": "ORDER_QR",
		"redirectUrl": "http://localhost:8000/",
		"redirectType": "WEB_LINK"
	}
}
`)

func TestCreateQR(t *testing.T) {
	mock, transport := NewMockClient(201, CreateQRJSON)
	c, err := NewClient("api-key", "api-secret", 0000000000000000, false, mock)
	if err != nil {
		t.Error(err)
	}

	qr, err := c.CreateQR(CreateQRRequest{
		MerchantPaymentID: "3b5b82d675834d418c1fbf27445ac5xij",
		CodeType: "ORDER_QR",
		OrderItems: []OrderItem{
			{
				Name: "test1",
				Category: "test",
				Quantity: 1,
				ProductID: "test1",
				UnitPrice: UnitPrice{
					Amount: 1,
					Currency: "JPY",
				},
			},
		},
		RedirectURL: "http://localhost:8000/",
		RedirectType: "WEB_LINK",
	})
	if transport.URL != "https://stg-api.sandbox.paypay.ne.jp/v2/codes" {
		t.Errorf("URL is wrong: %s", transport.URL)
	}
	if transport.Method != "POST" {
		t.Errorf("Method should be Get, but %s", transport.Method)
	}
	if err != nil {
		t.Errorf("err should be nil, but %v", err)
		return
	} else if qr == nil {
		t.Error("qr should not be nil")
	} else if qr.ResultInfo.Code != "SUCCESS" {
		t.Errorf("ResultInfo Code should be SUCCESS, but %s", qr.ResultInfo.Code)
	} else if qr.Data.MerchantPaymentID != "3b5b82d675834d418c1fbf27445ac5xij" {
		t.Errorf("MerchantPaymentId should be 3b5b82d675834d418c1fbf27445ac5xij, but %s", qr.Data.MerchantPaymentID)
	}
}