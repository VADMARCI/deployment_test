// Copyright 2022-present Vlabs Development Kft
//
// All rights reserved under a proprietary license.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"dealership/ent/cars"
	"dealership/ent/dealership"
	"dealership/ent/predicate"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DealershipUpdate is the builder for updating Dealership entities.
type DealershipUpdate struct {
	config
	hooks     []Hook
	mutation  *DealershipMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the DealershipUpdate builder.
func (du *DealershipUpdate) Where(ps ...predicate.Dealership) *DealershipUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetCity sets the "city" field.
func (du *DealershipUpdate) SetCity(s string) *DealershipUpdate {
	du.mutation.SetCity(s)
	return du
}

// SetName sets the "name" field.
func (du *DealershipUpdate) SetName(s string) *DealershipUpdate {
	du.mutation.SetName(s)
	return du
}

// AddCarIDs adds the "cars" edge to the Cars entity by IDs.
func (du *DealershipUpdate) AddCarIDs(ids ...int) *DealershipUpdate {
	du.mutation.AddCarIDs(ids...)
	return du
}

// AddCars adds the "cars" edges to the Cars entity.
func (du *DealershipUpdate) AddCars(c ...*Cars) *DealershipUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return du.AddCarIDs(ids...)
}

// Mutation returns the DealershipMutation object of the builder.
func (du *DealershipUpdate) Mutation() *DealershipMutation {
	return du.mutation
}

// ClearCars clears all "cars" edges to the Cars entity.
func (du *DealershipUpdate) ClearCars() *DealershipUpdate {
	du.mutation.ClearCars()
	return du
}

// RemoveCarIDs removes the "cars" edge to Cars entities by IDs.
func (du *DealershipUpdate) RemoveCarIDs(ids ...int) *DealershipUpdate {
	du.mutation.RemoveCarIDs(ids...)
	return du
}

// RemoveCars removes "cars" edges to Cars entities.
func (du *DealershipUpdate) RemoveCars(c ...*Cars) *DealershipUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return du.RemoveCarIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DealershipUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DealershipUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DealershipUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DealershipUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (du *DealershipUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DealershipUpdate {
	du.modifiers = append(du.modifiers, modifiers...)
	return du
}

func (du *DealershipUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(dealership.Table, dealership.Columns, sqlgraph.NewFieldSpec(dealership.FieldID, field.TypeInt))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.City(); ok {
		_spec.SetField(dealership.FieldCity, field.TypeString, value)
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.SetField(dealership.FieldName, field.TypeString, value)
	}
	if du.mutation.CarsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dealership.CarsTable,
			Columns: dealership.CarsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cars.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedCarsIDs(); len(nodes) > 0 && !du.mutation.CarsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dealership.CarsTable,
			Columns: dealership.CarsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cars.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.CarsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dealership.CarsTable,
			Columns: dealership.CarsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cars.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(du.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dealership.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DealershipUpdateOne is the builder for updating a single Dealership entity.
type DealershipUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *DealershipMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCity sets the "city" field.
func (duo *DealershipUpdateOne) SetCity(s string) *DealershipUpdateOne {
	duo.mutation.SetCity(s)
	return duo
}

// SetName sets the "name" field.
func (duo *DealershipUpdateOne) SetName(s string) *DealershipUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// AddCarIDs adds the "cars" edge to the Cars entity by IDs.
func (duo *DealershipUpdateOne) AddCarIDs(ids ...int) *DealershipUpdateOne {
	duo.mutation.AddCarIDs(ids...)
	return duo
}

// AddCars adds the "cars" edges to the Cars entity.
func (duo *DealershipUpdateOne) AddCars(c ...*Cars) *DealershipUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return duo.AddCarIDs(ids...)
}

// Mutation returns the DealershipMutation object of the builder.
func (duo *DealershipUpdateOne) Mutation() *DealershipMutation {
	return duo.mutation
}

// ClearCars clears all "cars" edges to the Cars entity.
func (duo *DealershipUpdateOne) ClearCars() *DealershipUpdateOne {
	duo.mutation.ClearCars()
	return duo
}

// RemoveCarIDs removes the "cars" edge to Cars entities by IDs.
func (duo *DealershipUpdateOne) RemoveCarIDs(ids ...int) *DealershipUpdateOne {
	duo.mutation.RemoveCarIDs(ids...)
	return duo
}

// RemoveCars removes "cars" edges to Cars entities.
func (duo *DealershipUpdateOne) RemoveCars(c ...*Cars) *DealershipUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return duo.RemoveCarIDs(ids...)
}

// Where appends a list predicates to the DealershipUpdate builder.
func (duo *DealershipUpdateOne) Where(ps ...predicate.Dealership) *DealershipUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DealershipUpdateOne) Select(field string, fields ...string) *DealershipUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Dealership entity.
func (duo *DealershipUpdateOne) Save(ctx context.Context) (*Dealership, error) {
	return withHooks(ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DealershipUpdateOne) SaveX(ctx context.Context) *Dealership {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DealershipUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DealershipUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (duo *DealershipUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DealershipUpdateOne {
	duo.modifiers = append(duo.modifiers, modifiers...)
	return duo
}

func (duo *DealershipUpdateOne) sqlSave(ctx context.Context) (_node *Dealership, err error) {
	_spec := sqlgraph.NewUpdateSpec(dealership.Table, dealership.Columns, sqlgraph.NewFieldSpec(dealership.FieldID, field.TypeInt))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Dealership.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dealership.FieldID)
		for _, f := range fields {
			if !dealership.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dealership.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.City(); ok {
		_spec.SetField(dealership.FieldCity, field.TypeString, value)
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.SetField(dealership.FieldName, field.TypeString, value)
	}
	if duo.mutation.CarsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dealership.CarsTable,
			Columns: dealership.CarsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cars.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedCarsIDs(); len(nodes) > 0 && !duo.mutation.CarsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dealership.CarsTable,
			Columns: dealership.CarsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cars.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.CarsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dealership.CarsTable,
			Columns: dealership.CarsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cars.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(duo.modifiers...)
	_node = &Dealership{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dealership.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}
