package repository

import (
	"context"
	"errors"
	"wishlist-wrangler-api/dto"
	"wishlist-wrangler-api/ent"
	"wishlist-wrangler-api/ent/user"
	"wishlist-wrangler-api/ent/wishlisttemplate"

	"github.com/google/uuid"
)

func CreateWishlistTemplate(dbClient *ent.Client, templateDto dto.CreateWishlistTemplateDto, creator *ent.User) (*ent.WishlistTemplate, error) {
	sections, err := CreateWishlistTemplateSections(dbClient, templateDto.TemplateSections)

	if err != nil {
		return nil, err
	}

	if sections == nil {
		return nil, errors.New("Unable to create wishlist template from an empty array of sections")
	}

	template, err := dbClient.WishlistTemplate.Create().
		SetTitle(templateDto.Title).
		SetDescription(templateDto.Description).
		AddSections(sections...).
		AddCreator(creator).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return dbClient.WishlistTemplate.Query().
		Where(wishlisttemplate.ID(template.ID)).
		WithSections().
		Only(context.Background())
}

func GetWishlistTemplate(dbClient *ent.Client, id uuid.UUID) (*ent.WishlistTemplate, error) {
	return dbClient.WishlistTemplate.Query().
		Where(wishlisttemplate.ID(id)).
		WithSections().
		Only(context.Background())
}

func GetWishlistTemplatesByUser(dbClient *ent.Client, userId uuid.UUID) ([]*ent.WishlistTemplate, error) {
	return dbClient.WishlistTemplate.
		Query().
		Where(
			wishlisttemplate.And(
				wishlisttemplate.HasCreatorWith(user.ID(userId)),
				wishlisttemplate.StatusEQ(wishlisttemplate.StatusACTIVE),
			),
		).
		WithSections().
		All(context.Background())
}
