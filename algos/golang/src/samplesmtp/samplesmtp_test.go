package samplesmtp

import "testing"

func TestSendMail(t *testing.T) {

	// server := "127.0.0.1:2526"
	server := "192.168.214.44:25"
	// server := "10.182.128.220:25"
	from := "abhijith.da@veritas.com"
	to := "abhijith.da@veritas.com"
	body := "This is the email body"

	expResp := true

	resp := SendMail(server, from, to, body)

	if resp != expResp {
		t.Errorf("got %v, want %v", resp, expResp)
	}
}
