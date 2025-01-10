package auth

import (
	"bytes"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

// Verify username and password token from basic auth request using PBKDF2 hashing algorithm.
// PBKDF2 hashing Specification:
// - Key length = 64 bits
// - Pseudo-random function = SHA512
// - Salt and iteration will be respected from sender
func basicAuthVerification(username, password []byte, usernameToken, passwordToken string) (bool, error) {
	usernameTokenSliced := strings.Split(usernameToken, "$")
	passwordTokenSliced := strings.Split(passwordToken, "$")

	userTokensIteration, errorGetUserToken := strconv.Atoi(usernameTokenSliced[1])
	if errorGetUserToken != nil {
		return false, fmt.Errorf("failed get iteration number on username token: %s", errorGetUserToken)
	}

	passwordTokensIteration, errorGetPasswordToken := strconv.Atoi(passwordTokenSliced[1])
	if errorGetPasswordToken != nil {
		return false, fmt.Errorf("failed get iteration number on password token: %s", errorGetPasswordToken)
	}

	userKey := pbkdf2.Key(
		username,
		[]byte(usernameTokenSliced[2]),
		userTokensIteration,
		64,
		sha512.New,
	)

	passwordKey := pbkdf2.Key(
		password,
		[]byte(passwordTokenSliced[2]),
		passwordTokensIteration,
		64,
		sha512.New,
	)

	hashedUsername, errorGetHashedUsername := base64.StdEncoding.DecodeString(usernameTokenSliced[3])
	if errorGetHashedUsername != nil {
		return false, fmt.Errorf("failed decode username: %s", errorGetHashedUsername)
	}

	hashedPassword, errorGetHashedPassword := base64.StdEncoding.DecodeString(passwordTokenSliced[3])
	if errorGetHashedPassword != nil {
		return false, fmt.Errorf("failed decode password: %s", errorGetHashedPassword)
	}

	if bytes.Compare(userKey, hashedUsername) != 0 {
		return false, fmt.Errorf("incorrect username")
	}
	if bytes.Compare(passwordKey, hashedPassword) != 0 {
		return false, fmt.Errorf("incorrect password")
	}

	return true, nil
}
