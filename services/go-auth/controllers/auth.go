package controllers

import (
	"authService/lib"
	"authService/models"
	b64 "encoding/base64"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type (
	SignUpReqBody struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginReqBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func SignUp(c *fiber.Ctx) error {
	var body SignUpReqBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if body.Email == "" || body.Password == "" || body.Username == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if user, err := models.MakeUser(body.Username, body.Email, body.Password); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err.Error(),
		})
	} else {

		token, err := lib.EncryptJWT(user.Username, user.Id.Hex(), user.Email)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"err": err.Error(),
			})
		}

		c.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return c.Status(fiber.StatusCreated).JSON(user)
	}
}

func Login(c *fiber.Ctx) error {
	var authHeader string = c.Get("Authorization")
	if authHeader == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	authParts := strings.Split(authHeader, "Basic ")
	if len(authParts) != 2 {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	authString := authParts[1]
	
	decodedAuth, err := b64.StdEncoding.DecodeString(authString)
	if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"err": "Invalid base64 encoding",
			})
	}

	authParts = strings.Split(string(decodedAuth), ":")
	if len(authParts) != 2 {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	email := authParts[0]
	password := authParts[1]

	user, err := models.GetUserByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	if ok := lib.ValidatePSWHash(password, user.Password); !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "bad credentials",
		})
	}

	token, err := lib.EncryptJWT(user.Username, user.Id.Hex(), user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	c.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	return c.Status(fiber.StatusOK).JSON(models.ToLogedInUser(user, token))
}

func Validate(c *fiber.Ctx) error {

	headerToken := c.Get("Authorization")
	if headerToken == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	headerParts := strings.Split(headerToken, "Bearer ")
	if len(headerParts) != 2 {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	jwtToken := headerParts[1]

	claims, err := lib.ValidateJWT(jwtToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	user := models.ClaimToAuthorizedUser(*claims)

	token, err := lib.EncryptJWT(claims.Username, claims.Id, claims.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": "User not found",
		})
	}


	c.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	return c.Status(fiber.StatusOK).JSON(user)
}
