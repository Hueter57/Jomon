// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/traPtitech/Jomon/ent/groupbudget"
	"github.com/traPtitech/Jomon/ent/predicate"
)

// GroupBudgetDelete is the builder for deleting a GroupBudget entity.
type GroupBudgetDelete struct {
	config
	hooks    []Hook
	mutation *GroupBudgetMutation
}

// Where appends a list predicates to the GroupBudgetDelete builder.
func (gbd *GroupBudgetDelete) Where(ps ...predicate.GroupBudget) *GroupBudgetDelete {
	gbd.mutation.Where(ps...)
	return gbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gbd *GroupBudgetDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, GroupBudgetMutation](ctx, gbd.sqlExec, gbd.mutation, gbd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (gbd *GroupBudgetDelete) ExecX(ctx context.Context) int {
	n, err := gbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gbd *GroupBudgetDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(groupbudget.Table, sqlgraph.NewFieldSpec(groupbudget.FieldID, field.TypeUUID))
	if ps := gbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gbd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	gbd.mutation.done = true
	return affected, err
}

// GroupBudgetDeleteOne is the builder for deleting a single GroupBudget entity.
type GroupBudgetDeleteOne struct {
	gbd *GroupBudgetDelete
}

// Where appends a list predicates to the GroupBudgetDelete builder.
func (gbdo *GroupBudgetDeleteOne) Where(ps ...predicate.GroupBudget) *GroupBudgetDeleteOne {
	gbdo.gbd.mutation.Where(ps...)
	return gbdo
}

// Exec executes the deletion query.
func (gbdo *GroupBudgetDeleteOne) Exec(ctx context.Context) error {
	n, err := gbdo.gbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{groupbudget.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gbdo *GroupBudgetDeleteOne) ExecX(ctx context.Context) {
	if err := gbdo.Exec(ctx); err != nil {
		panic(err)
	}
}
