// Copyright 2022-present Vlabs Development Kft
//
// All rights reserved under a proprietary license.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"dealership/ent/cars"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Cars is the model entity for the Cars schema.
type Cars struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CarsQuery when eager-loading is set.
	Edges        CarsEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CarsEdges holds the relations/edges for other nodes in the graph.
type CarsEdges struct {
	// DealershipCars holds the value of the dealership_cars edge.
	DealershipCars []*Dealership `json:"dealership_cars,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool

	namedDealershipCars map[string][]*Dealership
}

// DealershipCarsOrErr returns the DealershipCars value or an error if the edge
// was not loaded in eager-loading.
func (e CarsEdges) DealershipCarsOrErr() ([]*Dealership, error) {
	if e.loadedTypes[0] {
		return e.DealershipCars, nil
	}
	return nil, &NotLoadedError{edge: "dealership_cars"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Cars) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cars.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Cars fields.
func (c *Cars) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cars.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Cars.
// This includes values selected through modifiers, order, etc.
func (c *Cars) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryDealershipCars queries the "dealership_cars" edge of the Cars entity.
func (c *Cars) QueryDealershipCars() *DealershipQuery {
	return NewCarsClient(c.config).QueryDealershipCars(c)
}

// Update returns a builder for updating this Cars.
// Note that you need to call Cars.Unwrap() before calling this method if this Cars
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Cars) Update() *CarsUpdateOne {
	return NewCarsClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Cars entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Cars) Unwrap() *Cars {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Cars is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Cars) String() string {
	var builder strings.Builder
	builder.WriteString("Cars(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteByte(')')
	return builder.String()
}

// NamedDealershipCars returns the DealershipCars named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Cars) NamedDealershipCars(name string) ([]*Dealership, error) {
	if c.Edges.namedDealershipCars == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedDealershipCars[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Cars) appendNamedDealershipCars(name string, edges ...*Dealership) {
	if c.Edges.namedDealershipCars == nil {
		c.Edges.namedDealershipCars = make(map[string][]*Dealership)
	}
	if len(edges) == 0 {
		c.Edges.namedDealershipCars[name] = []*Dealership{}
	} else {
		c.Edges.namedDealershipCars[name] = append(c.Edges.namedDealershipCars[name], edges...)
	}
}

// CarsSlice is a parsable slice of Cars.
type CarsSlice []*Cars
