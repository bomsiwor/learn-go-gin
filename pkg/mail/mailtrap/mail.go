package mail

import (
	"fmt"
	forgot_password "golang-bootcamp-1/internal/forgot-password/dto"
	register "golang-bootcamp-1/internal/register/dto"
	globalMail "golang-bootcamp-1/pkg/mail"
	"log"
	"net/smtp"
	"os"
	"time"
)

type mailUsecase struct {
}

func NewMailUsecase() globalMail.IMail {
	return &mailUsecase{}
}

// SendForgotPassword implements mail.IMail.
func (uc *mailUsecase) SendForgotPassword(to string, data forgot_password.ForgotPasswordEmailBody) {
	// Get email template path
	path := globalMail.GetMailTemplateFile("forgot_password_email.html")

	// Parse
	result, err := globalMail.ParseEmailHTML(path, data)
	if err != nil {
		log.Println(err)
	}

	// Send mail
	uc.SendMail(to, data.Subject, result)
}

// SendVerification implements mail.IMail.
func (uc *mailUsecase) SendVerification(to string, data register.EmailVerification) {
	// Get email template filepath
	path := globalMail.GetMailTemplateFile("verification_email.html")

	// Parse file
	result, err := globalMail.ParseEmailHTML(path, data)
	if err != nil {
		log.Println(err)
	}

	uc.SendMail(to, data.Subject, result)
}

// sendMail implements mail.IMail.
func (*mailUsecase) SendMail(to string, subject string, content string) {
	host := os.Getenv("MAILTRAP_HOST") + ":" + os.Getenv("MAILTRAP_PORT")
	// host := "127.0.0.1:1025"

	// Build the message
	mime := "MIME-version: 1.0\nContent-Type: text/html; charset=\"UTF-8\"\n\n"
	header := fmt.Sprintf("From: %s\r\n", os.Getenv("MAIL_SENDER_NAME"))
	header += fmt.Sprintf("To: %s\r\n", to)
	header += fmt.Sprintf("Subject: %s\r\n", subject)
	header += fmt.Sprintf("Date: %s\r\n", time.Now().Format(time.RFC1123))
	header += mime

	auth := smtp.PlainAuth(
		"",
		os.Getenv("MAILTRAP_USER"),
		os.Getenv("MAILTRAP_PASSWORD"),
		os.Getenv("MAILTRAP_HOST"),
	)

	receiver := []string{to}
	msg := []byte(header + "\n" + content)

	err := smtp.SendMail(
		host,
		auth,
		os.Getenv("MAIL_SENDER_NAME"),
		receiver,
		msg,
	)
	log.Println(err)
	if err != nil {
		log.Println(err)
	}
}
