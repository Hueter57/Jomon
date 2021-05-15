// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/traPtitech/Jomon/ent/predicate"
	"github.com/traPtitech/Jomon/ent/requesttarget"
)

// RequestTargetDelete is the builder for deleting a RequestTarget entity.
type RequestTargetDelete struct {
	config
	hooks    []Hook
	mutation *RequestTargetMutation
}

// Where adds a new predicate to the RequestTargetDelete builder.
func (rtd *RequestTargetDelete) Where(ps ...predicate.RequestTarget) *RequestTargetDelete {
	rtd.mutation.predicates = append(rtd.mutation.predicates, ps...)
	return rtd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rtd *RequestTargetDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rtd.hooks) == 0 {
		affected, err = rtd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RequestTargetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rtd.mutation = mutation
			affected, err = rtd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rtd.hooks) - 1; i >= 0; i-- {
			mut = rtd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rtd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtd *RequestTargetDelete) ExecX(ctx context.Context) int {
	n, err := rtd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rtd *RequestTargetDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: requesttarget.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: requesttarget.FieldID,
			},
		},
	}
	if ps := rtd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, rtd.driver, _spec)
}

// RequestTargetDeleteOne is the builder for deleting a single RequestTarget entity.
type RequestTargetDeleteOne struct {
	rtd *RequestTargetDelete
}

// Exec executes the deletion query.
func (rtdo *RequestTargetDeleteOne) Exec(ctx context.Context) error {
	n, err := rtdo.rtd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{requesttarget.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rtdo *RequestTargetDeleteOne) ExecX(ctx context.Context) {
	rtdo.rtd.ExecX(ctx)
}
