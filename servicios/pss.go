package servicios

import "golang.org/x/crypto/bcrypt"

// PasswordHash hashea una contraseña utilizando bcrypt
func PasswordHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
