package security

import "golang.org/x/crypto/bcrypt"

func HashPW(pw string) (string, error) {
	bPW := []byte(pw)
	bytes, err := bcrypt.GenerateFromPassword(bPW, 4)
	hash := string(bytes)
	return hash, err
}

func CheckHash(pw, hash string) bool {
	bPW := []byte(pw)
	bHash := []byte(hash)
	err := bcrypt.CompareHashAndPassword(bHash, bPW)
	return err == nil
}
