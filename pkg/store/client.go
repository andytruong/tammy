// Code generated by ent, DO NOT EDIT.

package store

import (
	"context"
	"fmt"
	"log"

	"tammy/pkg/store/migrate"

	"tammy/pkg/store/account"
	"tammy/pkg/store/accountfield"
	"tammy/pkg/store/portal"
	"tammy/pkg/store/portallegal"
	"tammy/pkg/store/portalmetadata"
	"tammy/pkg/store/user"
	"tammy/pkg/store/useremail"
	"tammy/pkg/store/userpassword"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Account is the client for interacting with the Account builders.
	Account *AccountClient
	// AccountField is the client for interacting with the AccountField builders.
	AccountField *AccountFieldClient
	// Portal is the client for interacting with the Portal builders.
	Portal *PortalClient
	// PortalLegal is the client for interacting with the PortalLegal builders.
	PortalLegal *PortalLegalClient
	// PortalMetadata is the client for interacting with the PortalMetadata builders.
	PortalMetadata *PortalMetadataClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// UserEmail is the client for interacting with the UserEmail builders.
	UserEmail *UserEmailClient
	// UserPassword is the client for interacting with the UserPassword builders.
	UserPassword *UserPasswordClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Account = NewAccountClient(c.config)
	c.AccountField = NewAccountFieldClient(c.config)
	c.Portal = NewPortalClient(c.config)
	c.PortalLegal = NewPortalLegalClient(c.config)
	c.PortalMetadata = NewPortalMetadataClient(c.config)
	c.User = NewUserClient(c.config)
	c.UserEmail = NewUserEmailClient(c.config)
	c.UserPassword = NewUserPasswordClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("store: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("store: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:            ctx,
		config:         cfg,
		Account:        NewAccountClient(cfg),
		AccountField:   NewAccountFieldClient(cfg),
		Portal:         NewPortalClient(cfg),
		PortalLegal:    NewPortalLegalClient(cfg),
		PortalMetadata: NewPortalMetadataClient(cfg),
		User:           NewUserClient(cfg),
		UserEmail:      NewUserEmailClient(cfg),
		UserPassword:   NewUserPasswordClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:            ctx,
		config:         cfg,
		Account:        NewAccountClient(cfg),
		AccountField:   NewAccountFieldClient(cfg),
		Portal:         NewPortalClient(cfg),
		PortalLegal:    NewPortalLegalClient(cfg),
		PortalMetadata: NewPortalMetadataClient(cfg),
		User:           NewUserClient(cfg),
		UserEmail:      NewUserEmailClient(cfg),
		UserPassword:   NewUserPasswordClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Account.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Account.Use(hooks...)
	c.AccountField.Use(hooks...)
	c.Portal.Use(hooks...)
	c.PortalLegal.Use(hooks...)
	c.PortalMetadata.Use(hooks...)
	c.User.Use(hooks...)
	c.UserEmail.Use(hooks...)
	c.UserPassword.Use(hooks...)
}

// AccountClient is a client for the Account schema.
type AccountClient struct {
	config
}

// NewAccountClient returns a client for the Account from the given config.
func NewAccountClient(c config) *AccountClient {
	return &AccountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `account.Hooks(f(g(h())))`.
func (c *AccountClient) Use(hooks ...Hook) {
	c.hooks.Account = append(c.hooks.Account, hooks...)
}

// Create returns a builder for creating a Account entity.
func (c *AccountClient) Create() *AccountCreate {
	mutation := newAccountMutation(c.config, OpCreate)
	return &AccountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Account entities.
func (c *AccountClient) CreateBulk(builders ...*AccountCreate) *AccountCreateBulk {
	return &AccountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Account.
func (c *AccountClient) Update() *AccountUpdate {
	mutation := newAccountMutation(c.config, OpUpdate)
	return &AccountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountClient) UpdateOne(a *Account) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccount(a))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountClient) UpdateOneID(id uint32) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccountID(id))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Account.
func (c *AccountClient) Delete() *AccountDelete {
	mutation := newAccountMutation(c.config, OpDelete)
	return &AccountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AccountClient) DeleteOne(a *Account) *AccountDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *AccountClient) DeleteOneID(id uint32) *AccountDeleteOne {
	builder := c.Delete().Where(account.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AccountDeleteOne{builder}
}

// Query returns a query builder for Account.
func (c *AccountClient) Query() *AccountQuery {
	return &AccountQuery{
		config: c.config,
	}
}

// Get returns a Account entity by its id.
func (c *AccountClient) Get(ctx context.Context, id uint32) (*Account, error) {
	return c.Query().Where(account.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountClient) GetX(ctx context.Context, id uint32) *Account {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Account.
func (c *AccountClient) QueryUser(a *Account) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, account.UserTable, account.UserColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPortal queries the portal edge of a Account.
func (c *AccountClient) QueryPortal(a *Account) *PortalQuery {
	query := &PortalQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, id),
			sqlgraph.To(portal.Table, portal.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, account.PortalTable, account.PortalPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFields queries the fields edge of a Account.
func (c *AccountClient) QueryFields(a *Account) *AccountFieldQuery {
	query := &AccountFieldQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, id),
			sqlgraph.To(accountfield.Table, accountfield.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, account.FieldsTable, account.FieldsColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AccountClient) Hooks() []Hook {
	return c.hooks.Account
}

// AccountFieldClient is a client for the AccountField schema.
type AccountFieldClient struct {
	config
}

// NewAccountFieldClient returns a client for the AccountField from the given config.
func NewAccountFieldClient(c config) *AccountFieldClient {
	return &AccountFieldClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `accountfield.Hooks(f(g(h())))`.
func (c *AccountFieldClient) Use(hooks ...Hook) {
	c.hooks.AccountField = append(c.hooks.AccountField, hooks...)
}

// Create returns a builder for creating a AccountField entity.
func (c *AccountFieldClient) Create() *AccountFieldCreate {
	mutation := newAccountFieldMutation(c.config, OpCreate)
	return &AccountFieldCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AccountField entities.
func (c *AccountFieldClient) CreateBulk(builders ...*AccountFieldCreate) *AccountFieldCreateBulk {
	return &AccountFieldCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AccountField.
func (c *AccountFieldClient) Update() *AccountFieldUpdate {
	mutation := newAccountFieldMutation(c.config, OpUpdate)
	return &AccountFieldUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountFieldClient) UpdateOne(af *AccountField) *AccountFieldUpdateOne {
	mutation := newAccountFieldMutation(c.config, OpUpdateOne, withAccountField(af))
	return &AccountFieldUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountFieldClient) UpdateOneID(id uint32) *AccountFieldUpdateOne {
	mutation := newAccountFieldMutation(c.config, OpUpdateOne, withAccountFieldID(id))
	return &AccountFieldUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AccountField.
func (c *AccountFieldClient) Delete() *AccountFieldDelete {
	mutation := newAccountFieldMutation(c.config, OpDelete)
	return &AccountFieldDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AccountFieldClient) DeleteOne(af *AccountField) *AccountFieldDeleteOne {
	return c.DeleteOneID(af.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *AccountFieldClient) DeleteOneID(id uint32) *AccountFieldDeleteOne {
	builder := c.Delete().Where(accountfield.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AccountFieldDeleteOne{builder}
}

// Query returns a query builder for AccountField.
func (c *AccountFieldClient) Query() *AccountFieldQuery {
	return &AccountFieldQuery{
		config: c.config,
	}
}

// Get returns a AccountField entity by its id.
func (c *AccountFieldClient) Get(ctx context.Context, id uint32) (*AccountField, error) {
	return c.Query().Where(accountfield.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountFieldClient) GetX(ctx context.Context, id uint32) *AccountField {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAccount queries the account edge of a AccountField.
func (c *AccountFieldClient) QueryAccount(af *AccountField) *AccountQuery {
	query := &AccountQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := af.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(accountfield.Table, accountfield.FieldID, id),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, accountfield.AccountTable, accountfield.AccountColumn),
		)
		fromV = sqlgraph.Neighbors(af.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AccountFieldClient) Hooks() []Hook {
	return c.hooks.AccountField
}

// PortalClient is a client for the Portal schema.
type PortalClient struct {
	config
}

// NewPortalClient returns a client for the Portal from the given config.
func NewPortalClient(c config) *PortalClient {
	return &PortalClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `portal.Hooks(f(g(h())))`.
func (c *PortalClient) Use(hooks ...Hook) {
	c.hooks.Portal = append(c.hooks.Portal, hooks...)
}

// Create returns a builder for creating a Portal entity.
func (c *PortalClient) Create() *PortalCreate {
	mutation := newPortalMutation(c.config, OpCreate)
	return &PortalCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Portal entities.
func (c *PortalClient) CreateBulk(builders ...*PortalCreate) *PortalCreateBulk {
	return &PortalCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Portal.
func (c *PortalClient) Update() *PortalUpdate {
	mutation := newPortalMutation(c.config, OpUpdate)
	return &PortalUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PortalClient) UpdateOne(po *Portal) *PortalUpdateOne {
	mutation := newPortalMutation(c.config, OpUpdateOne, withPortal(po))
	return &PortalUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PortalClient) UpdateOneID(id uint32) *PortalUpdateOne {
	mutation := newPortalMutation(c.config, OpUpdateOne, withPortalID(id))
	return &PortalUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Portal.
func (c *PortalClient) Delete() *PortalDelete {
	mutation := newPortalMutation(c.config, OpDelete)
	return &PortalDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PortalClient) DeleteOne(po *Portal) *PortalDeleteOne {
	return c.DeleteOneID(po.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *PortalClient) DeleteOneID(id uint32) *PortalDeleteOne {
	builder := c.Delete().Where(portal.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PortalDeleteOne{builder}
}

// Query returns a query builder for Portal.
func (c *PortalClient) Query() *PortalQuery {
	return &PortalQuery{
		config: c.config,
	}
}

// Get returns a Portal entity by its id.
func (c *PortalClient) Get(ctx context.Context, id uint32) (*Portal, error) {
	return c.Query().Where(portal.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PortalClient) GetX(ctx context.Context, id uint32) *Portal {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMembers queries the members edge of a Portal.
func (c *PortalClient) QueryMembers(po *Portal) *AccountQuery {
	query := &AccountQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(portal.Table, portal.FieldID, id),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, portal.MembersTable, portal.MembersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMetadata queries the metadata edge of a Portal.
func (c *PortalClient) QueryMetadata(po *Portal) *PortalMetadataQuery {
	query := &PortalMetadataQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(portal.Table, portal.FieldID, id),
			sqlgraph.To(portalmetadata.Table, portalmetadata.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, portal.MetadataTable, portal.MetadataColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLegal queries the legal edge of a Portal.
func (c *PortalClient) QueryLegal(po *Portal) *PortalLegalQuery {
	query := &PortalLegalQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(portal.Table, portal.FieldID, id),
			sqlgraph.To(portallegal.Table, portallegal.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, portal.LegalTable, portal.LegalColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PortalClient) Hooks() []Hook {
	return c.hooks.Portal
}

// PortalLegalClient is a client for the PortalLegal schema.
type PortalLegalClient struct {
	config
}

// NewPortalLegalClient returns a client for the PortalLegal from the given config.
func NewPortalLegalClient(c config) *PortalLegalClient {
	return &PortalLegalClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `portallegal.Hooks(f(g(h())))`.
func (c *PortalLegalClient) Use(hooks ...Hook) {
	c.hooks.PortalLegal = append(c.hooks.PortalLegal, hooks...)
}

// Create returns a builder for creating a PortalLegal entity.
func (c *PortalLegalClient) Create() *PortalLegalCreate {
	mutation := newPortalLegalMutation(c.config, OpCreate)
	return &PortalLegalCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PortalLegal entities.
func (c *PortalLegalClient) CreateBulk(builders ...*PortalLegalCreate) *PortalLegalCreateBulk {
	return &PortalLegalCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PortalLegal.
func (c *PortalLegalClient) Update() *PortalLegalUpdate {
	mutation := newPortalLegalMutation(c.config, OpUpdate)
	return &PortalLegalUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PortalLegalClient) UpdateOne(pl *PortalLegal) *PortalLegalUpdateOne {
	mutation := newPortalLegalMutation(c.config, OpUpdateOne, withPortalLegal(pl))
	return &PortalLegalUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PortalLegalClient) UpdateOneID(id uint32) *PortalLegalUpdateOne {
	mutation := newPortalLegalMutation(c.config, OpUpdateOne, withPortalLegalID(id))
	return &PortalLegalUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PortalLegal.
func (c *PortalLegalClient) Delete() *PortalLegalDelete {
	mutation := newPortalLegalMutation(c.config, OpDelete)
	return &PortalLegalDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PortalLegalClient) DeleteOne(pl *PortalLegal) *PortalLegalDeleteOne {
	return c.DeleteOneID(pl.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *PortalLegalClient) DeleteOneID(id uint32) *PortalLegalDeleteOne {
	builder := c.Delete().Where(portallegal.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PortalLegalDeleteOne{builder}
}

// Query returns a query builder for PortalLegal.
func (c *PortalLegalClient) Query() *PortalLegalQuery {
	return &PortalLegalQuery{
		config: c.config,
	}
}

// Get returns a PortalLegal entity by its id.
func (c *PortalLegalClient) Get(ctx context.Context, id uint32) (*PortalLegal, error) {
	return c.Query().Where(portallegal.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PortalLegalClient) GetX(ctx context.Context, id uint32) *PortalLegal {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPortal queries the portal edge of a PortalLegal.
func (c *PortalLegalClient) QueryPortal(pl *PortalLegal) *PortalQuery {
	query := &PortalQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(portallegal.Table, portallegal.FieldID, id),
			sqlgraph.To(portal.Table, portal.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, portallegal.PortalTable, portallegal.PortalColumn),
		)
		fromV = sqlgraph.Neighbors(pl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PortalLegalClient) Hooks() []Hook {
	return c.hooks.PortalLegal
}

// PortalMetadataClient is a client for the PortalMetadata schema.
type PortalMetadataClient struct {
	config
}

// NewPortalMetadataClient returns a client for the PortalMetadata from the given config.
func NewPortalMetadataClient(c config) *PortalMetadataClient {
	return &PortalMetadataClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `portalmetadata.Hooks(f(g(h())))`.
func (c *PortalMetadataClient) Use(hooks ...Hook) {
	c.hooks.PortalMetadata = append(c.hooks.PortalMetadata, hooks...)
}

// Create returns a builder for creating a PortalMetadata entity.
func (c *PortalMetadataClient) Create() *PortalMetadataCreate {
	mutation := newPortalMetadataMutation(c.config, OpCreate)
	return &PortalMetadataCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PortalMetadata entities.
func (c *PortalMetadataClient) CreateBulk(builders ...*PortalMetadataCreate) *PortalMetadataCreateBulk {
	return &PortalMetadataCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PortalMetadata.
func (c *PortalMetadataClient) Update() *PortalMetadataUpdate {
	mutation := newPortalMetadataMutation(c.config, OpUpdate)
	return &PortalMetadataUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PortalMetadataClient) UpdateOne(pm *PortalMetadata) *PortalMetadataUpdateOne {
	mutation := newPortalMetadataMutation(c.config, OpUpdateOne, withPortalMetadata(pm))
	return &PortalMetadataUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PortalMetadataClient) UpdateOneID(id uint32) *PortalMetadataUpdateOne {
	mutation := newPortalMetadataMutation(c.config, OpUpdateOne, withPortalMetadataID(id))
	return &PortalMetadataUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PortalMetadata.
func (c *PortalMetadataClient) Delete() *PortalMetadataDelete {
	mutation := newPortalMetadataMutation(c.config, OpDelete)
	return &PortalMetadataDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PortalMetadataClient) DeleteOne(pm *PortalMetadata) *PortalMetadataDeleteOne {
	return c.DeleteOneID(pm.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *PortalMetadataClient) DeleteOneID(id uint32) *PortalMetadataDeleteOne {
	builder := c.Delete().Where(portalmetadata.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PortalMetadataDeleteOne{builder}
}

// Query returns a query builder for PortalMetadata.
func (c *PortalMetadataClient) Query() *PortalMetadataQuery {
	return &PortalMetadataQuery{
		config: c.config,
	}
}

// Get returns a PortalMetadata entity by its id.
func (c *PortalMetadataClient) Get(ctx context.Context, id uint32) (*PortalMetadata, error) {
	return c.Query().Where(portalmetadata.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PortalMetadataClient) GetX(ctx context.Context, id uint32) *PortalMetadata {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPortal queries the portal edge of a PortalMetadata.
func (c *PortalMetadataClient) QueryPortal(pm *PortalMetadata) *PortalQuery {
	query := &PortalQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(portalmetadata.Table, portalmetadata.FieldID, id),
			sqlgraph.To(portal.Table, portal.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, portalmetadata.PortalTable, portalmetadata.PortalColumn),
		)
		fromV = sqlgraph.Neighbors(pm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PortalMetadataClient) Hooks() []Hook {
	return c.hooks.PortalMetadata
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uint32) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uint32) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uint32) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uint32) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPassword queries the password edge of a User.
func (c *UserClient) QueryPassword(u *User) *UserPasswordQuery {
	query := &UserPasswordQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(userpassword.Table, userpassword.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.PasswordTable, user.PasswordColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAccounts queries the accounts edge of a User.
func (c *UserClient) QueryAccounts(u *User) *AccountQuery {
	query := &AccountQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.AccountsTable, user.AccountsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEmails queries the emails edge of a User.
func (c *UserClient) QueryEmails(u *User) *UserEmailQuery {
	query := &UserEmailQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(useremail.Table, useremail.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.EmailsTable, user.EmailsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// UserEmailClient is a client for the UserEmail schema.
type UserEmailClient struct {
	config
}

// NewUserEmailClient returns a client for the UserEmail from the given config.
func NewUserEmailClient(c config) *UserEmailClient {
	return &UserEmailClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `useremail.Hooks(f(g(h())))`.
func (c *UserEmailClient) Use(hooks ...Hook) {
	c.hooks.UserEmail = append(c.hooks.UserEmail, hooks...)
}

// Create returns a builder for creating a UserEmail entity.
func (c *UserEmailClient) Create() *UserEmailCreate {
	mutation := newUserEmailMutation(c.config, OpCreate)
	return &UserEmailCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserEmail entities.
func (c *UserEmailClient) CreateBulk(builders ...*UserEmailCreate) *UserEmailCreateBulk {
	return &UserEmailCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserEmail.
func (c *UserEmailClient) Update() *UserEmailUpdate {
	mutation := newUserEmailMutation(c.config, OpUpdate)
	return &UserEmailUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserEmailClient) UpdateOne(ue *UserEmail) *UserEmailUpdateOne {
	mutation := newUserEmailMutation(c.config, OpUpdateOne, withUserEmail(ue))
	return &UserEmailUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserEmailClient) UpdateOneID(id uint32) *UserEmailUpdateOne {
	mutation := newUserEmailMutation(c.config, OpUpdateOne, withUserEmailID(id))
	return &UserEmailUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserEmail.
func (c *UserEmailClient) Delete() *UserEmailDelete {
	mutation := newUserEmailMutation(c.config, OpDelete)
	return &UserEmailDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserEmailClient) DeleteOne(ue *UserEmail) *UserEmailDeleteOne {
	return c.DeleteOneID(ue.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserEmailClient) DeleteOneID(id uint32) *UserEmailDeleteOne {
	builder := c.Delete().Where(useremail.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserEmailDeleteOne{builder}
}

// Query returns a query builder for UserEmail.
func (c *UserEmailClient) Query() *UserEmailQuery {
	return &UserEmailQuery{
		config: c.config,
	}
}

// Get returns a UserEmail entity by its id.
func (c *UserEmailClient) Get(ctx context.Context, id uint32) (*UserEmail, error) {
	return c.Query().Where(useremail.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserEmailClient) GetX(ctx context.Context, id uint32) *UserEmail {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a UserEmail.
func (c *UserEmailClient) QueryUser(ue *UserEmail) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ue.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(useremail.Table, useremail.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, useremail.UserTable, useremail.UserColumn),
		)
		fromV = sqlgraph.Neighbors(ue.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserEmailClient) Hooks() []Hook {
	return c.hooks.UserEmail
}

// UserPasswordClient is a client for the UserPassword schema.
type UserPasswordClient struct {
	config
}

// NewUserPasswordClient returns a client for the UserPassword from the given config.
func NewUserPasswordClient(c config) *UserPasswordClient {
	return &UserPasswordClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `userpassword.Hooks(f(g(h())))`.
func (c *UserPasswordClient) Use(hooks ...Hook) {
	c.hooks.UserPassword = append(c.hooks.UserPassword, hooks...)
}

// Create returns a builder for creating a UserPassword entity.
func (c *UserPasswordClient) Create() *UserPasswordCreate {
	mutation := newUserPasswordMutation(c.config, OpCreate)
	return &UserPasswordCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserPassword entities.
func (c *UserPasswordClient) CreateBulk(builders ...*UserPasswordCreate) *UserPasswordCreateBulk {
	return &UserPasswordCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserPassword.
func (c *UserPasswordClient) Update() *UserPasswordUpdate {
	mutation := newUserPasswordMutation(c.config, OpUpdate)
	return &UserPasswordUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserPasswordClient) UpdateOne(up *UserPassword) *UserPasswordUpdateOne {
	mutation := newUserPasswordMutation(c.config, OpUpdateOne, withUserPassword(up))
	return &UserPasswordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserPasswordClient) UpdateOneID(id uint32) *UserPasswordUpdateOne {
	mutation := newUserPasswordMutation(c.config, OpUpdateOne, withUserPasswordID(id))
	return &UserPasswordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserPassword.
func (c *UserPasswordClient) Delete() *UserPasswordDelete {
	mutation := newUserPasswordMutation(c.config, OpDelete)
	return &UserPasswordDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserPasswordClient) DeleteOne(up *UserPassword) *UserPasswordDeleteOne {
	return c.DeleteOneID(up.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserPasswordClient) DeleteOneID(id uint32) *UserPasswordDeleteOne {
	builder := c.Delete().Where(userpassword.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserPasswordDeleteOne{builder}
}

// Query returns a query builder for UserPassword.
func (c *UserPasswordClient) Query() *UserPasswordQuery {
	return &UserPasswordQuery{
		config: c.config,
	}
}

// Get returns a UserPassword entity by its id.
func (c *UserPasswordClient) Get(ctx context.Context, id uint32) (*UserPassword, error) {
	return c.Query().Where(userpassword.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserPasswordClient) GetX(ctx context.Context, id uint32) *UserPassword {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a UserPassword.
func (c *UserPasswordClient) QueryUser(up *UserPassword) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := up.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(userpassword.Table, userpassword.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, userpassword.UserTable, userpassword.UserColumn),
		)
		fromV = sqlgraph.Neighbors(up.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserPasswordClient) Hooks() []Hook {
	return c.hooks.UserPassword
}
