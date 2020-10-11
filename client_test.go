package paypay

import "testing"

func TestNewClient(t *testing.T) {
	c, err := NewClient("api-key", "api-secret", 0000000000000000, false, nil)
	if err != nil {
		t.Errorf("%v", err)
	}
	if c == nil {
		t.Error("service should be valid")
	}
	if c.APIBase() != SandBoxBaseURL {
		t.Errorf(`APIBase should be "%s", but "%s"`, SandBoxBaseURL, c.APIBase())
	}
}