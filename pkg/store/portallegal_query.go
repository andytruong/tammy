// Code generated by ent, DO NOT EDIT.

package store

import (
	"context"
	"fmt"
	"math"
	"tammy/pkg/store/portal"
	"tammy/pkg/store/portallegal"
	"tammy/pkg/store/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PortalLegalQuery is the builder for querying PortalLegal entities.
type PortalLegalQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PortalLegal
	// eager-loading edges.
	withPortal *PortalQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PortalLegalQuery builder.
func (plq *PortalLegalQuery) Where(ps ...predicate.PortalLegal) *PortalLegalQuery {
	plq.predicates = append(plq.predicates, ps...)
	return plq
}

// Limit adds a limit step to the query.
func (plq *PortalLegalQuery) Limit(limit int) *PortalLegalQuery {
	plq.limit = &limit
	return plq
}

// Offset adds an offset step to the query.
func (plq *PortalLegalQuery) Offset(offset int) *PortalLegalQuery {
	plq.offset = &offset
	return plq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (plq *PortalLegalQuery) Unique(unique bool) *PortalLegalQuery {
	plq.unique = &unique
	return plq
}

// Order adds an order step to the query.
func (plq *PortalLegalQuery) Order(o ...OrderFunc) *PortalLegalQuery {
	plq.order = append(plq.order, o...)
	return plq
}

// QueryPortal chains the current query on the "portal" edge.
func (plq *PortalLegalQuery) QueryPortal() *PortalQuery {
	query := &PortalQuery{config: plq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := plq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := plq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(portallegal.Table, portallegal.FieldID, selector),
			sqlgraph.To(portal.Table, portal.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, portallegal.PortalTable, portallegal.PortalColumn),
		)
		fromU = sqlgraph.SetNeighbors(plq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PortalLegal entity from the query.
// Returns a *NotFoundError when no PortalLegal was found.
func (plq *PortalLegalQuery) First(ctx context.Context) (*PortalLegal, error) {
	nodes, err := plq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{portallegal.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (plq *PortalLegalQuery) FirstX(ctx context.Context) *PortalLegal {
	node, err := plq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PortalLegal ID from the query.
// Returns a *NotFoundError when no PortalLegal ID was found.
func (plq *PortalLegalQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = plq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{portallegal.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (plq *PortalLegalQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := plq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PortalLegal entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PortalLegal entity is found.
// Returns a *NotFoundError when no PortalLegal entities are found.
func (plq *PortalLegalQuery) Only(ctx context.Context) (*PortalLegal, error) {
	nodes, err := plq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{portallegal.Label}
	default:
		return nil, &NotSingularError{portallegal.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (plq *PortalLegalQuery) OnlyX(ctx context.Context) *PortalLegal {
	node, err := plq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PortalLegal ID in the query.
// Returns a *NotSingularError when more than one PortalLegal ID is found.
// Returns a *NotFoundError when no entities are found.
func (plq *PortalLegalQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = plq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{portallegal.Label}
	default:
		err = &NotSingularError{portallegal.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (plq *PortalLegalQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := plq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PortalLegals.
func (plq *PortalLegalQuery) All(ctx context.Context) ([]*PortalLegal, error) {
	if err := plq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return plq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (plq *PortalLegalQuery) AllX(ctx context.Context) []*PortalLegal {
	nodes, err := plq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PortalLegal IDs.
func (plq *PortalLegalQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := plq.Select(portallegal.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (plq *PortalLegalQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := plq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (plq *PortalLegalQuery) Count(ctx context.Context) (int, error) {
	if err := plq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return plq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (plq *PortalLegalQuery) CountX(ctx context.Context) int {
	count, err := plq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (plq *PortalLegalQuery) Exist(ctx context.Context) (bool, error) {
	if err := plq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return plq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (plq *PortalLegalQuery) ExistX(ctx context.Context) bool {
	exist, err := plq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PortalLegalQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (plq *PortalLegalQuery) Clone() *PortalLegalQuery {
	if plq == nil {
		return nil
	}
	return &PortalLegalQuery{
		config:     plq.config,
		limit:      plq.limit,
		offset:     plq.offset,
		order:      append([]OrderFunc{}, plq.order...),
		predicates: append([]predicate.PortalLegal{}, plq.predicates...),
		withPortal: plq.withPortal.Clone(),
		// clone intermediate query.
		sql:    plq.sql.Clone(),
		path:   plq.path,
		unique: plq.unique,
	}
}

// WithPortal tells the query-builder to eager-load the nodes that are connected to
// the "portal" edge. The optional arguments are used to configure the query builder of the edge.
func (plq *PortalLegalQuery) WithPortal(opts ...func(*PortalQuery)) *PortalLegalQuery {
	query := &PortalQuery{config: plq.config}
	for _, opt := range opts {
		opt(query)
	}
	plq.withPortal = query
	return plq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PortalLegal.Query().
//		GroupBy(portallegal.FieldCreatedAt).
//		Aggregate(store.Count()).
//		Scan(ctx, &v)
//
func (plq *PortalLegalQuery) GroupBy(field string, fields ...string) *PortalLegalGroupBy {
	grbuild := &PortalLegalGroupBy{config: plq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := plq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return plq.sqlQuery(ctx), nil
	}
	grbuild.label = portallegal.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty"`
//	}
//
//	client.PortalLegal.Query().
//		Select(portallegal.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (plq *PortalLegalQuery) Select(fields ...string) *PortalLegalSelect {
	plq.fields = append(plq.fields, fields...)
	selbuild := &PortalLegalSelect{PortalLegalQuery: plq}
	selbuild.label = portallegal.Label
	selbuild.flds, selbuild.scan = &plq.fields, selbuild.Scan
	return selbuild
}

func (plq *PortalLegalQuery) prepareQuery(ctx context.Context) error {
	for _, f := range plq.fields {
		if !portallegal.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("store: invalid field %q for query", f)}
		}
	}
	if plq.path != nil {
		prev, err := plq.path(ctx)
		if err != nil {
			return err
		}
		plq.sql = prev
	}
	return nil
}

func (plq *PortalLegalQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PortalLegal, error) {
	var (
		nodes       = []*PortalLegal{}
		withFKs     = plq.withFKs
		_spec       = plq.querySpec()
		loadedTypes = [1]bool{
			plq.withPortal != nil,
		}
	)
	if plq.withPortal != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, portallegal.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*PortalLegal).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &PortalLegal{config: plq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, plq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := plq.withPortal; query != nil {
		ids := make([]uint32, 0, len(nodes))
		nodeids := make(map[uint32][]*PortalLegal)
		for i := range nodes {
			if nodes[i].portal_legal == nil {
				continue
			}
			fk := *nodes[i].portal_legal
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(portal.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "portal_legal" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Portal = n
			}
		}
	}

	return nodes, nil
}

func (plq *PortalLegalQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := plq.querySpec()
	_spec.Node.Columns = plq.fields
	if len(plq.fields) > 0 {
		_spec.Unique = plq.unique != nil && *plq.unique
	}
	return sqlgraph.CountNodes(ctx, plq.driver, _spec)
}

func (plq *PortalLegalQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := plq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("store: check existence: %w", err)
	}
	return n > 0, nil
}

func (plq *PortalLegalQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   portallegal.Table,
			Columns: portallegal.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: portallegal.FieldID,
			},
		},
		From:   plq.sql,
		Unique: true,
	}
	if unique := plq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := plq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, portallegal.FieldID)
		for i := range fields {
			if fields[i] != portallegal.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := plq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := plq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := plq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := plq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (plq *PortalLegalQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(plq.driver.Dialect())
	t1 := builder.Table(portallegal.Table)
	columns := plq.fields
	if len(columns) == 0 {
		columns = portallegal.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if plq.sql != nil {
		selector = plq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if plq.unique != nil && *plq.unique {
		selector.Distinct()
	}
	for _, p := range plq.predicates {
		p(selector)
	}
	for _, p := range plq.order {
		p(selector)
	}
	if offset := plq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := plq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PortalLegalGroupBy is the group-by builder for PortalLegal entities.
type PortalLegalGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (plgb *PortalLegalGroupBy) Aggregate(fns ...AggregateFunc) *PortalLegalGroupBy {
	plgb.fns = append(plgb.fns, fns...)
	return plgb
}

// Scan applies the group-by query and scans the result into the given value.
func (plgb *PortalLegalGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := plgb.path(ctx)
	if err != nil {
		return err
	}
	plgb.sql = query
	return plgb.sqlScan(ctx, v)
}

func (plgb *PortalLegalGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range plgb.fields {
		if !portallegal.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := plgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := plgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (plgb *PortalLegalGroupBy) sqlQuery() *sql.Selector {
	selector := plgb.sql.Select()
	aggregation := make([]string, 0, len(plgb.fns))
	for _, fn := range plgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(plgb.fields)+len(plgb.fns))
		for _, f := range plgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(plgb.fields...)...)
}

// PortalLegalSelect is the builder for selecting fields of PortalLegal entities.
type PortalLegalSelect struct {
	*PortalLegalQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pls *PortalLegalSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pls.prepareQuery(ctx); err != nil {
		return err
	}
	pls.sql = pls.PortalLegalQuery.sqlQuery(ctx)
	return pls.sqlScan(ctx, v)
}

func (pls *PortalLegalSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pls.sql.Query()
	if err := pls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}