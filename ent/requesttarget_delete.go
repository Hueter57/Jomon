// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

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

// Where appends a list predicates to the RequestTargetDelete builder.
func (rtd *RequestTargetDelete) Where(ps ...predicate.RequestTarget) *RequestTargetDelete {
	rtd.mutation.Where(ps...)
	return rtd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rtd *RequestTargetDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, RequestTargetMutation](ctx, rtd.sqlExec, rtd.mutation, rtd.hooks)
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
	_spec := sqlgraph.NewDeleteSpec(requesttarget.Table, sqlgraph.NewFieldSpec(requesttarget.FieldID, field.TypeUUID))
	if ps := rtd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rtd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rtd.mutation.done = true
	return affected, err
}

// RequestTargetDeleteOne is the builder for deleting a single RequestTarget entity.
type RequestTargetDeleteOne struct {
	rtd *RequestTargetDelete
}

// Where appends a list predicates to the RequestTargetDelete builder.
func (rtdo *RequestTargetDeleteOne) Where(ps ...predicate.RequestTarget) *RequestTargetDeleteOne {
	rtdo.rtd.mutation.Where(ps...)
	return rtdo
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
	if err := rtdo.Exec(ctx); err != nil {
		panic(err)
	}
}
