package otp

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"strings"
	"text/template"
)

var auth smtp.Auth

type otp struct {
	cfg Config
}

func (o otp) SMS(phone string, content string) error {
	accountSid := o.cfg.TwilioSid
	authToken := o.cfg.TwilioToken
	urlStr := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", accountSid)

	msgData := url.Values{}
	msgData.Set("To", "+62"+phone[1:])
	msgData.Set("From", os.Getenv("OTP_TWILIO_SENDER"))
	msgData.Set("Body", fmt.Sprintf("Verification OTP Code %s", content))
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("sms otp error : %+v", err)
	}
	if resp.StatusCode == 201 && resp.StatusCode < 300 {
		return nil
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("sms otp error with status code : %d body : %s", resp.StatusCode, string(body))
	}
}

func (o otp) Email(email string, content string) error {
	t, err := template.ParseFiles(o.cfg.TemplatePath)
	if err != nil {
		return fmt.Errorf("no file found in this path error : %+v", err)
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, struct{ Content string }{Content: content}); err != nil {
		return fmt.Errorf("mailler send error : %+v", err)
	}
	o.cfg.Body = buf.String()

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + "Admin response to your request!" + "!\n"
	msg := []byte(subject + mime + "\n" + o.cfg.Body)
	addr := "smtp.gmail.com:587"

	if err := smtp.SendMail(addr, auth, o.cfg.Email, []string{email}, msg); err != nil {
		return err
	}
	return nil
}

func NewOTP(config Config) OTP {
	auth = smtp.PlainAuth("", config.Email, config.Password, "smtp.gmail.com")
	return &otp{
		cfg: config,
	}
}
