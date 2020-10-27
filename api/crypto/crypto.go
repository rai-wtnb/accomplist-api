package crypto

import "golang.org/x/crypto/bcrypt"

// PasswordEncrypt generates hashed password
func PasswordEncrypt(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

// Verify conpares hash and enterd password
func Verify(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
