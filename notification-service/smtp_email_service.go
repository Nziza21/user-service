package notificationservice

import (
	"fmt"
	"net/smtp"
)

type SMTPEmailService struct {
	username string
	password string
	host     string
	port     string
	from     string
}

func NewSMTPEmailService() *SMTPEmailService {
	username := "<smtp@mailtrap.io>" 
	password := "<41c92c339184eb9fc7a6b6dbf6add059>" 
	host := "live.smtp.mailtrap.io"
	port := "587"
	from := "hello@nzizasamuel.com"

	return &SMTPEmailService{
		username: username,
		password: password,
		host:     host,
		port:     port,
		from:     from,
	}
}

func (s *SMTPEmailService) SendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", s.username, s.password, s.host)

	msg := []byte(fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/plain; charset=\"utf-8\"\r\n"+
			"\r\n%s",
		s.from, to, subject, body,
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
	body := fmt.Sprintf(
		"Hello,\n\nYour OTP is %s. It expires in 5 minutes.\n\nBest,\nUser Service Team",
		otp,
	)
	return s.SendEmail(to, subject, body)
}
