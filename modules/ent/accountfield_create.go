// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"tammy/ent/account"
	"tammy/ent/accountfield"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccountFieldCreate is the builder for creating a AccountField entity.
type AccountFieldCreate struct {
	config
	mutation *AccountFieldMutation
	hooks    []Hook
}

// SetFid sets the "fid" field.
func (afc *AccountFieldCreate) SetFid(i int) *AccountFieldCreate {
	afc.mutation.SetFid(i)
	return afc
}

// SetKey sets the "key" field.
func (afc *AccountFieldCreate) SetKey(s string) *AccountFieldCreate {
	afc.mutation.SetKey(s)
	return afc
}

// SetValue sets the "value" field.
func (afc *AccountFieldCreate) SetValue(s string) *AccountFieldCreate {
	afc.mutation.SetValue(s)
	return afc
}

// SetID sets the "id" field.
func (afc *AccountFieldCreate) SetID(i int) *AccountFieldCreate {
	afc.mutation.SetID(i)
	return afc
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (afc *AccountFieldCreate) SetAccountID(id int) *AccountFieldCreate {
	afc.mutation.SetAccountID(id)
	return afc
}

// SetAccount sets the "account" edge to the Account entity.
func (afc *AccountFieldCreate) SetAccount(a *Account) *AccountFieldCreate {
	return afc.SetAccountID(a.ID)
}

// Mutation returns the AccountFieldMutation object of the builder.
func (afc *AccountFieldCreate) Mutation() *AccountFieldMutation {
	return afc.mutation
}

// Save creates the AccountField in the database.
func (afc *AccountFieldCreate) Save(ctx context.Context) (*AccountField, error) {
	var (
		err  error
		node *AccountField
	)
	if len(afc.hooks) == 0 {
		if err = afc.check(); err != nil {
			return nil, err
		}
		node, err = afc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountFieldMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = afc.check(); err != nil {
				return nil, err
			}
			afc.mutation = mutation
			if node, err = afc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(afc.hooks) - 1; i >= 0; i-- {
			if afc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = afc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, afc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AccountField)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AccountFieldMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (afc *AccountFieldCreate) SaveX(ctx context.Context) *AccountField {
	v, err := afc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (afc *AccountFieldCreate) Exec(ctx context.Context) error {
	_, err := afc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (afc *AccountFieldCreate) ExecX(ctx context.Context) {
	if err := afc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (afc *AccountFieldCreate) check() error {
	if _, ok := afc.mutation.Fid(); !ok {
		return &ValidationError{Name: "fid", err: errors.New(`ent: missing required field "AccountField.fid"`)}
	}
	if _, ok := afc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "AccountField.key"`)}
	}
	if _, ok := afc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "AccountField.value"`)}
	}
	if _, ok := afc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account", err: errors.New(`ent: missing required edge "AccountField.account"`)}
	}
	return nil
}

func (afc *AccountFieldCreate) sqlSave(ctx context.Context) (*AccountField, error) {
	_node, _spec := afc.createSpec()
	if err := sqlgraph.CreateNode(ctx, afc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	return _node, nil
}

func (afc *AccountFieldCreate) createSpec() (*AccountField, *sqlgraph.CreateSpec) {
	var (
		_node = &AccountField{config: afc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: accountfield.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: accountfield.FieldID,
			},
		}
	)
	if id, ok := afc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := afc.mutation.Fid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: accountfield.FieldFid,
		})
		_node.Fid = value
	}
	if value, ok := afc.mutation.Key(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accountfield.FieldKey,
		})
		_node.Key = value
	}
	if value, ok := afc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accountfield.FieldValue,
		})
		_node.Value = value
	}
	if nodes := afc.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accountfield.AccountTable,
			Columns: []string{accountfield.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.account_fields = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AccountFieldCreateBulk is the builder for creating many AccountField entities in bulk.
type AccountFieldCreateBulk struct {
	config
	builders []*AccountFieldCreate
}

// Save creates the AccountField entities in the database.
func (afcb *AccountFieldCreateBulk) Save(ctx context.Context) ([]*AccountField, error) {
	specs := make([]*sqlgraph.CreateSpec, len(afcb.builders))
	nodes := make([]*AccountField, len(afcb.builders))
	mutators := make([]Mutator, len(afcb.builders))
	for i := range afcb.builders {
		func(i int, root context.Context) {
			builder := afcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccountFieldMutation)
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
					_, err = mutators[i+1].Mutate(root, afcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, afcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, afcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (afcb *AccountFieldCreateBulk) SaveX(ctx context.Context) []*AccountField {
	v, err := afcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (afcb *AccountFieldCreateBulk) Exec(ctx context.Context) error {
	_, err := afcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (afcb *AccountFieldCreateBulk) ExecX(ctx context.Context) {
	if err := afcb.Exec(ctx); err != nil {
		panic(err)
	}
}
