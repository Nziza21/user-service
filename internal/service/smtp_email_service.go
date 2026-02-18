package service

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
		"Content-Type: text/html; charset=UTF-8\r\n" + // <-- this line is key
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
	
	// HTML body
	body := fmt.Sprintf(`
		<html>
		<head>
			<style>
				.container {
					font-family: Arial, sans-serif;
					padding: 20px;
					background-color: #f4f4f4;
					border-radius: 8px;
					max-width: 600px;
					margin: auto;
					text-align: center;
				}
				.otp {
					font-size: 24px;
					font-weight: bold;
					color: #2E86C1;
				}
				.footer {
					margin-top: 20px;
					font-size: 12px;
					color: #888888;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h2>Password Reset Request</h2>
				<p>Your One-Time Password (OTP) is:</p>
				<p class="otp">%s</p>
				<p>This OTP is valid for 5 minutes.</p>
				<div class="footer">
					If you didn't request this, please ignore this email.
				</div>
			</div>
		</body>
		</html>
	`, otp)

	return s.SendEmail(to, subject, body)
}