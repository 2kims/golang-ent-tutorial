// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/yowmark/go-ent-pokemon/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/yowmark/go-ent-pokemon/ent/battle"
	"github.com/yowmark/go-ent-pokemon/ent/pokemon"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Battle is the client for interacting with the Battle builders.
	Battle *BattleClient
	// Pokemon is the client for interacting with the Pokemon builders.
	Pokemon *PokemonClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Battle = NewBattleClient(c.config)
	c.Pokemon = NewPokemonClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Battle:  NewBattleClient(cfg),
		Pokemon: NewPokemonClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		ctx:     ctx,
		config:  cfg,
		Battle:  NewBattleClient(cfg),
		Pokemon: NewPokemonClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Battle.
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
	c.Battle.Use(hooks...)
	c.Pokemon.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Battle.Intercept(interceptors...)
	c.Pokemon.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *BattleMutation:
		return c.Battle.mutate(ctx, m)
	case *PokemonMutation:
		return c.Pokemon.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// BattleClient is a client for the Battle schema.
type BattleClient struct {
	config
}

// NewBattleClient returns a client for the Battle from the given config.
func NewBattleClient(c config) *BattleClient {
	return &BattleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `battle.Hooks(f(g(h())))`.
func (c *BattleClient) Use(hooks ...Hook) {
	c.hooks.Battle = append(c.hooks.Battle, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `battle.Intercept(f(g(h())))`.
func (c *BattleClient) Intercept(interceptors ...Interceptor) {
	c.inters.Battle = append(c.inters.Battle, interceptors...)
}

// Create returns a builder for creating a Battle entity.
func (c *BattleClient) Create() *BattleCreate {
	mutation := newBattleMutation(c.config, OpCreate)
	return &BattleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Battle entities.
func (c *BattleClient) CreateBulk(builders ...*BattleCreate) *BattleCreateBulk {
	return &BattleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Battle.
func (c *BattleClient) Update() *BattleUpdate {
	mutation := newBattleMutation(c.config, OpUpdate)
	return &BattleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BattleClient) UpdateOne(b *Battle) *BattleUpdateOne {
	mutation := newBattleMutation(c.config, OpUpdateOne, withBattle(b))
	return &BattleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BattleClient) UpdateOneID(id int) *BattleUpdateOne {
	mutation := newBattleMutation(c.config, OpUpdateOne, withBattleID(id))
	return &BattleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Battle.
func (c *BattleClient) Delete() *BattleDelete {
	mutation := newBattleMutation(c.config, OpDelete)
	return &BattleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BattleClient) DeleteOne(b *Battle) *BattleDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BattleClient) DeleteOneID(id int) *BattleDeleteOne {
	builder := c.Delete().Where(battle.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BattleDeleteOne{builder}
}

// Query returns a query builder for Battle.
func (c *BattleClient) Query() *BattleQuery {
	return &BattleQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeBattle},
		inters: c.Interceptors(),
	}
}

// Get returns a Battle entity by its id.
func (c *BattleClient) Get(ctx context.Context, id int) (*Battle, error) {
	return c.Query().Where(battle.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BattleClient) GetX(ctx context.Context, id int) *Battle {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryContender queries the contender edge of a Battle.
func (c *BattleClient) QueryContender(b *Battle) *PokemonQuery {
	query := (&PokemonClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(battle.Table, battle.FieldID, id),
			sqlgraph.To(pokemon.Table, pokemon.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, battle.ContenderTable, battle.ContenderColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOponent queries the oponent edge of a Battle.
func (c *BattleClient) QueryOponent(b *Battle) *PokemonQuery {
	query := (&PokemonClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(battle.Table, battle.FieldID, id),
			sqlgraph.To(pokemon.Table, pokemon.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, battle.OponentTable, battle.OponentColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BattleClient) Hooks() []Hook {
	return c.hooks.Battle
}

// Interceptors returns the client interceptors.
func (c *BattleClient) Interceptors() []Interceptor {
	return c.inters.Battle
}

func (c *BattleClient) mutate(ctx context.Context, m *BattleMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&BattleCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&BattleUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&BattleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&BattleDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Battle mutation op: %q", m.Op())
	}
}

// PokemonClient is a client for the Pokemon schema.
type PokemonClient struct {
	config
}

// NewPokemonClient returns a client for the Pokemon from the given config.
func NewPokemonClient(c config) *PokemonClient {
	return &PokemonClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `pokemon.Hooks(f(g(h())))`.
func (c *PokemonClient) Use(hooks ...Hook) {
	c.hooks.Pokemon = append(c.hooks.Pokemon, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `pokemon.Intercept(f(g(h())))`.
func (c *PokemonClient) Intercept(interceptors ...Interceptor) {
	c.inters.Pokemon = append(c.inters.Pokemon, interceptors...)
}

// Create returns a builder for creating a Pokemon entity.
func (c *PokemonClient) Create() *PokemonCreate {
	mutation := newPokemonMutation(c.config, OpCreate)
	return &PokemonCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Pokemon entities.
func (c *PokemonClient) CreateBulk(builders ...*PokemonCreate) *PokemonCreateBulk {
	return &PokemonCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Pokemon.
func (c *PokemonClient) Update() *PokemonUpdate {
	mutation := newPokemonMutation(c.config, OpUpdate)
	return &PokemonUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PokemonClient) UpdateOne(po *Pokemon) *PokemonUpdateOne {
	mutation := newPokemonMutation(c.config, OpUpdateOne, withPokemon(po))
	return &PokemonUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PokemonClient) UpdateOneID(id int) *PokemonUpdateOne {
	mutation := newPokemonMutation(c.config, OpUpdateOne, withPokemonID(id))
	return &PokemonUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Pokemon.
func (c *PokemonClient) Delete() *PokemonDelete {
	mutation := newPokemonMutation(c.config, OpDelete)
	return &PokemonDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PokemonClient) DeleteOne(po *Pokemon) *PokemonDeleteOne {
	return c.DeleteOneID(po.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PokemonClient) DeleteOneID(id int) *PokemonDeleteOne {
	builder := c.Delete().Where(pokemon.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PokemonDeleteOne{builder}
}

// Query returns a query builder for Pokemon.
func (c *PokemonClient) Query() *PokemonQuery {
	return &PokemonQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePokemon},
		inters: c.Interceptors(),
	}
}

// Get returns a Pokemon entity by its id.
func (c *PokemonClient) Get(ctx context.Context, id int) (*Pokemon, error) {
	return c.Query().Where(pokemon.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PokemonClient) GetX(ctx context.Context, id int) *Pokemon {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryFights queries the fights edge of a Pokemon.
func (c *PokemonClient) QueryFights(po *Pokemon) *BattleQuery {
	query := (&BattleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(pokemon.Table, pokemon.FieldID, id),
			sqlgraph.To(battle.Table, battle.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, pokemon.FightsTable, pokemon.FightsColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOpponents queries the opponents edge of a Pokemon.
func (c *PokemonClient) QueryOpponents(po *Pokemon) *BattleQuery {
	query := (&BattleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(pokemon.Table, pokemon.FieldID, id),
			sqlgraph.To(battle.Table, battle.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, pokemon.OpponentsTable, pokemon.OpponentsColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PokemonClient) Hooks() []Hook {
	return c.hooks.Pokemon
}

// Interceptors returns the client interceptors.
func (c *PokemonClient) Interceptors() []Interceptor {
	return c.inters.Pokemon
}

func (c *PokemonClient) mutate(ctx context.Context, m *PokemonMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PokemonCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PokemonUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PokemonUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PokemonDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Pokemon mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Battle, Pokemon []ent.Hook
	}
	inters struct {
		Battle, Pokemon []ent.Interceptor
	}
)
