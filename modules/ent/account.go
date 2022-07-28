// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"tammy/ent/account"
	"tammy/ent/user"

	"entgo.io/ent/dialect/sql"
)

// Account is the model entity for the Account schema.
type Account struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// IsActive holds the value of the "isActive" field.
	IsActive bool `json:"isActive,omitempty"`
	// IsBlocked holds the value of the "isBlocked" field.
	IsBlocked bool `json:"isBlocked,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountQuery when eager-loading is set.
	Edges         AccountEdges `json:"edges"`
	user_accounts *int
}

// AccountEdges holds the relations/edges for other nodes in the graph.
type AccountEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Portal holds the value of the portal edge.
	Portal []*Portal `json:"portal,omitempty"`
	// Fields holds the value of the fields edge.
	Fields []*AccountField `json:"fields,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// PortalOrErr returns the Portal value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) PortalOrErr() ([]*Portal, error) {
	if e.loadedTypes[1] {
		return e.Portal, nil
	}
	return nil, &NotLoadedError{edge: "portal"}
}

// FieldsOrErr returns the Fields value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) FieldsOrErr() ([]*AccountField, error) {
	if e.loadedTypes[2] {
		return e.Fields, nil
	}
	return nil, &NotLoadedError{edge: "fields"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Account) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case account.FieldIsActive, account.FieldIsBlocked:
			values[i] = new(sql.NullBool)
		case account.FieldID:
			values[i] = new(sql.NullInt64)
		case account.ForeignKeys[0]: // user_accounts
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Account", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Account fields.
func (a *Account) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case account.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case account.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field isActive", values[i])
			} else if value.Valid {
				a.IsActive = value.Bool
			}
		case account.FieldIsBlocked:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field isBlocked", values[i])
			} else if value.Valid {
				a.IsBlocked = value.Bool
			}
		case account.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_accounts", value)
			} else if value.Valid {
				a.user_accounts = new(int)
				*a.user_accounts = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Account entity.
func (a *Account) QueryUser() *UserQuery {
	return (&AccountClient{config: a.config}).QueryUser(a)
}

// QueryPortal queries the "portal" edge of the Account entity.
func (a *Account) QueryPortal() *PortalQuery {
	return (&AccountClient{config: a.config}).QueryPortal(a)
}

// QueryFields queries the "fields" edge of the Account entity.
func (a *Account) QueryFields() *AccountFieldQuery {
	return (&AccountClient{config: a.config}).QueryFields(a)
}

// Update returns a builder for updating this Account.
// Note that you need to call Account.Unwrap() before calling this method if this Account
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Account) Update() *AccountUpdateOne {
	return (&AccountClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Account entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Account) Unwrap() *Account {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Account is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Account) String() string {
	var builder strings.Builder
	builder.WriteString("Account(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("isActive=")
	builder.WriteString(fmt.Sprintf("%v", a.IsActive))
	builder.WriteString(", ")
	builder.WriteString("isBlocked=")
	builder.WriteString(fmt.Sprintf("%v", a.IsBlocked))
	builder.WriteByte(')')
	return builder.String()
}

// Accounts is a parsable slice of Account.
type Accounts []*Account

func (a Accounts) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
