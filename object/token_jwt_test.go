package object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClaimsCustomOmitsEmptyNonce(t *testing.T) {
	claims := Claims{
		Nonce: "",
	}

	result := getClaimsCustom(claims, nil, nil)

	_, exists := result["nonce"]
	assert.False(t, exists)
}

func TestGetClaimsCustomIncludesNonceWhenPresent(t *testing.T) {
	claims := Claims{
		Nonce: "test-nonce",
	}

	result := getClaimsCustom(claims, nil, nil)

	assert.Equal(t, "test-nonce", result["nonce"])
}
