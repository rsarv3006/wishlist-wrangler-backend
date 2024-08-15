package repository

import (
	"context"
	"errors"
	"wishlist-wrangler-api/dto"
	"wishlist-wrangler-api/ent"
	"wishlist-wrangler-api/ent/loginrequest"
	"wishlist-wrangler-api/helper"
)

func CreateLoginRequest(dbClient *ent.Client, createLoginRequestDto *dto.CreateLoginRequest) (*ent.LoginRequest, error) {
	loginCode, err := generateUniqueLoginCode(dbClient)

	if err != nil {
		return nil, err
	}

	loginRequest, err := dbClient.LoginRequest.Create().SetUserId(createLoginRequestDto.UserId).SetLoginRequestCode(loginCode).Save(context.Background())

	if err != nil {
		return nil, err
	}

	return loginRequest, nil
}

func generateUniqueLoginCode(dbClient *ent.Client) (string, error) {
	for i := 0; i < 10; i++ {
		loginCode := helper.GenerateRandomCode()

		doesCodeExist, err := dbClient.LoginRequest.Query().Where(loginrequest.LoginRequestCode(loginCode)).Exist(context.Background())

		if err != nil {
			return "", err
		}

		if !doesCodeExist {
			return loginCode, nil
		}
	}

	return "", errors.New("Unable to generate unique login code")
}

func GetLoginRequestByCode(dbClient *ent.Client, loginCode string) (*ent.LoginRequest, error) {
	loginRequest, err := dbClient.LoginRequest.Query().Where(
		loginrequest.And(
			loginrequest.LoginRequestCode(loginCode),
			// TODO finish this
			// loginrequest.Status(),
		),
	).First(context.Background())

	if err != nil {
		return nil, err
	}

	return loginRequest, nil
}
