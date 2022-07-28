// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"tammy/ent/predicate"
	"tammy/ent/userpassword"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserPasswordDelete is the builder for deleting a UserPassword entity.
type UserPasswordDelete struct {
	config
	hooks    []Hook
	mutation *UserPasswordMutation
}

// Where appends a list predicates to the UserPasswordDelete builder.
func (upd *UserPasswordDelete) Where(ps ...predicate.UserPassword) *UserPasswordDelete {
	upd.mutation.Where(ps...)
	return upd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (upd *UserPasswordDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(upd.hooks) == 0 {
		affected, err = upd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserPasswordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			upd.mutation = mutation
			affected, err = upd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(upd.hooks) - 1; i >= 0; i-- {
			if upd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = upd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, upd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (upd *UserPasswordDelete) ExecX(ctx context.Context) int {
	n, err := upd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (upd *UserPasswordDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: userpassword.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userpassword.FieldID,
			},
		},
	}
	if ps := upd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, upd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// UserPasswordDeleteOne is the builder for deleting a single UserPassword entity.
type UserPasswordDeleteOne struct {
	upd *UserPasswordDelete
}

// Exec executes the deletion query.
func (updo *UserPasswordDeleteOne) Exec(ctx context.Context) error {
	n, err := updo.upd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userpassword.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (updo *UserPasswordDeleteOne) ExecX(ctx context.Context) {
	updo.upd.ExecX(ctx)
}
