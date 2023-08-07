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

// CarsQuery is the builder for querying Cars entities.
type CarsQuery struct {
	config
	ctx                     *QueryContext
	order                   []cars.OrderOption
	inters                  []Interceptor
	predicates              []predicate.Cars
	withDealershipCars      *DealershipQuery
	loadTotal               []func(context.Context, []*Cars) error
	modifiers               []func(*sql.Selector)
	withNamedDealershipCars map[string]*DealershipQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CarsQuery builder.
func (cq *CarsQuery) Where(ps ...predicate.Cars) *CarsQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CarsQuery) Limit(limit int) *CarsQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CarsQuery) Offset(offset int) *CarsQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CarsQuery) Unique(unique bool) *CarsQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CarsQuery) Order(o ...cars.OrderOption) *CarsQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryDealershipCars chains the current query on the "dealership_cars" edge.
func (cq *CarsQuery) QueryDealershipCars() *DealershipQuery {
	query := (&DealershipClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cars.Table, cars.FieldID, selector),
			sqlgraph.To(dealership.Table, dealership.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, cars.DealershipCarsTable, cars.DealershipCarsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Cars entity from the query.
// Returns a *NotFoundError when no Cars was found.
func (cq *CarsQuery) First(ctx context.Context) (*Cars, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cars.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CarsQuery) FirstX(ctx context.Context) *Cars {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Cars ID from the query.
// Returns a *NotFoundError when no Cars ID was found.
func (cq *CarsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cars.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CarsQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Cars entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Cars entity is found.
// Returns a *NotFoundError when no Cars entities are found.
func (cq *CarsQuery) Only(ctx context.Context) (*Cars, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cars.Label}
	default:
		return nil, &NotSingularError{cars.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CarsQuery) OnlyX(ctx context.Context) *Cars {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Cars ID in the query.
// Returns a *NotSingularError when more than one Cars ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CarsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cars.Label}
	default:
		err = &NotSingularError{cars.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CarsQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CarsSlice.
func (cq *CarsQuery) All(ctx context.Context) ([]*Cars, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Cars, *CarsQuery]()
	return withInterceptors[[]*Cars](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CarsQuery) AllX(ctx context.Context) []*Cars {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Cars IDs.
func (cq *CarsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(cars.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CarsQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CarsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CarsQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CarsQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CarsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CarsQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CarsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CarsQuery) Clone() *CarsQuery {
	if cq == nil {
		return nil
	}
	return &CarsQuery{
		config:             cq.config,
		ctx:                cq.ctx.Clone(),
		order:              append([]cars.OrderOption{}, cq.order...),
		inters:             append([]Interceptor{}, cq.inters...),
		predicates:         append([]predicate.Cars{}, cq.predicates...),
		withDealershipCars: cq.withDealershipCars.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithDealershipCars tells the query-builder to eager-load the nodes that are connected to
// the "dealership_cars" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CarsQuery) WithDealershipCars(opts ...func(*DealershipQuery)) *CarsQuery {
	query := (&DealershipClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withDealershipCars = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (cq *CarsQuery) GroupBy(field string, fields ...string) *CarsGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CarsGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = cars.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (cq *CarsQuery) Select(fields ...string) *CarsSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CarsSelect{CarsQuery: cq}
	sbuild.label = cars.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CarsSelect configured with the given aggregations.
func (cq *CarsQuery) Aggregate(fns ...AggregateFunc) *CarsSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CarsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !cars.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CarsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Cars, error) {
	var (
		nodes       = []*Cars{}
		_spec       = cq.querySpec()
		loadedTypes = [1]bool{
			cq.withDealershipCars != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Cars).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Cars{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withDealershipCars; query != nil {
		if err := cq.loadDealershipCars(ctx, query, nodes,
			func(n *Cars) { n.Edges.DealershipCars = []*Dealership{} },
			func(n *Cars, e *Dealership) { n.Edges.DealershipCars = append(n.Edges.DealershipCars, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedDealershipCars {
		if err := cq.loadDealershipCars(ctx, query, nodes,
			func(n *Cars) { n.appendNamedDealershipCars(name) },
			func(n *Cars, e *Dealership) { n.appendNamedDealershipCars(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range cq.loadTotal {
		if err := cq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CarsQuery) loadDealershipCars(ctx context.Context, query *DealershipQuery, nodes []*Cars, init func(*Cars), assign func(*Cars, *Dealership)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Cars)
	nids := make(map[int]map[*Cars]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(cars.DealershipCarsTable)
		s.Join(joinT).On(s.C(dealership.FieldID), joinT.C(cars.DealershipCarsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(cars.DealershipCarsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(cars.DealershipCarsPrimaryKey[0]))
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
					nids[inValue] = map[*Cars]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Dealership](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "dealership_cars" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (cq *CarsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CarsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(cars.Table, cars.Columns, sqlgraph.NewFieldSpec(cars.FieldID, field.TypeInt))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cars.FieldID)
		for i := range fields {
			if fields[i] != cars.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CarsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(cars.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = cars.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range cq.modifiers {
		m(selector)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cq *CarsQuery) Modify(modifiers ...func(s *sql.Selector)) *CarsSelect {
	cq.modifiers = append(cq.modifiers, modifiers...)
	return cq.Select()
}

// WithNamedDealershipCars tells the query-builder to eager-load the nodes that are connected to the "dealership_cars"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *CarsQuery) WithNamedDealershipCars(name string, opts ...func(*DealershipQuery)) *CarsQuery {
	query := (&DealershipClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedDealershipCars == nil {
		cq.withNamedDealershipCars = make(map[string]*DealershipQuery)
	}
	cq.withNamedDealershipCars[name] = query
	return cq
}

// CarsGroupBy is the group-by builder for Cars entities.
type CarsGroupBy struct {
	selector
	build *CarsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CarsGroupBy) Aggregate(fns ...AggregateFunc) *CarsGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CarsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CarsQuery, *CarsGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CarsGroupBy) sqlScan(ctx context.Context, root *CarsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CarsSelect is the builder for selecting fields of Cars entities.
type CarsSelect struct {
	*CarsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CarsSelect) Aggregate(fns ...AggregateFunc) *CarsSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CarsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CarsQuery, *CarsSelect](ctx, cs.CarsQuery, cs, cs.inters, v)
}

func (cs *CarsSelect) sqlScan(ctx context.Context, root *CarsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cs *CarsSelect) Modify(modifiers ...func(s *sql.Selector)) *CarsSelect {
	cs.modifiers = append(cs.modifiers, modifiers...)
	return cs
}
