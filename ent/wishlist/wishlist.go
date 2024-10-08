// Code generated by ent, DO NOT EDIT.

package wishlist

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the wishlist type in the database.
	Label = "wishlist"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldTemplateID holds the string denoting the template_id field in the database.
	FieldTemplateID = "template_id"
	// FieldCreatorID holds the string denoting the creator_id field in the database.
	FieldCreatorID = "creator_id"
	// Table holds the table name of the wishlist in the database.
	Table = "wishlists"
)

// Columns holds all SQL columns for wishlist fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldCreatedAt,
	FieldStatus,
	FieldTemplateID,
	FieldCreatorID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// StatusPENDING is the default value of the Status enum.
const DefaultStatus = StatusPENDING

// Status values.
const (
	StatusPENDING   Status = "PENDING"
	StatusACTIVE    Status = "ACTIVE"
	StatusREMOVED   Status = "REMOVED"
	StatusCOMPLETED Status = "COMPLETED"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusPENDING, StatusACTIVE, StatusREMOVED, StatusCOMPLETED:
		return nil
	default:
		return fmt.Errorf("wishlist: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Wishlist queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByTemplateID orders the results by the template_id field.
func ByTemplateID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTemplateID, opts...).ToFunc()
}

// ByCreatorID orders the results by the creator_id field.
func ByCreatorID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatorID, opts...).ToFunc()
}
