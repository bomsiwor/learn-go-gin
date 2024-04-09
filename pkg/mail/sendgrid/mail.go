package mail

import (
	forgot_password "golang-bootcamp-1/internal/forgot-password/dto"
	registerDto "golang-bootcamp-1/internal/register/dto"
	globalMail "golang-bootcamp-1/pkg/mail"
	"log"
	"os"
	"path/filepath"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type mailUsecase struct {
}

func NewMailUsecase() globalMail.IMail {
	return &mailUsecase{}
}

// SendForgotPassword implements mail.IMail.
func (*mailUsecase) SendForgotPassword(to string, data forgot_password.ForgotPasswordEmailBody) {
	panic("unimplemented")
}

// Send email verification
// Read HTML file
// Send email based on HTML File
func (uc *mailUsecase) SendVerification(to string, data registerDto.EmailVerification) {
	// Parse file
	root, _ := os.Getwd()
	path := filepath.Join(root, "templates", "email", "verification_email.html")

	result, err := globalMail.ParseEmailHTML(path, data)
	if err != nil {
		log.Println(err)
	}

	uc.SendMail(to, data.Subject, result)
}

func (uc *mailUsecase) SendMail(to string, subject string, content string) {
	// Create from
	from := mail.NewEmail(os.Getenv("MAIL_SENDER_NAME"), os.Getenv("MAIL_SENDER_NAME"))
	// Create receiver
	receiver := mail.NewEmail(to, to)

	// Create mail data for client sending
	message := mail.NewSingleEmail(from, subject, receiver, "", content)

	// Init client
	client := sendgrid.NewSendClient(os.Getenv("MAIL_KEY"))

	// Send and get response
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else if response.StatusCode != 200 {
		log.Println(response.StatusCode, response)
	} else {
		log.Println(response.Body)
	}
}
