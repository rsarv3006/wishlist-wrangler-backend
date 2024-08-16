package handler

import (
	"wishlist-wrangler-api/dto"
	"wishlist-wrangler-api/ent"
	"wishlist-wrangler-api/helper"
	"wishlist-wrangler-api/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func SignUpEndpoint(dbClient *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		createUserDto := new(dto.CreateUserDto)

		if err := c.BodyParser(createUserDto); err != nil {
			return sendBadRequestResponse(c, err, "Failed to parse create user data")
		}

		isEmailValid := helper.DidUserEmailPassValidation(createUserDto.Email)

		if !isEmailValid {
			log.Warn("User submitted an invalid email address", createUserDto.Email)
			return sendBadRequestResponse(c, nil, "Email is not valid")
		}

		isDisplayNameValid := helper.DidUserDisplayNamePassValidation(createUserDto.DisplayName)

		if !isDisplayNameValid {
			log.Warn("User submitted an invalid email address", createUserDto.DisplayName)
			return sendBadRequestResponse(c, nil, "DisplayName is not valid")
		}

		createdUser, err := repository.CreateUser(dbClient, *createUserDto)

		if err != nil {
			if ent.IsConstraintError(err) {
				return sendBadRequestResponse(c, err, "Email is already in use")
			}
			return sendInternalServerErrorResponse(c, err)
		}

		loginRequest, err := repository.CreateLoginRequest(dbClient, &dto.CreateLoginRequest{
			UserId: createdUser.ID,
			Email:  createUserDto.Email,
		})

		if err != nil {
			return sendInternalServerErrorResponse(c, err)
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message":        "User Created",
			"loginRequestId": loginRequest.ID,
		})
	}
}

func LoginEndpoint(dbClient *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		loginRequestDto := new(dto.UserLoginDto)

		if err := c.BodyParser(loginRequestDto); err != nil {
			return sendBadRequestResponse(c, err, "Failed to parse login data")
		}

		user, err := repository.GetUserByEmail(dbClient, loginRequestDto.Email)

		if err != nil {
			return sendInternalServerErrorResponse(c, err)
		}

		loginRequest, err := repository.CreateLoginRequest(dbClient, &dto.CreateLoginRequest{
			UserId: user.ID,
			Email:  loginRequestDto.Email,
		})

		if err != nil {
			return sendInternalServerErrorResponse(c, err)
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"loginRequestId": loginRequest.ID,
		})
	}
}

func UserGetTokenEndpoint(dbClient *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userTokenDto := new(dto.UserTokenDto)

		if err := c.BodyParser(userTokenDto); err != nil {
			return sendBadRequestResponse(c, err, "Failed to parse user token data")
		}

		user, err := repository.GetUserByEmail(dbClient, userTokenDto.Email)

		if err != nil {
			return sendInternalServerErrorResponse(c, err)
		}

		loginRequest, err := repository.GetLoginRequestByCode(dbClient, userTokenDto.LoginCode, user.Email)

		if err != nil {
			return sendInternalServerErrorResponse(c, err)
		}

		if loginRequest.Email != user.Email {
			return sendBadRequestResponse(c, nil, "Login request id does not match")
		}

		// TODO: Create token

		if err != nil {
			return sendInternalServerErrorResponse(c, err)
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"token": "TODO: Create token",
		})
	}
}
