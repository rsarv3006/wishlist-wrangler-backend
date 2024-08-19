// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// LoginRequestsColumns holds the columns for the "login_requests" table.
	LoginRequestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "email", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "login_request_code", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"PENDING", "COMPLETED", "EXPIRED"}, Default: "PENDING"},
	}
	// LoginRequestsTable holds the schema information for the "login_requests" table.
	LoginRequestsTable = &schema.Table{
		Name:       "login_requests",
		Columns:    LoginRequestsColumns,
		PrimaryKey: []*schema.Column{LoginRequestsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "display_name", Type: field.TypeString, Size: 255},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 255},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"ACTIVE", "PENDING", "DELETED"}, Default: "PENDING"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// WishlistsColumns holds the columns for the "wishlists" table.
	WishlistsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString, Size: 255},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"PENDING", "ACTIVE", "REMOVED", "COMPLETED"}, Default: "PENDING"},
	}
	// WishlistsTable holds the schema information for the "wishlists" table.
	WishlistsTable = &schema.Table{
		Name:       "wishlists",
		Columns:    WishlistsColumns,
		PrimaryKey: []*schema.Column{WishlistsColumns[0]},
	}
	// WishlistSectionsColumns holds the columns for the "wishlist_sections" table.
	WishlistSectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"WISHLIST_SECTION_TYPE_TEXT", "WISHLIST_SECTION_TYPE_IMAGE", "WISHLIST_SECTION_TYPE_VIDEO", "WISHLIST_SECTION_TYPE_LINK", "WISHLIST_SECTION_TYPE_BOOLEAN"}},
		{Name: "text_value", Type: field.TypeString, Size: 1024},
		{Name: "created_at", Type: field.TypeTime},
	}
	// WishlistSectionsTable holds the schema information for the "wishlist_sections" table.
	WishlistSectionsTable = &schema.Table{
		Name:       "wishlist_sections",
		Columns:    WishlistSectionsColumns,
		PrimaryKey: []*schema.Column{WishlistSectionsColumns[0]},
	}
	// WishlistTemplatesColumns holds the columns for the "wishlist_templates" table.
	WishlistTemplatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString, Size: 255},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "description", Type: field.TypeString, Size: 525},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"ACTIVE", "REMOVED"}, Default: "ACTIVE"},
		{Name: "creator_id", Type: field.TypeUUID},
	}
	// WishlistTemplatesTable holds the schema information for the "wishlist_templates" table.
	WishlistTemplatesTable = &schema.Table{
		Name:       "wishlist_templates",
		Columns:    WishlistTemplatesColumns,
		PrimaryKey: []*schema.Column{WishlistTemplatesColumns[0]},
	}
	// WishlistTemplateSectionsColumns holds the columns for the "wishlist_template_sections" table.
	WishlistTemplateSectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString, Size: 255},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "wishlist_template_id", Type: field.TypeUUID},
	}
	// WishlistTemplateSectionsTable holds the schema information for the "wishlist_template_sections" table.
	WishlistTemplateSectionsTable = &schema.Table{
		Name:       "wishlist_template_sections",
		Columns:    WishlistTemplateSectionsColumns,
		PrimaryKey: []*schema.Column{WishlistTemplateSectionsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		LoginRequestsTable,
		UsersTable,
		WishlistsTable,
		WishlistSectionsTable,
		WishlistTemplatesTable,
		WishlistTemplateSectionsTable,
	}
)

func init() {
}
