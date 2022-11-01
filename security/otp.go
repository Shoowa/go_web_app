package security

import (
	"bytes"
	"image"
	"image/png"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	timedPW "github.com/pquerna/otp/totp"
)

func GenerateKey(email string) (*otp.Key, error) {
	subject := timedPW.GenerateOpts{
		Issuer:      "asymblur.com",
		AccountName: email,
	}

	key, err := timedPW.Generate(subject)
	return key, err
}

func CreateQR(key *otp.Key) (image.Image, error) {
	img, err := key.Image(200, 200)
	return img, err
}

func BufferQR(img image.Image) bytes.Buffer {
	var buf bytes.Buffer
	png.Encode(&buf, img)
	return buf
}

func CheckCode(code string, secret string) bool {
	valid := totp.Validate(code, secret)
	return valid
}
