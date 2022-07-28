// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"tammy/ent/account"
	"tammy/ent/accountfield"
	"tammy/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccountFieldUpdate is the builder for updating AccountField entities.
type AccountFieldUpdate struct {
	config
	hooks    []Hook
	mutation *AccountFieldMutation
}

// Where appends a list predicates to the AccountFieldUpdate builder.
func (afu *AccountFieldUpdate) Where(ps ...predicate.AccountField) *AccountFieldUpdate {
	afu.mutation.Where(ps...)
	return afu
}

// SetFid sets the "fid" field.
func (afu *AccountFieldUpdate) SetFid(i int) *AccountFieldUpdate {
	afu.mutation.ResetFid()
	afu.mutation.SetFid(i)
	return afu
}

// AddFid adds i to the "fid" field.
func (afu *AccountFieldUpdate) AddFid(i int) *AccountFieldUpdate {
	afu.mutation.AddFid(i)
	return afu
}

// SetKey sets the "key" field.
func (afu *AccountFieldUpdate) SetKey(s string) *AccountFieldUpdate {
	afu.mutation.SetKey(s)
	return afu
}

// SetValue sets the "value" field.
func (afu *AccountFieldUpdate) SetValue(s string) *AccountFieldUpdate {
	afu.mutation.SetValue(s)
	return afu
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (afu *AccountFieldUpdate) SetAccountID(id int) *AccountFieldUpdate {
	afu.mutation.SetAccountID(id)
	return afu
}

// SetAccount sets the "account" edge to the Account entity.
func (afu *AccountFieldUpdate) SetAccount(a *Account) *AccountFieldUpdate {
	return afu.SetAccountID(a.ID)
}

// Mutation returns the AccountFieldMutation object of the builder.
func (afu *AccountFieldUpdate) Mutation() *AccountFieldMutation {
	return afu.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (afu *AccountFieldUpdate) ClearAccount() *AccountFieldUpdate {
	afu.mutation.ClearAccount()
	return afu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (afu *AccountFieldUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(afu.hooks) == 0 {
		if err = afu.check(); err != nil {
			return 0, err
		}
		affected, err = afu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountFieldMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = afu.check(); err != nil {
				return 0, err
			}
			afu.mutation = mutation
			affected, err = afu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(afu.hooks) - 1; i >= 0; i-- {
			if afu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = afu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, afu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (afu *AccountFieldUpdate) SaveX(ctx context.Context) int {
	affected, err := afu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (afu *AccountFieldUpdate) Exec(ctx context.Context) error {
	_, err := afu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (afu *AccountFieldUpdate) ExecX(ctx context.Context) {
	if err := afu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (afu *AccountFieldUpdate) check() error {
	if _, ok := afu.mutation.AccountID(); afu.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AccountField.account"`)
	}
	return nil
}

func (afu *AccountFieldUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   accountfield.Table,
			Columns: accountfield.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: accountfield.FieldID,
			},
		},
	}
	if ps := afu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := afu.mutation.Fid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: accountfield.FieldFid,
		})
	}
	if value, ok := afu.mutation.AddedFid(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: accountfield.FieldFid,
		})
	}
	if value, ok := afu.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accountfield.FieldKey,
		})
	}
	if value, ok := afu.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accountfield.FieldValue,
		})
	}
	if afu.mutation.AccountCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := afu.mutation.AccountIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, afu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accountfield.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AccountFieldUpdateOne is the builder for updating a single AccountField entity.
type AccountFieldUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccountFieldMutation
}

// SetFid sets the "fid" field.
func (afuo *AccountFieldUpdateOne) SetFid(i int) *AccountFieldUpdateOne {
	afuo.mutation.ResetFid()
	afuo.mutation.SetFid(i)
	return afuo
}

// AddFid adds i to the "fid" field.
func (afuo *AccountFieldUpdateOne) AddFid(i int) *AccountFieldUpdateOne {
	afuo.mutation.AddFid(i)
	return afuo
}

// SetKey sets the "key" field.
func (afuo *AccountFieldUpdateOne) SetKey(s string) *AccountFieldUpdateOne {
	afuo.mutation.SetKey(s)
	return afuo
}

// SetValue sets the "value" field.
func (afuo *AccountFieldUpdateOne) SetValue(s string) *AccountFieldUpdateOne {
	afuo.mutation.SetValue(s)
	return afuo
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (afuo *AccountFieldUpdateOne) SetAccountID(id int) *AccountFieldUpdateOne {
	afuo.mutation.SetAccountID(id)
	return afuo
}

// SetAccount sets the "account" edge to the Account entity.
func (afuo *AccountFieldUpdateOne) SetAccount(a *Account) *AccountFieldUpdateOne {
	return afuo.SetAccountID(a.ID)
}

// Mutation returns the AccountFieldMutation object of the builder.
func (afuo *AccountFieldUpdateOne) Mutation() *AccountFieldMutation {
	return afuo.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (afuo *AccountFieldUpdateOne) ClearAccount() *AccountFieldUpdateOne {
	afuo.mutation.ClearAccount()
	return afuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (afuo *AccountFieldUpdateOne) Select(field string, fields ...string) *AccountFieldUpdateOne {
	afuo.fields = append([]string{field}, fields...)
	return afuo
}

// Save executes the query and returns the updated AccountField entity.
func (afuo *AccountFieldUpdateOne) Save(ctx context.Context) (*AccountField, error) {
	var (
		err  error
		node *AccountField
	)
	if len(afuo.hooks) == 0 {
		if err = afuo.check(); err != nil {
			return nil, err
		}
		node, err = afuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountFieldMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = afuo.check(); err != nil {
				return nil, err
			}
			afuo.mutation = mutation
			node, err = afuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(afuo.hooks) - 1; i >= 0; i-- {
			if afuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = afuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, afuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (afuo *AccountFieldUpdateOne) SaveX(ctx context.Context) *AccountField {
	node, err := afuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (afuo *AccountFieldUpdateOne) Exec(ctx context.Context) error {
	_, err := afuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (afuo *AccountFieldUpdateOne) ExecX(ctx context.Context) {
	if err := afuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (afuo *AccountFieldUpdateOne) check() error {
	if _, ok := afuo.mutation.AccountID(); afuo.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AccountField.account"`)
	}
	return nil
}

func (afuo *AccountFieldUpdateOne) sqlSave(ctx context.Context) (_node *AccountField, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   accountfield.Table,
			Columns: accountfield.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: accountfield.FieldID,
			},
		},
	}
	id, ok := afuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AccountField.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := afuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accountfield.FieldID)
		for _, f := range fields {
			if !accountfield.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != accountfield.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := afuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := afuo.mutation.Fid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: accountfield.FieldFid,
		})
	}
	if value, ok := afuo.mutation.AddedFid(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: accountfield.FieldFid,
		})
	}
	if value, ok := afuo.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accountfield.FieldKey,
		})
	}
	if value, ok := afuo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accountfield.FieldValue,
		})
	}
	if afuo.mutation.AccountCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := afuo.mutation.AccountIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AccountField{config: afuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, afuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accountfield.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
