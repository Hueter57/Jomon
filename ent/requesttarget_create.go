// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/traPtitech/Jomon/ent/request"
	"github.com/traPtitech/Jomon/ent/requesttarget"
)

// RequestTargetCreate is the builder for creating a RequestTarget entity.
type RequestTargetCreate struct {
	config
	mutation *RequestTargetMutation
	hooks    []Hook
}

// SetTarget sets the "target" field.
func (rtc *RequestTargetCreate) SetTarget(s string) *RequestTargetCreate {
	rtc.mutation.SetTarget(s)
	return rtc
}

// SetPaidAt sets the "paid_at" field.
func (rtc *RequestTargetCreate) SetPaidAt(t time.Time) *RequestTargetCreate {
	rtc.mutation.SetPaidAt(t)
	return rtc
}

// SetNillablePaidAt sets the "paid_at" field if the given value is not nil.
func (rtc *RequestTargetCreate) SetNillablePaidAt(t *time.Time) *RequestTargetCreate {
	if t != nil {
		rtc.SetPaidAt(*t)
	}
	return rtc
}

// SetCreatedAt sets the "created_at" field.
func (rtc *RequestTargetCreate) SetCreatedAt(t time.Time) *RequestTargetCreate {
	rtc.mutation.SetCreatedAt(t)
	return rtc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rtc *RequestTargetCreate) SetNillableCreatedAt(t *time.Time) *RequestTargetCreate {
	if t != nil {
		rtc.SetCreatedAt(*t)
	}
	return rtc
}

// SetID sets the "id" field.
func (rtc *RequestTargetCreate) SetID(u uuid.UUID) *RequestTargetCreate {
	rtc.mutation.SetID(u)
	return rtc
}

// SetRequestID sets the "request" edge to the Request entity by ID.
func (rtc *RequestTargetCreate) SetRequestID(id uuid.UUID) *RequestTargetCreate {
	rtc.mutation.SetRequestID(id)
	return rtc
}

// SetRequest sets the "request" edge to the Request entity.
func (rtc *RequestTargetCreate) SetRequest(r *Request) *RequestTargetCreate {
	return rtc.SetRequestID(r.ID)
}

// Mutation returns the RequestTargetMutation object of the builder.
func (rtc *RequestTargetCreate) Mutation() *RequestTargetMutation {
	return rtc.mutation
}

// Save creates the RequestTarget in the database.
func (rtc *RequestTargetCreate) Save(ctx context.Context) (*RequestTarget, error) {
	var (
		err  error
		node *RequestTarget
	)
	rtc.defaults()
	if len(rtc.hooks) == 0 {
		if err = rtc.check(); err != nil {
			return nil, err
		}
		node, err = rtc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RequestTargetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rtc.check(); err != nil {
				return nil, err
			}
			rtc.mutation = mutation
			node, err = rtc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rtc.hooks) - 1; i >= 0; i-- {
			mut = rtc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rtc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rtc *RequestTargetCreate) SaveX(ctx context.Context) *RequestTarget {
	v, err := rtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (rtc *RequestTargetCreate) defaults() {
	if _, ok := rtc.mutation.CreatedAt(); !ok {
		v := requesttarget.DefaultCreatedAt()
		rtc.mutation.SetCreatedAt(v)
	}
	if _, ok := rtc.mutation.ID(); !ok {
		v := requesttarget.DefaultID()
		rtc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rtc *RequestTargetCreate) check() error {
	if _, ok := rtc.mutation.Target(); !ok {
		return &ValidationError{Name: "target", err: errors.New("ent: missing required field \"target\"")}
	}
	if _, ok := rtc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := rtc.mutation.RequestID(); !ok {
		return &ValidationError{Name: "request", err: errors.New("ent: missing required edge \"request\"")}
	}
	return nil
}

func (rtc *RequestTargetCreate) sqlSave(ctx context.Context) (*RequestTarget, error) {
	_node, _spec := rtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rtc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (rtc *RequestTargetCreate) createSpec() (*RequestTarget, *sqlgraph.CreateSpec) {
	var (
		_node = &RequestTarget{config: rtc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: requesttarget.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: requesttarget.FieldID,
			},
		}
	)
	if id, ok := rtc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rtc.mutation.Target(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: requesttarget.FieldTarget,
		})
		_node.Target = value
	}
	if value, ok := rtc.mutation.PaidAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: requesttarget.FieldPaidAt,
		})
		_node.PaidAt = &value
	}
	if value, ok := rtc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: requesttarget.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := rtc.mutation.RequestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   requesttarget.RequestTable,
			Columns: []string{requesttarget.RequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: request.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.request_target = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RequestTargetCreateBulk is the builder for creating many RequestTarget entities in bulk.
type RequestTargetCreateBulk struct {
	config
	builders []*RequestTargetCreate
}

// Save creates the RequestTarget entities in the database.
func (rtcb *RequestTargetCreateBulk) Save(ctx context.Context) ([]*RequestTarget, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rtcb.builders))
	nodes := make([]*RequestTarget, len(rtcb.builders))
	mutators := make([]Mutator, len(rtcb.builders))
	for i := range rtcb.builders {
		func(i int, root context.Context) {
			builder := rtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RequestTargetMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rtcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rtcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rtcb *RequestTargetCreateBulk) SaveX(ctx context.Context) []*RequestTarget {
	v, err := rtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
