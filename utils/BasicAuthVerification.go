package utils

import (
	"bytes"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

// Verify username and password token from basic auth request using PBKDF2 hashing algorithm.
func BasicAuthVerification(usernameToken, passwordToken string) (bool, error) {
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
		[]byte(os.Getenv("VIREST_STORAGE_POOL_APPLICATION_BA_USER")),
		[]byte(usernameTokenSliced[2]),
		userTokensIteration,
		64,
		sha512.New,
	)

	passwordKey := pbkdf2.Key(
		[]byte(os.Getenv("VIREST_STORAGE_POOL_APPLICATION_BA_PASSWORD")),
		[]byte(passwordTokenSliced[2]),
		passwordTokensIteration,
		64,
		sha512.New,
	)

	hashedUsername, errorGetHashedUsername := base64.StdEncoding.DecodeString(usernameTokenSliced[3])
	if errorGetHashedUsername != nil {
		return false, fmt.Errorf("failed get hashed username: %s", errorGetHashedUsername)
	}

	hashedPassword, errorGetHashedPassword := base64.StdEncoding.DecodeString(passwordTokenSliced[3])
	if errorGetHashedPassword != nil {
		return false, fmt.Errorf("failed get hashed password: %s", errorGetHashedPassword)
	}

	if bytes.Compare(userKey, hashedUsername) != 0 {
		return false, fmt.Errorf("incorrect username")
	}
	if bytes.Compare(passwordKey, hashedPassword) != 0 {
		return false, fmt.Errorf("incorrect password")
	}

	return true, nil
}
