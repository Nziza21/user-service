package notificationservice

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

type SMTPEmailService struct{}

func NewSMTPEmailService() *SMTPEmailService {
	return &SMTPEmailService{}
}

func (s *SMTPEmailService) SendEmail(to, subject, body string) error {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	from := os.Getenv("SMTP_FROM")

	if host == "" || port == "" || username == "" || password == "" || from == "" {
		return fmt.Errorf("SMTP credentials are not set in .env")
	}

	msg := []byte(
		"From: " + from + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" + body + "\r\n",
	)

	auth := smtp.PlainAuth("", username, password, host)

	if err := smtp.SendMail(host+":"+port, auth, from, []string{to}, msg); err != nil {
		log.Println("send email failed:", err)
		return err
	}

	log.Println("email sent successfully to", to)
	return nil
}

func (s *SMTPEmailService) SendOTPEmail(to, otp string) error {
	subject := "Reset Password OTP"
	body := fmt.Sprintf("Your OTP is: %s", otp)
	return s.SendEmail(to, subject, body)
}
