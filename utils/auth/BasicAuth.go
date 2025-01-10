package auth

import (
	"net/http"
	"time"

	"github.com/Hari-Kiri/virest-utilities/utils/structures/virest"
	"github.com/golang-jwt/jwt"
	"libvirt.org/go/libvirt"
)

// Authenticate and generate JWT token for user with basic auth. After authentication succeed, new JWT will be generated with issuer name
// from argument 'applicationName' and valid until argument 'jwtLifetimeDuration'. This function use hashing PBKDF2 for basic auth.
// PBKDF2 hashing Specification:
//
// - Key length = 64 bits
//
// - Pseudo-random function = SHA512
//
// - Salt and iteration will be respected from sender
func BasicAuth(
	httpRequest *http.Request,
	username,
	password,
	applicationName string,
	jwtLifetimeDuration time.Duration,
	jwtSigningMethod *jwt.SigningMethodHMAC,
	jwtSignatureKey []byte,
) (string, virest.Error, bool) {
	usernameToken, passwordToken, ok := httpRequest.BasicAuth()
	if !ok {
		return "", virest.Error{Error: libvirt.Error{
			Code:    libvirt.ERR_AUTH_FAILED,
			Domain:  libvirt.FROM_AUTH,
			Message: "basic authentication credential not found",
			Level:   2,
		}}, true
	}

	succeed, errorBasicAuth := basicAuthVerification([]byte(username), []byte(password), usernameToken, passwordToken)
	if !succeed {
		return "", virest.Error{Error: libvirt.Error{
			Code:    libvirt.ERR_AUTH_FAILED,
			Domain:  libvirt.FROM_AUTH,
			Message: errorBasicAuth.Error(),
			Level:   2,
		}}, true
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		jwt.StandardClaims{
			Issuer:    applicationName,
			ExpiresAt: time.Now().Add(jwtLifetimeDuration).Unix(),
		},
	)

	signedToken, errorSigningToken := token.SignedString(jwtSignatureKey)
	if errorSigningToken != nil {
		libvirtError := libvirt.Error{
			Code:    libvirt.ERR_AUTH_FAILED,
			Domain:  libvirt.FROM_AUTH,
			Message: errorSigningToken.Error(),
			Level:   2,
		}
		return "", virest.Error{Error: libvirtError}, true
	}

	return signedToken, virest.Error{}, false
}
