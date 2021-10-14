package otp

type Config struct {
	Password, Email        string
	TemplatePath, Body     string
	TwilioSid, TwilioToken string
}

type OTP interface {
	SMS(string, string) error
	Email(string, string) error
}

