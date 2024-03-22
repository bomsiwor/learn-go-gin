package hasher

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(raw string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(raw), 10)

	if err != nil {
		return "", err
	}

	return string(password), nil
}

// An API for bcrypt compare and hash
func ValidatePassword(password, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}
