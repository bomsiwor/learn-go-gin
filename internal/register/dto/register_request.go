package register

type EmailVerification struct {
	Subject          string
	Email            string
	VerificationCode string
}
