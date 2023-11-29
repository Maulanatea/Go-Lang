package utils

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {
	hashedBybe, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(hashedBybe), err
}

func CheckPasswordHash(password, HashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(HashedPassword), []byte(password))
	return err == nil

}
