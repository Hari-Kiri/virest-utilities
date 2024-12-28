package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Hari-Kiri/virest-utilities/utils/structures/virest"
	"github.com/golang-jwt/jwt"
	"libvirt.org/go/libvirt"
)

// Parse and validate request bearer token.
func BearerTokenAuth(
	httpRequest *http.Request,
	applicationName string,
	jwtSigningMethod *jwt.SigningMethodHMAC,
	jwtSignatureKey []byte,
) (virest.Error, bool) {
	bearerToken := httpRequest.Header.Get("Authorization")
	if !strings.Contains(bearerToken, "Bearer") {
		libvirtError := libvirt.Error{
			Code:    libvirt.ERR_AUTH_FAILED,
			Domain:  libvirt.FROM_NET,
			Message: "bearer token not exist",
			Level:   libvirt.ERR_ERROR,
		}
		return virest.Error{Error: libvirtError}, true
	}
	tokenString := strings.Replace(bearerToken, "Bearer ", "", -1)

	token, errorTokenValidation := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Signing method invalid")
		}
		if method != jwtSigningMethod {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return jwtSignatureKey, nil
	})
	if errorTokenValidation != nil {
		libvirtError := libvirt.Error{
			Code:    libvirt.ERR_AUTH_FAILED,
			Domain:  libvirt.FROM_NET,
			Message: fmt.Sprintf("failed validate JWT token: %s", errorTokenValidation),
			Level:   libvirt.ERR_ERROR,
		}
		return virest.Error{Error: libvirtError}, true
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		libvirtError := libvirt.Error{
			Code:    libvirt.ERR_AUTH_FAILED,
			Domain:  libvirt.FROM_NET,
			Message: "failed get JWT claims",
			Level:   libvirt.ERR_ERROR,
		}
		return virest.Error{Error: libvirtError}, true
	}

	if !claims.VerifyIssuer(applicationName, true) {
		libvirtError := libvirt.Error{
			Code:    libvirt.ERR_AUTH_FAILED,
			Domain:  libvirt.FROM_NET,
			Message: fmt.Sprintf("failed validate JWT token: Issuer not from %s", applicationName),
			Level:   libvirt.ERR_ERROR,
		}
		return virest.Error{Error: libvirtError}, true
	}

	return virest.Error{}, false
}
