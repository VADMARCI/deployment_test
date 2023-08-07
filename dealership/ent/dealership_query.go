// Copyright 2022-present Vlabs Development Kft
//
// All rights reserved under a proprietary license.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"dealership/ent/cars"
	"dealership/ent/dealership"
	"dealership/ent/predicate"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DealershipQuery is the builder for querying Dealership entities.
type DealershipQuery struct {
	config
	ctx           *QueryContext
	order         []dealership.OrderOption
	inters        []Interceptor
	predicates    []predicate.Dealership
	withCars      *CarsQuery
	loadTotal     []func(context.Context, []*Dealership) error
	modifiers     []func(*sql.Selector)
	withNamedCars map[string]*CarsQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DealershipQuery builder.
func (dq *DealershipQuery) Where(ps ...predicate.Dealership) *DealershipQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DealershipQuery) Limit(limit int) *DealershipQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DealershipQuery) Offset(offset int) *DealershipQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DealershipQuery) Unique(unique bool) *DealershipQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DealershipQuery) Order(o ...dealership.OrderOption) *DealershipQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryCars chains the current query on the "cars" edge.
func (dq *DealershipQuery) QueryCars() *CarsQuery {
	query := (&CarsClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dealership.Table, dealership.FieldID, selector),
			sqlgraph.To(cars.Table, cars.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, dealership.CarsTable, dealership.CarsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Dealership entity from the query.
// Returns a *NotFoundError when no Dealership was found.
func (dq *DealershipQuery) First(ctx context.Context) (*Dealership, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dealership.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DealershipQuery) FirstX(ctx context.Context) *Dealership {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Dealership ID from the query.
// Returns a *NotFoundError when no Dealership ID was found.
func (dq *DealershipQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dealership.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DealershipQuery) FirstIDX(ctx context.Context) int {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Dealership entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Dealership entity is found.
// Returns a *NotFoundError when no Dealership entities are found.
func (dq *DealershipQuery) Only(ctx context.Context) (*Dealership, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dealership.Label}
	default:
		return nil, &NotSingularError{dealership.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DealershipQuery) OnlyX(ctx context.Context) *Dealership {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Dealership ID in the query.
// Returns a *NotSingularError when more than one Dealership ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DealershipQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dealership.Label}
	default:
		err = &NotSingularError{dealership.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DealershipQuery) OnlyIDX(ctx context.Context) int {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Dealerships.
func (dq *DealershipQuery) All(ctx context.Context) ([]*Dealership, error) {
	ctx = setContextOp(ctx, dq.ctx, "All")
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Dealership, *DealershipQuery]()
	return withInterceptors[[]*Dealership](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DealershipQuery) AllX(ctx context.Context) []*Dealership {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Dealership IDs.
func (dq *DealershipQuery) IDs(ctx context.Context) (ids []int, err error) {
	if dq.ctx.Unique == nil && dq.path != nil {
		dq.Unique(true)
	}
	ctx = setContextOp(ctx, dq.ctx, "IDs")
	if err = dq.Select(dealership.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DealershipQuery) IDsX(ctx context.Context) []int {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DealershipQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, "Count")
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DealershipQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DealershipQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DealershipQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, "Exist")
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DealershipQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DealershipQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DealershipQuery) Clone() *DealershipQuery {
	if dq == nil {
		return nil
	}
	return &DealershipQuery{
		config:     dq.config,
		ctx:        dq.ctx.Clone(),
		order:      append([]dealership.OrderOption{}, dq.order...),
		inters:     append([]Interceptor{}, dq.inters...),
		predicates: append([]predicate.Dealership{}, dq.predicates...),
		withCars:   dq.withCars.Clone(),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

// WithCars tells the query-builder to eager-load the nodes that are connected to
// the "cars" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DealershipQuery) WithCars(opts ...func(*CarsQuery)) *DealershipQuery {
	query := (&CarsClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withCars = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		City string `json:"city,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Dealership.Query().
//		GroupBy(dealership.FieldCity).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DealershipQuery) GroupBy(field string, fields ...string) *DealershipGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DealershipGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = dealership.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		City string `json:"city,omitempty"`
//	}
//
//	client.Dealership.Query().
//		Select(dealership.FieldCity).
//		Scan(ctx, &v)
func (dq *DealershipQuery) Select(fields ...string) *DealershipSelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DealershipSelect{DealershipQuery: dq}
	sbuild.label = dealership.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DealershipSelect configured with the given aggregations.
func (dq *DealershipQuery) Aggregate(fns ...AggregateFunc) *DealershipSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DealershipQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !dealership.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DealershipQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Dealership, error) {
	var (
		nodes       = []*Dealership{}
		_spec       = dq.querySpec()
		loadedTypes = [1]bool{
			dq.withCars != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Dealership).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Dealership{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(dq.modifiers) > 0 {
		_spec.Modifiers = dq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withCars; query != nil {
		if err := dq.loadCars(ctx, query, nodes,
			func(n *Dealership) { n.Edges.Cars = []*Cars{} },
			func(n *Dealership, e *Cars) { n.Edges.Cars = append(n.Edges.Cars, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range dq.withNamedCars {
		if err := dq.loadCars(ctx, query, nodes,
			func(n *Dealership) { n.appendNamedCars(name) },
			func(n *Dealership, e *Cars) { n.appendNamedCars(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range dq.loadTotal {
		if err := dq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DealershipQuery) loadCars(ctx context.Context, query *CarsQuery, nodes []*Dealership, init func(*Dealership), assign func(*Dealership, *Cars)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Dealership)
	nids := make(map[int]map[*Dealership]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(dealership.CarsTable)
		s.Join(joinT).On(s.C(cars.FieldID), joinT.C(dealership.CarsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(dealership.CarsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(dealership.CarsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Dealership]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Cars](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "cars" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (dq *DealershipQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	if len(dq.modifiers) > 0 {
		_spec.Modifiers = dq.modifiers
	}
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DealershipQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(dealership.Table, dealership.Columns, sqlgraph.NewFieldSpec(dealership.FieldID, field.TypeInt))
	_spec.From = dq.sql
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dq.path != nil {
		_spec.Unique = true
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dealership.FieldID)
		for i := range fields {
			if fields[i] != dealership.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DealershipQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(dealership.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = dealership.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range dq.modifiers {
		m(selector)
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (dq *DealershipQuery) Modify(modifiers ...func(s *sql.Selector)) *DealershipSelect {
	dq.modifiers = append(dq.modifiers, modifiers...)
	return dq.Select()
}

// WithNamedCars tells the query-builder to eager-load the nodes that are connected to the "cars"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (dq *DealershipQuery) WithNamedCars(name string, opts ...func(*CarsQuery)) *DealershipQuery {
	query := (&CarsClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if dq.withNamedCars == nil {
		dq.withNamedCars = make(map[string]*CarsQuery)
	}
	dq.withNamedCars[name] = query
	return dq
}

// DealershipGroupBy is the group-by builder for Dealership entities.
type DealershipGroupBy struct {
	selector
	build *DealershipQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DealershipGroupBy) Aggregate(fns ...AggregateFunc) *DealershipGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DealershipGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, "GroupBy")
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DealershipQuery, *DealershipGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DealershipGroupBy) sqlScan(ctx context.Context, root *DealershipQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DealershipSelect is the builder for selecting fields of Dealership entities.
type DealershipSelect struct {
	*DealershipQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DealershipSelect) Aggregate(fns ...AggregateFunc) *DealershipSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DealershipSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, "Select")
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DealershipQuery, *DealershipSelect](ctx, ds.DealershipQuery, ds, ds.inters, v)
}

func (ds *DealershipSelect) sqlScan(ctx context.Context, root *DealershipQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ds *DealershipSelect) Modify(modifiers ...func(s *sql.Selector)) *DealershipSelect {
	ds.modifiers = append(ds.modifiers, modifiers...)
	return ds
}
