package mail

import (
	"bytes"
	forgotPasswordDto "golang-bootcamp-1/internal/forgot-password/dto"
	registerDto "golang-bootcamp-1/internal/register/dto"
	"html/template"
	"os"
	"path/filepath"
)

type IMail interface {
	SendVerification(to string, data registerDto.EmailVerification)
	SendForgotPassword(to string, data forgotPasswordDto.ForgotPasswordEmailBody)
	SendMail(to string, subject string, content string)
}

// Parse HTML and generate data for HTML template
func ParseEmailHTML(path string, data any) (string, error) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	err = tmpl.Execute(buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func GetMailTemplateFile(filename string) string {
	// Parse file
	root, _ := os.Getwd()
	path := filepath.Join(root, "templates", "email", filename)

	return path
}
