package coresharedcontroller

import "strings"

func isForeignKeyViolation(err error) bool {
	// This is an example check. You need to adjust it based on your database driver and error message.
	return strings.Contains(err.Error(), "violates foreign key constraint")
}
