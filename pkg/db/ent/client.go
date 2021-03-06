// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/NpoolPlatform/third-login-gateway/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/third-login-gateway/pkg/db/ent/auth"
	"github.com/NpoolPlatform/third-login-gateway/pkg/db/ent/thirdparty"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Auth is the client for interacting with the Auth builders.
	Auth *AuthClient
	// ThirdParty is the client for interacting with the ThirdParty builders.
	ThirdParty *ThirdPartyClient
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
	c.Auth = NewAuthClient(c.config)
	c.ThirdParty = NewThirdPartyClient(c.config)
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
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Auth:       NewAuthClient(cfg),
		ThirdParty: NewThirdPartyClient(cfg),
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
		ctx:        ctx,
		config:     cfg,
		Auth:       NewAuthClient(cfg),
		ThirdParty: NewThirdPartyClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Auth.
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
	c.Auth.Use(hooks...)
	c.ThirdParty.Use(hooks...)
}

// AuthClient is a client for the Auth schema.
type AuthClient struct {
	config
}

// NewAuthClient returns a client for the Auth from the given config.
func NewAuthClient(c config) *AuthClient {
	return &AuthClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `auth.Hooks(f(g(h())))`.
func (c *AuthClient) Use(hooks ...Hook) {
	c.hooks.Auth = append(c.hooks.Auth, hooks...)
}

// Create returns a create builder for Auth.
func (c *AuthClient) Create() *AuthCreate {
	mutation := newAuthMutation(c.config, OpCreate)
	return &AuthCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Auth entities.
func (c *AuthClient) CreateBulk(builders ...*AuthCreate) *AuthCreateBulk {
	return &AuthCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Auth.
func (c *AuthClient) Update() *AuthUpdate {
	mutation := newAuthMutation(c.config, OpUpdate)
	return &AuthUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AuthClient) UpdateOne(a *Auth) *AuthUpdateOne {
	mutation := newAuthMutation(c.config, OpUpdateOne, withAuth(a))
	return &AuthUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AuthClient) UpdateOneID(id uuid.UUID) *AuthUpdateOne {
	mutation := newAuthMutation(c.config, OpUpdateOne, withAuthID(id))
	return &AuthUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Auth.
func (c *AuthClient) Delete() *AuthDelete {
	mutation := newAuthMutation(c.config, OpDelete)
	return &AuthDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AuthClient) DeleteOne(a *Auth) *AuthDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AuthClient) DeleteOneID(id uuid.UUID) *AuthDeleteOne {
	builder := c.Delete().Where(auth.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AuthDeleteOne{builder}
}

// Query returns a query builder for Auth.
func (c *AuthClient) Query() *AuthQuery {
	return &AuthQuery{
		config: c.config,
	}
}

// Get returns a Auth entity by its id.
func (c *AuthClient) Get(ctx context.Context, id uuid.UUID) (*Auth, error) {
	return c.Query().Where(auth.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AuthClient) GetX(ctx context.Context, id uuid.UUID) *Auth {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AuthClient) Hooks() []Hook {
	hooks := c.hooks.Auth
	return append(hooks[:len(hooks):len(hooks)], auth.Hooks[:]...)
}

// ThirdPartyClient is a client for the ThirdParty schema.
type ThirdPartyClient struct {
	config
}

// NewThirdPartyClient returns a client for the ThirdParty from the given config.
func NewThirdPartyClient(c config) *ThirdPartyClient {
	return &ThirdPartyClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `thirdparty.Hooks(f(g(h())))`.
func (c *ThirdPartyClient) Use(hooks ...Hook) {
	c.hooks.ThirdParty = append(c.hooks.ThirdParty, hooks...)
}

// Create returns a create builder for ThirdParty.
func (c *ThirdPartyClient) Create() *ThirdPartyCreate {
	mutation := newThirdPartyMutation(c.config, OpCreate)
	return &ThirdPartyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ThirdParty entities.
func (c *ThirdPartyClient) CreateBulk(builders ...*ThirdPartyCreate) *ThirdPartyCreateBulk {
	return &ThirdPartyCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ThirdParty.
func (c *ThirdPartyClient) Update() *ThirdPartyUpdate {
	mutation := newThirdPartyMutation(c.config, OpUpdate)
	return &ThirdPartyUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ThirdPartyClient) UpdateOne(tp *ThirdParty) *ThirdPartyUpdateOne {
	mutation := newThirdPartyMutation(c.config, OpUpdateOne, withThirdParty(tp))
	return &ThirdPartyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ThirdPartyClient) UpdateOneID(id uuid.UUID) *ThirdPartyUpdateOne {
	mutation := newThirdPartyMutation(c.config, OpUpdateOne, withThirdPartyID(id))
	return &ThirdPartyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ThirdParty.
func (c *ThirdPartyClient) Delete() *ThirdPartyDelete {
	mutation := newThirdPartyMutation(c.config, OpDelete)
	return &ThirdPartyDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ThirdPartyClient) DeleteOne(tp *ThirdParty) *ThirdPartyDeleteOne {
	return c.DeleteOneID(tp.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ThirdPartyClient) DeleteOneID(id uuid.UUID) *ThirdPartyDeleteOne {
	builder := c.Delete().Where(thirdparty.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ThirdPartyDeleteOne{builder}
}

// Query returns a query builder for ThirdParty.
func (c *ThirdPartyClient) Query() *ThirdPartyQuery {
	return &ThirdPartyQuery{
		config: c.config,
	}
}

// Get returns a ThirdParty entity by its id.
func (c *ThirdPartyClient) Get(ctx context.Context, id uuid.UUID) (*ThirdParty, error) {
	return c.Query().Where(thirdparty.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ThirdPartyClient) GetX(ctx context.Context, id uuid.UUID) *ThirdParty {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ThirdPartyClient) Hooks() []Hook {
	hooks := c.hooks.ThirdParty
	return append(hooks[:len(hooks):len(hooks)], thirdparty.Hooks[:]...)
}
