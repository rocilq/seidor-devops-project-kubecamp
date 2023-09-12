package controllers

import (
	"authService/lib"
	"fmt"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

const (
	testMail     = "mail@testmail.com"
	testPsw      = "ok_thispsw"
	testUsername = "pepito"
	testUserId	 = "1234567890"
)

func TestSignUp(t *testing.T) {

	app := fiber.New()
	app.Post("/users", SignUp)

	reqBody := fmt.Sprintf(`{
    "email": "%s",
    "password": "%s",
    "username": "%s"
	}`, testMail, testPsw, testUsername)

	req := httptest.NewRequest("POST", "/users", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

	invalidReqBody := fmt.Sprintf(`{
		"email": "%s",
		"password": "%s"
	}`, testMail, testPsw)

	invalidReq := httptest.NewRequest("POST", "/users", strings.NewReader(invalidReqBody))
	invalidReq.Header.Set("Content-Type", "application/json")

	resp, err = app.Test(invalidReq)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

	repeatedEmailReq := httptest.NewRequest("POST", "/users", strings.NewReader(reqBody))
	repeatedEmailReq.Header.Set("Content-Type", "application/json")

	resp, err = app.Test(repeatedEmailReq)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}

func TestLogin(t *testing.T) {
	app := fiber.New()
	app.Post("/login", Login)

	reqBody := fmt.Sprintf(`{
    "email": "%s",
    "password": "%s"
	}`, testMail, testPsw)

	req := httptest.NewRequest("POST", "/login", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, resp.Header.Get("Authorization"))

	badPsw := "randomeswsafm"
	badCredentialsReqBody := fmt.Sprintf(`{
		"email": "%s",
		"password": "%s"
	}`, testMail, badPsw)

	badCredentialsReq := httptest.NewRequest("POST", "/login", strings.NewReader(badCredentialsReqBody))
	badCredentialsReq.Header.Set("Content-Type", "application/json")

	resp, err = app.Test(badCredentialsReq)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)

	badReqBody := fmt.Sprintf(`{
		"email": "%s",
	}`, "random@imail.com")

	badReq := httptest.NewRequest("POST", "/login", strings.NewReader(badReqBody))
	badReq.Header.Set("Content-Type", "application/json")

	resp, err = app.Test(badReq)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestValidate(t *testing.T) {
	os.Setenv("SECRET_KEY", "SECRET_KEY")
	app := fiber.New()
	app.Get("/validate", Validate)

	testToken, err := lib.EncryptJWT(testUsername, testUserId, testMail)

	if err != nil {
		t.Error("error making test token")
	}

	req := httptest.NewRequest("GET", "/validate", strings.NewReader(""))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testToken))

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, resp.Header.Get("Authorization"))
	assert.NotEqual(t, testToken, resp.Header.Get("Authorization"))

	invalidToken := "some_random_string"

	invalidReq := httptest.NewRequest("GET", "/validate", strings.NewReader(""))
	invalidReq.Header.Set("Content-Type", "application/json")
	invalidReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", invalidToken))

	resp, err = app.Test(invalidReq)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
}
