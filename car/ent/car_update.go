// Copyright 2022-present Vlabs Development Kft
//
// All rights reserved under a proprietary license.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"car/ent/car"
	"car/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CarUpdate is the builder for updating Car entities.
type CarUpdate struct {
	config
	hooks     []Hook
	mutation  *CarMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CarUpdate builder.
func (cu *CarUpdate) Where(ps ...predicate.Car) *CarUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetIsSold sets the "is_sold" field.
func (cu *CarUpdate) SetIsSold(b bool) *CarUpdate {
	cu.mutation.SetIsSold(b)
	return cu
}

// SetNillableIsSold sets the "is_sold" field if the given value is not nil.
func (cu *CarUpdate) SetNillableIsSold(b *bool) *CarUpdate {
	if b != nil {
		cu.SetIsSold(*b)
	}
	return cu
}

// SetName sets the "name" field.
func (cu *CarUpdate) SetName(s string) *CarUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetPrice sets the "price" field.
func (cu *CarUpdate) SetPrice(i int) *CarUpdate {
	cu.mutation.ResetPrice()
	cu.mutation.SetPrice(i)
	return cu
}

// AddPrice adds i to the "price" field.
func (cu *CarUpdate) AddPrice(i int) *CarUpdate {
	cu.mutation.AddPrice(i)
	return cu
}

// Mutation returns the CarMutation object of the builder.
func (cu *CarUpdate) Mutation() *CarMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CarUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CarUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CarUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CarUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *CarUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CarUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *CarUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(car.Table, car.Columns, sqlgraph.NewFieldSpec(car.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.IsSold(); ok {
		_spec.SetField(car.FieldIsSold, field.TypeBool, value)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(car.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Price(); ok {
		_spec.SetField(car.FieldPrice, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedPrice(); ok {
		_spec.AddField(car.FieldPrice, field.TypeInt, value)
	}
	_spec.AddModifiers(cu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{car.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CarUpdateOne is the builder for updating a single Car entity.
type CarUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CarMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetIsSold sets the "is_sold" field.
func (cuo *CarUpdateOne) SetIsSold(b bool) *CarUpdateOne {
	cuo.mutation.SetIsSold(b)
	return cuo
}

// SetNillableIsSold sets the "is_sold" field if the given value is not nil.
func (cuo *CarUpdateOne) SetNillableIsSold(b *bool) *CarUpdateOne {
	if b != nil {
		cuo.SetIsSold(*b)
	}
	return cuo
}

// SetName sets the "name" field.
func (cuo *CarUpdateOne) SetName(s string) *CarUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetPrice sets the "price" field.
func (cuo *CarUpdateOne) SetPrice(i int) *CarUpdateOne {
	cuo.mutation.ResetPrice()
	cuo.mutation.SetPrice(i)
	return cuo
}

// AddPrice adds i to the "price" field.
func (cuo *CarUpdateOne) AddPrice(i int) *CarUpdateOne {
	cuo.mutation.AddPrice(i)
	return cuo
}

// Mutation returns the CarMutation object of the builder.
func (cuo *CarUpdateOne) Mutation() *CarMutation {
	return cuo.mutation
}

// Where appends a list predicates to the CarUpdate builder.
func (cuo *CarUpdateOne) Where(ps ...predicate.Car) *CarUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CarUpdateOne) Select(field string, fields ...string) *CarUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Car entity.
func (cuo *CarUpdateOne) Save(ctx context.Context) (*Car, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CarUpdateOne) SaveX(ctx context.Context) *Car {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CarUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CarUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *CarUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CarUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *CarUpdateOne) sqlSave(ctx context.Context) (_node *Car, err error) {
	_spec := sqlgraph.NewUpdateSpec(car.Table, car.Columns, sqlgraph.NewFieldSpec(car.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Car.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, car.FieldID)
		for _, f := range fields {
			if !car.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != car.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.IsSold(); ok {
		_spec.SetField(car.FieldIsSold, field.TypeBool, value)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(car.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Price(); ok {
		_spec.SetField(car.FieldPrice, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedPrice(); ok {
		_spec.AddField(car.FieldPrice, field.TypeInt, value)
	}
	_spec.AddModifiers(cuo.modifiers...)
	_node = &Car{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{car.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
