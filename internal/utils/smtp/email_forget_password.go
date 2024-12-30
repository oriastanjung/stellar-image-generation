package smtp

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/oriastanjung/stellar/internal/config"
)

// sendEmail is a helper function to send emails with a given subject and body
func sendEmail(to, subject, body string) error {
	cfg := config.LoadEnv()
	from := cfg.GmailEmail             // Sender email
	password := cfg.GmailPassword      // Sender email password
	smtpServer := "smtp.gmail.com:587" // Gmail SMTP server with TLS port

	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		"\r\n" +
		body)

	// Create a channel to receive errors
	errCh := make(chan error)

	// Start a goroutine to send the email
	go func() {
		auth := smtp.PlainAuth("", from, password, strings.Split(smtpServer, ":")[0])
		err := smtp.SendMail(smtpServer, auth, from, []string{to}, msg)
		if err != nil {
			errCh <- err
			return
		}
		errCh <- nil
	}()

	// Wait for the result from the goroutine
	err := <-errCh
	return err
}

// SendEmailForgetPassword sends a password reset email with a dynamic button
func SendEmailForgetPassword(to, forgetPasswordToken, link string) error {
	subject := "Reset Password Account"
	body := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Reset Password</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				background-color: #f9f9f9;
				color: #333;
				margin: 0;
				padding: 0;
			}
			.header {
				text-align: center;
				margin-bottom: 20px;
			}
			.header img {
				width: 150px;
				height: auto;
			}
			.container {
				width: 100%%;
				max-width: 600px;
				margin: 30px auto;
				background-color: #fff;
				padding: 20px;
				border: 1px solid #e5e5e5;
				border-radius: 8px;
			}
			.title {
				font-size: 24px;
				font-weight: bold;
				color: #000;
				margin-bottom: 10px;
				text-align: left;
			}
			.text {
				font-size: 16px;
				line-height: 1.6;
				color: #555;
				margin-bottom: 20px;
				text-align: left;
			}
			.button-container {
				text-align: center;
			}
			.button {
				background-color: black;
				color: white;
				padding: 15px 32px;
				text-decoration: none;
				border-radius: 4px;
				font-size: 16px;
				margin-top: 10px;
				display: inline-block;
			}
			.footer {
				font-size: 12px;
				color: #888;
				text-align: center;
				margin-top: 20px;
			}
		</style>
	</head>
	<body>
		<div class="container">
		<div class="header">
				<img src="https://res.cloudinary.com/drsfd7hqt/image/upload/v1732671727/d1xwixs02xkmjwqmocx4.png" alt="Stellar">
			</div>
			<div class="title">Reset Password Request</div>
			<div class="text">
				<p>Kami menerima request anda untuk mereset password anda, silahkan klik tombol di bawah ini:</p>
			</div>
			<div class="button-container">
				<a href="%s/%s" class="button" 
				style="
						background-color: #000; 
						color: #fff; 
						padding: 15px 32px; 
						font-size: 16px; 
						border: none; 
						border-radius: 4px; 
						text-decoration: none; 
						display: inline-block; 
						margin-top: 10px; 
						cursor: pointer; 
						font-family: Arial, sans-serif;
						text-align: center;"
				>Reset Password</a>
			</div>
			<div class="text">
				<p>Jika anda merasa tidak pernah mendaftar akun Stellar, silahkan abaikan email ini.</p>
				<p>Salam Hangat,<br>Stellar</p>
			</div>
			<div class="footer">
				<p>© Stellar.</p>
			</div>
		</div>
	</body>
	</html>
	`, link, forgetPasswordToken)

	return sendEmail(to, subject, body)
}
