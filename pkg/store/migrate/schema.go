// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "user_id", Type: field.TypeUint32},
		{Name: "is_active", Type: field.TypeBool},
		{Name: "is_blocked", Type: field.TypeBool},
		{Name: "user_accounts", Type: field.TypeUint32},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "accounts_users_accounts",
				Columns:    []*schema.Column{AccountsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// AccountFieldsColumns holds the columns for the "account_fields" table.
	AccountFieldsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "fid", Type: field.TypeUint32},
		{Name: "key", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "account_fields", Type: field.TypeUint32},
	}
	// AccountFieldsTable holds the schema information for the "account_fields" table.
	AccountFieldsTable = &schema.Table{
		Name:       "account_fields",
		Columns:    AccountFieldsColumns,
		PrimaryKey: []*schema.Column{AccountFieldsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "account_fields_accounts_fields",
				Columns:    []*schema.Column{AccountFieldsColumns[4]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PortalsColumns holds the columns for the "portals" table.
	PortalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "is_active", Type: field.TypeBool, Default: false},
		{Name: "slug", Type: field.TypeString, Unique: true},
	}
	// PortalsTable holds the schema information for the "portals" table.
	PortalsTable = &schema.Table{
		Name:       "portals",
		Columns:    PortalsColumns,
		PrimaryKey: []*schema.Column{PortalsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "portal_slug",
				Unique:  false,
				Columns: []*schema.Column{PortalsColumns[4]},
			},
		},
	}
	// PortalLegalsColumns holds the columns for the "portal_legals" table.
	PortalLegalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "privacy_policy", Type: field.TypeString},
		{Name: "term_of_service", Type: field.TypeString},
		{Name: "copyright", Type: field.TypeString},
		{Name: "online_training_agreement", Type: field.TypeString},
		{Name: "portal_legal", Type: field.TypeUint32},
	}
	// PortalLegalsTable holds the schema information for the "portal_legals" table.
	PortalLegalsTable = &schema.Table{
		Name:       "portal_legals",
		Columns:    PortalLegalsColumns,
		PrimaryKey: []*schema.Column{PortalLegalsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "portal_legals_portals_legal",
				Columns:    []*schema.Column{PortalLegalsColumns[7]},
				RefColumns: []*schema.Column{PortalsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PortalMetadataColumns holds the columns for the "portal_metadata" table.
	PortalMetadataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "kind", Type: field.TypeEnum, Enums: []string{"KIND_DEVELOPMENT", "KIND_INTERNAL", "KIND_CONTENT_PARTNER", "KIND_DISTRIBUTION_PARTNER", "KIND_CUSTOMER", "KIND_COMPLISPACE", "KIND_TEAM"}, Default: "KIND_DEVELOPMENT"},
		{Name: "lifecycle", Type: field.TypeEnum, Enums: []string{"LIFECYCLE_TRIAL", "LIFECYCLE_ONBOARD", "LIFECYCLE_LIVE", "LIFECYCLE_EXPIRED_TRIAL", "LIFECYCLE_SUSPENDED", "LIFECYCLE_CANCELLED", "LIFECYCLE_TEST", "LIFECYCLE_SAMPLE", "LIFECYCLE_TEMPLATE"}, Default: "LIFECYCLE_TRIAL"},
		{Name: "portal_metadata", Type: field.TypeUint32},
	}
	// PortalMetadataTable holds the schema information for the "portal_metadata" table.
	PortalMetadataTable = &schema.Table{
		Name:       "portal_metadata",
		Columns:    PortalMetadataColumns,
		PrimaryKey: []*schema.Column{PortalMetadataColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "portal_metadata_portals_metadata",
				Columns:    []*schema.Column{PortalMetadataColumns[5]},
				RefColumns: []*schema.Column{PortalsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserEmailsColumns holds the columns for the "user_emails" table.
	UserEmailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "value", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "is_verified", Type: field.TypeBool},
		{Name: "user_emails", Type: field.TypeUint32, Nullable: true},
	}
	// UserEmailsTable holds the schema information for the "user_emails" table.
	UserEmailsTable = &schema.Table{
		Name:       "user_emails",
		Columns:    UserEmailsColumns,
		PrimaryKey: []*schema.Column{UserEmailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_emails_users_emails",
				Columns:    []*schema.Column{UserEmailsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserPasswordsColumns holds the columns for the "user_passwords" table.
	UserPasswordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "user_id", Type: field.TypeUint32},
		{Name: "hashed_value", Type: field.TypeString},
		{Name: "is_active", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "user_password", Type: field.TypeUint32, Unique: true},
	}
	// UserPasswordsTable holds the schema information for the "user_passwords" table.
	UserPasswordsTable = &schema.Table{
		Name:       "user_passwords",
		Columns:    UserPasswordsColumns,
		PrimaryKey: []*schema.Column{UserPasswordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_passwords_users_password",
				Columns:    []*schema.Column{UserPasswordsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PortalMembersColumns holds the columns for the "portal_members" table.
	PortalMembersColumns = []*schema.Column{
		{Name: "portal_id", Type: field.TypeUint32},
		{Name: "account_id", Type: field.TypeUint32},
	}
	// PortalMembersTable holds the schema information for the "portal_members" table.
	PortalMembersTable = &schema.Table{
		Name:       "portal_members",
		Columns:    PortalMembersColumns,
		PrimaryKey: []*schema.Column{PortalMembersColumns[0], PortalMembersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "portal_members_portal_id",
				Columns:    []*schema.Column{PortalMembersColumns[0]},
				RefColumns: []*schema.Column{PortalsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "portal_members_account_id",
				Columns:    []*schema.Column{PortalMembersColumns[1]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		AccountFieldsTable,
		PortalsTable,
		PortalLegalsTable,
		PortalMetadataTable,
		UsersTable,
		UserEmailsTable,
		UserPasswordsTable,
		PortalMembersTable,
	}
)

func init() {
	AccountsTable.ForeignKeys[0].RefTable = UsersTable
	AccountFieldsTable.ForeignKeys[0].RefTable = AccountsTable
	PortalLegalsTable.ForeignKeys[0].RefTable = PortalsTable
	PortalMetadataTable.ForeignKeys[0].RefTable = PortalsTable
	UserEmailsTable.ForeignKeys[0].RefTable = UsersTable
	UserPasswordsTable.ForeignKeys[0].RefTable = UsersTable
	PortalMembersTable.ForeignKeys[0].RefTable = PortalsTable
	PortalMembersTable.ForeignKeys[1].RefTable = AccountsTable
}
