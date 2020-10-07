package paypay

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// reference url : https://dev.to/plutov/writing-rest-api-client-in-go-3fkg

type Client struct {
	apiKey     string
	apiSecret  string
	apiBase    string
	merchantID int
	Client     *http.Client
}

func NewClient(key, secret string, merchantID int, isProduction bool, client *http.Client) (*Client, error) {
	if key == "" || secret == "" {
		return nil, errors.New("APIKey, APISecret and APIBase are required to create a Client")
	}
	if client == nil {
		client = &http.Client{}
	}

	c := &Client{
		apiKey:     key,
		apiSecret:  secret,
		merchantID: merchantID,
		Client:     client,
	}
	if isProduction {
		c.apiBase = ProductionBaseURL
	} else {
		c.apiBase = SandBoxBaseURL
	}
	return c, nil
}

func (c *Client) APIBase() string {
	return c.apiBase
}

func (c *Client) header(method, resourceUrl string, body []byte) map[string]string {
	contentType := "application/json"
	payload := string(body)
	switch method {
	case http.MethodGet, http.MethodDelete:
		contentType = "empty"
		payload = "empty"
	case http.MethodPost, http.MethodPut:
		if body != nil {
			md5 := md5.New()
			md5.Write([]byte(contentType))
			md5.Write([]byte(payload))
			payload = base64.StdEncoding.EncodeToString(md5.Sum(nil))
		}
	default:
		return nil
	}

	epoch := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := uuid.New().String()[:8]
	signatureRawList := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s", resourceUrl, method, nonce, epoch, contentType, payload)
	h := hmac.New(sha256.New, []byte(c.apiSecret))
	h.Write([]byte(signatureRawList))
	hashed64 := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return map[string]string{
		"Authorization":     fmt.Sprintf("hmac OPA-Auth:%s:%s:%s:%s:%s", c.apiKey, hashed64, nonce, epoch, payload),
		"X-ASSUME-MERCHANT": string(c.merchantID),
		"Content-Type":      contentType,
	}
}

func (c *Client) doRequest(method, path string, query map[string]string, data []byte) (body []byte, err error) {
	baseURL, err := url.Parse(c.apiBase)
	if err != nil {
		// TODO エラー時の処理はあとですべて書き換える
		return nil, err
	}
	apiURL, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	endpoint := baseURL.ResolveReference(apiURL).String()

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	for key, value := range c.header(method, req.URL.RequestURI(), data) {
		req.Header.Add(key, value)
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
