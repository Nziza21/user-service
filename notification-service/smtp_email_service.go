package notificationservice

import (
	"fmt"
	"net/smtp"
	"os"
)

type SMTPEmailService struct {
	username string
	password string
	host     string
	port     string
	from     string
}

func NewSMTPEmailService() *SMTPEmailService {
	username := os.Getenv("GMAIL_USERNAME")
	password := os.Getenv("GMAIL_APP_PASSWORD")

	if username == "" || password == "" {
		panic("GMAIL_USERNAME or GMAIL_APP_PASSWORD not set in environment")
	}

	return &SMTPEmailService{
		username: username,
		password: password,
		host:     "smtp.gmail.com",
		port:     "587",
		from:     username,
	}
}

func (s *SMTPEmailService) SendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", s.username, s.password, s.host)

	msg := []byte(fmt.Sprintf(
		"To: %s\r\n"+
			"Subject: %s\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/plain; charset=\"utf-8\"\r\n"+
			"\r\n%s",
		to, subject, body,
	))

	addr := s.host + ":" + s.port
	if err := smtp.SendMail(addr, auth, s.from, []string{to}, msg); err != nil {
		fmt.Println("SMTP error:", err)
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Println("Email sent to:", to)
	return nil
}

func (s *SMTPEmailService) SendOTPEmail(to, otp string) error {
	subject := "Password Reset OTP"
	body := fmt.Sprintf("Hello,\n\nYour OTP is %s. It expires in 5 minutes.\n\nBest,\nUser Service Team", otp)
	return s.SendEmail(to, subject, body)
}
