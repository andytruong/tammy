// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"tammy/ent/accountfield"
	"tammy/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccountFieldDelete is the builder for deleting a AccountField entity.
type AccountFieldDelete struct {
	config
	hooks    []Hook
	mutation *AccountFieldMutation
}

// Where appends a list predicates to the AccountFieldDelete builder.
func (afd *AccountFieldDelete) Where(ps ...predicate.AccountField) *AccountFieldDelete {
	afd.mutation.Where(ps...)
	return afd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (afd *AccountFieldDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(afd.hooks) == 0 {
		affected, err = afd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountFieldMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			afd.mutation = mutation
			affected, err = afd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(afd.hooks) - 1; i >= 0; i-- {
			if afd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = afd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, afd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (afd *AccountFieldDelete) ExecX(ctx context.Context) int {
	n, err := afd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (afd *AccountFieldDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: accountfield.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: accountfield.FieldID,
			},
		},
	}
	if ps := afd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, afd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// AccountFieldDeleteOne is the builder for deleting a single AccountField entity.
type AccountFieldDeleteOne struct {
	afd *AccountFieldDelete
}

// Exec executes the deletion query.
func (afdo *AccountFieldDeleteOne) Exec(ctx context.Context) error {
	n, err := afdo.afd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{accountfield.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (afdo *AccountFieldDeleteOne) ExecX(ctx context.Context) {
	afdo.afd.ExecX(ctx)
}
