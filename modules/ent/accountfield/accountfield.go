// Code generated by ent, DO NOT EDIT.

package accountfield

const (
	// Label holds the string label denoting the accountfield type in the database.
	Label = "account_field"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFid holds the string denoting the fid field in the database.
	FieldFid = "fid"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// EdgeAccount holds the string denoting the account edge name in mutations.
	EdgeAccount = "account"
	// Table holds the table name of the accountfield in the database.
	Table = "account_fields"
	// AccountTable is the table that holds the account relation/edge.
	AccountTable = "account_fields"
	// AccountInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AccountInverseTable = "accounts"
	// AccountColumn is the table column denoting the account relation/edge.
	AccountColumn = "account_fields"
)

// Columns holds all SQL columns for accountfield fields.
var Columns = []string{
	FieldID,
	FieldFid,
	FieldKey,
	FieldValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "account_fields"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"account_fields",
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
