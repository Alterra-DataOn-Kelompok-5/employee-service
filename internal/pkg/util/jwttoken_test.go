package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseJWTTokenInvalidSigningMethod(t *testing.T) {
	tk := "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJlbWFpbCI6InZpbmNlbnRsaHViYmFyZEBzdXBlcnJpdG8uY29tIiwicm9sZV9pZCI6MSwiZGl2aXNpb25faWQiOjEsImV4cCI6MTY1NzE3OTIzNn0.4Vtn8sCiM96wvNzaOEcmwMT0q4boMZ3lHYH7GRUFxWuxtq0DRAJdA6h26NjLMb_g"
	_, err := ParseJWTToken(fmt.Sprintf("Bearer %s", tk))
	if assert.Error(t, err) {
		assert.Equal(t, "invalid signing method", err.Error())
	}
}

func TestParseJWTTokenNoAuthHeader(t *testing.T) {
	tk := "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJlbWFpbCI6InZpbmNlbnRsaHViYmFyZEBzdXBlcnJpdG8uY29tIiwicm9sZV9pZCI6MSwiZGl2aXNpb25faWQiOjEsImV4cCI6MTY1NzE3OTIzNn0.4Vtn8sCiM96wvNzaOEcmwMT0q4boMZ3lHYH7GRUFxWuxtq0DRAJdA6h26NjLMb_g"
	_, err := ParseJWTToken(tk)
	if assert.Error(t, err) {
		assert.Equal(t, "authorization not found", err.Error())
	}
}

func TestParseJWTTokenExpiredToken(t *testing.T) {
	tk := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJlbWFpbCI6InZpbmNlbnRsaHViYmFyZEBzdXBlcnJpdG8uY29tIiwicm9sZV9pZCI6MSwiZGl2aXNpb25faWQiOjEsImV4cCI6MTY1NzE4OTcwN30.xii7Nrv8RN3PBzOLS994I7MOcTzrtuPlos4am9pxCB8"
	_, err := ParseJWTToken(fmt.Sprintf("Bearer %s", tk))
	if assert.Error(t, err) {
		assert.Equal(t, "Token is expired", err.Error())
	}
}

func TestParseJWTTokenSuccess(t *testing.T) {
	var (
		email      string = "vincentlhubbard@superrito.com"
		userID     uint   = 1
		roleID     uint   = 1
		divisionID uint   = 1
	)
	tk, err := CreateJWTToken(CreateJWTClaims(email, userID, roleID, divisionID))
	if err != nil {
		t.Fatal(err)
	}
	res, err := ParseJWTToken(fmt.Sprintf("Bearer %s", tk))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, email, res.Email)
	assert.Equal(t, userID, res.UserID)
	assert.Equal(t, roleID, res.RoleID)
	assert.Equal(t, divisionID, res.DivisionID)
	assert.NotNil(t, res.ExpiresAt)
}
