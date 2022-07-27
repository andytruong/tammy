// Code generated by ent, DO NOT EDIT.

package account

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserId holds the string denoting the userid field in the database.
	FieldUserId = "user_id"
	// FieldIsActive holds the string denoting the isactive field in the database.
	FieldIsActive = "is_active"
	// FieldIsBlocked holds the string denoting the isblocked field in the database.
	FieldIsBlocked = "is_blocked"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeFields holds the string denoting the fields edge name in mutations.
	EdgeFields = "fields"
	// Table holds the table name of the account in the database.
	Table = "accounts"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "accounts"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_accounts"
	// FieldsTable is the table that holds the fields relation/edge.
	FieldsTable = "account_fields"
	// FieldsInverseTable is the table name for the AccountField entity.
	// It exists in this package in order to avoid circular dependency with the "accountfield" package.
	FieldsInverseTable = "account_fields"
	// FieldsColumn is the table column denoting the fields relation/edge.
	FieldsColumn = "account_fields"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldUserId,
	FieldIsActive,
	FieldIsBlocked,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "accounts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_accounts",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}