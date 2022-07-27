// Code generated by ent, DO NOT EDIT.

package store

import (
	"tammy/ent"
	"tammy/pkg/store/user"
	"time"

	"entgo.io/ent"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := ent.User{}.Fields()
	_ = userFields
	// userDescIsActive is the schema descriptor for isActive field.
	userDescIsActive := userFields[1].Descriptor()
	// user.DefaultIsActive holds the default value on creation for the isActive field.
	user.DefaultIsActive = userDescIsActive.Default.(bool)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[2].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[3].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
}