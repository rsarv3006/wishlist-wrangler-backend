// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
	"wishlist-wrangler-api/ent/loginrequest"
	"wishlist-wrangler-api/ent/schema"
	"wishlist-wrangler-api/ent/user"
	"wishlist-wrangler-api/ent/wishlist"
	"wishlist-wrangler-api/ent/wishlistsection"
	"wishlist-wrangler-api/ent/wishlisttemplate"
	"wishlist-wrangler-api/ent/wishlisttemplatesection"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	loginrequestFields := schema.LoginRequest{}.Fields()
	_ = loginrequestFields
	// loginrequestDescCreatedAt is the schema descriptor for created_at field.
	loginrequestDescCreatedAt := loginrequestFields[3].Descriptor()
	// loginrequest.DefaultCreatedAt holds the default value on creation for the created_at field.
	loginrequest.DefaultCreatedAt = loginrequestDescCreatedAt.Default.(func() time.Time)
	// loginrequestDescID is the schema descriptor for id field.
	loginrequestDescID := loginrequestFields[0].Descriptor()
	// loginrequest.DefaultID holds the default value on creation for the id field.
	loginrequest.DefaultID = loginrequestDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescDisplayName is the schema descriptor for displayName field.
	userDescDisplayName := userFields[1].Descriptor()
	// user.DisplayNameValidator is a validator for the "displayName" field. It is called by the builders before save.
	user.DisplayNameValidator = func() func(string) error {
		validators := userDescDisplayName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(displayName string) error {
			for _, fn := range fns {
				if err := fn(displayName); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = func() func(string) error {
		validators := userDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
	wishlistFields := schema.Wishlist{}.Fields()
	_ = wishlistFields
	// wishlistDescTitle is the schema descriptor for title field.
	wishlistDescTitle := wishlistFields[1].Descriptor()
	// wishlist.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	wishlist.TitleValidator = func() func(string) error {
		validators := wishlistDescTitle.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(title string) error {
			for _, fn := range fns {
				if err := fn(title); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// wishlistDescCreatedAt is the schema descriptor for created_at field.
	wishlistDescCreatedAt := wishlistFields[2].Descriptor()
	// wishlist.DefaultCreatedAt holds the default value on creation for the created_at field.
	wishlist.DefaultCreatedAt = wishlistDescCreatedAt.Default.(func() time.Time)
	// wishlistDescID is the schema descriptor for id field.
	wishlistDescID := wishlistFields[0].Descriptor()
	// wishlist.DefaultID holds the default value on creation for the id field.
	wishlist.DefaultID = wishlistDescID.Default.(func() uuid.UUID)
	wishlistsectionFields := schema.WishlistSection{}.Fields()
	_ = wishlistsectionFields
	// wishlistsectionDescValue is the schema descriptor for value field.
	wishlistsectionDescValue := wishlistsectionFields[1].Descriptor()
	// wishlistsection.ValueValidator is a validator for the "value" field. It is called by the builders before save.
	wishlistsection.ValueValidator = func() func(string) error {
		validators := wishlistsectionDescValue.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(value string) error {
			for _, fn := range fns {
				if err := fn(value); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// wishlistsectionDescCreatedAt is the schema descriptor for created_at field.
	wishlistsectionDescCreatedAt := wishlistsectionFields[2].Descriptor()
	// wishlistsection.DefaultCreatedAt holds the default value on creation for the created_at field.
	wishlistsection.DefaultCreatedAt = wishlistsectionDescCreatedAt.Default.(func() time.Time)
	// wishlistsectionDescID is the schema descriptor for id field.
	wishlistsectionDescID := wishlistsectionFields[0].Descriptor()
	// wishlistsection.DefaultID holds the default value on creation for the id field.
	wishlistsection.DefaultID = wishlistsectionDescID.Default.(func() uuid.UUID)
	wishlisttemplateFields := schema.WishlistTemplate{}.Fields()
	_ = wishlisttemplateFields
	// wishlisttemplateDescTitle is the schema descriptor for title field.
	wishlisttemplateDescTitle := wishlisttemplateFields[1].Descriptor()
	// wishlisttemplate.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	wishlisttemplate.TitleValidator = func() func(string) error {
		validators := wishlisttemplateDescTitle.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(title string) error {
			for _, fn := range fns {
				if err := fn(title); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// wishlisttemplateDescCreatedAt is the schema descriptor for created_at field.
	wishlisttemplateDescCreatedAt := wishlisttemplateFields[2].Descriptor()
	// wishlisttemplate.DefaultCreatedAt holds the default value on creation for the created_at field.
	wishlisttemplate.DefaultCreatedAt = wishlisttemplateDescCreatedAt.Default.(func() time.Time)
	// wishlisttemplateDescDescription is the schema descriptor for description field.
	wishlisttemplateDescDescription := wishlisttemplateFields[3].Descriptor()
	// wishlisttemplate.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	wishlisttemplate.DescriptionValidator = wishlisttemplateDescDescription.Validators[0].(func(string) error)
	// wishlisttemplateDescID is the schema descriptor for id field.
	wishlisttemplateDescID := wishlisttemplateFields[0].Descriptor()
	// wishlisttemplate.DefaultID holds the default value on creation for the id field.
	wishlisttemplate.DefaultID = wishlisttemplateDescID.Default.(func() uuid.UUID)
	wishlisttemplatesectionFields := schema.WishlistTemplateSection{}.Fields()
	_ = wishlisttemplatesectionFields
	// wishlisttemplatesectionDescTitle is the schema descriptor for title field.
	wishlisttemplatesectionDescTitle := wishlisttemplatesectionFields[1].Descriptor()
	// wishlisttemplatesection.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	wishlisttemplatesection.TitleValidator = func() func(string) error {
		validators := wishlisttemplatesectionDescTitle.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(title string) error {
			for _, fn := range fns {
				if err := fn(title); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// wishlisttemplatesectionDescCreatedAt is the schema descriptor for created_at field.
	wishlisttemplatesectionDescCreatedAt := wishlisttemplatesectionFields[2].Descriptor()
	// wishlisttemplatesection.DefaultCreatedAt holds the default value on creation for the created_at field.
	wishlisttemplatesection.DefaultCreatedAt = wishlisttemplatesectionDescCreatedAt.Default.(func() time.Time)
	// wishlisttemplatesectionDescSectionId is the schema descriptor for sectionId field.
	wishlisttemplatesectionDescSectionId := wishlisttemplatesectionFields[4].Descriptor()
	// wishlisttemplatesection.SectionIdValidator is a validator for the "sectionId" field. It is called by the builders before save.
	wishlisttemplatesection.SectionIdValidator = func() func(string) error {
		validators := wishlisttemplatesectionDescSectionId.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(sectionId string) error {
			for _, fn := range fns {
				if err := fn(sectionId); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// wishlisttemplatesectionDescSortOrder is the schema descriptor for sort_order field.
	wishlisttemplatesectionDescSortOrder := wishlisttemplatesectionFields[6].Descriptor()
	// wishlisttemplatesection.SortOrderValidator is a validator for the "sort_order" field. It is called by the builders before save.
	wishlisttemplatesection.SortOrderValidator = wishlisttemplatesectionDescSortOrder.Validators[0].(func(int) error)
	// wishlisttemplatesectionDescID is the schema descriptor for id field.
	wishlisttemplatesectionDescID := wishlisttemplatesectionFields[0].Descriptor()
	// wishlisttemplatesection.DefaultID holds the default value on creation for the id field.
	wishlisttemplatesection.DefaultID = wishlisttemplatesectionDescID.Default.(func() uuid.UUID)
}
