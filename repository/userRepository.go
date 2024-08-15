package repository

import (
	"context"
	"errors"
	"wishlist-wrangler-api/dto"
	"wishlist-wrangler-api/ent"
	"wishlist-wrangler-api/ent/user"
	"wishlist-wrangler-api/helper"

	"github.com/google/uuid"
)

func CreateUser(dbClient *ent.Client, createUserDto dto.CreateUserDto) (*ent.User, error) {
	user, err := dbClient.User.Create().SetEmail(createUserDto.Email).SetDisplayName(createUserDto.DisplayName).Save(context.Background())

	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserById(dbClient *ent.Client, userId uuid.UUID) (*ent.User, error) {
	user, err := dbClient.User.Get(context.Background(), userId)

	return user, err
}

func GetUserByEmail(dbClient *ent.Client, email string) (*ent.User, error) {
	user, err := dbClient.User.Query().Where(user.Email(email)).First(context.Background())
	return user, err
}

func DeleteUser(dbClient *ent.Client, userId uuid.UUID) error {
	return dbClient.User.Update().Where(user.ID(userId)).SetStatus(user.StatusDELETED).Exec(context.Background())
}

func UpdateUser(dbClient *ent.Client, updateUserDto *dto.UpdateUserDto) error {
	if updateUserDto == nil {
		return errors.New("Unable to update user from an empty object")
	}

	isDisplayNameValid := helper.DidUserDisplayNamePassValidation(updateUserDto.DisplayName)

	if !isDisplayNameValid {
		return errors.New("Display name is invalid")
	}

	err := dbClient.User.Update().Where(user.ID(updateUserDto.UserId)).SetDisplayName(updateUserDto.DisplayName).Exec(context.Background())

	return err
}
