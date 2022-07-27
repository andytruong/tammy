// Code generated by ent, DO NOT EDIT.

package store

import (
	"context"
	"errors"
	"fmt"
	"tammy/pkg/store/portal"
	"tammy/pkg/store/portallegal"
	"tammy/pkg/store/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PortalLegalUpdate is the builder for updating PortalLegal entities.
type PortalLegalUpdate struct {
	config
	hooks    []Hook
	mutation *PortalLegalMutation
}

// Where appends a list predicates to the PortalLegalUpdate builder.
func (plu *PortalLegalUpdate) Where(ps ...predicate.PortalLegal) *PortalLegalUpdate {
	plu.mutation.Where(ps...)
	return plu
}

// SetUpdatedAt sets the "updatedAt" field.
func (plu *PortalLegalUpdate) SetUpdatedAt(t time.Time) *PortalLegalUpdate {
	plu.mutation.SetUpdatedAt(t)
	return plu
}

// SetPrivacyPolicy sets the "privacyPolicy" field.
func (plu *PortalLegalUpdate) SetPrivacyPolicy(s string) *PortalLegalUpdate {
	plu.mutation.SetPrivacyPolicy(s)
	return plu
}

// SetTermOfService sets the "termOfService" field.
func (plu *PortalLegalUpdate) SetTermOfService(s string) *PortalLegalUpdate {
	plu.mutation.SetTermOfService(s)
	return plu
}

// SetCopyright sets the "copyright" field.
func (plu *PortalLegalUpdate) SetCopyright(s string) *PortalLegalUpdate {
	plu.mutation.SetCopyright(s)
	return plu
}

// SetOnlineTrainingAgreement sets the "onlineTrainingAgreement" field.
func (plu *PortalLegalUpdate) SetOnlineTrainingAgreement(s string) *PortalLegalUpdate {
	plu.mutation.SetOnlineTrainingAgreement(s)
	return plu
}

// SetPortalID sets the "portal" edge to the Portal entity by ID.
func (plu *PortalLegalUpdate) SetPortalID(id uint32) *PortalLegalUpdate {
	plu.mutation.SetPortalID(id)
	return plu
}

// SetPortal sets the "portal" edge to the Portal entity.
func (plu *PortalLegalUpdate) SetPortal(p *Portal) *PortalLegalUpdate {
	return plu.SetPortalID(p.ID)
}

// Mutation returns the PortalLegalMutation object of the builder.
func (plu *PortalLegalUpdate) Mutation() *PortalLegalMutation {
	return plu.mutation
}

// ClearPortal clears the "portal" edge to the Portal entity.
func (plu *PortalLegalUpdate) ClearPortal() *PortalLegalUpdate {
	plu.mutation.ClearPortal()
	return plu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (plu *PortalLegalUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(plu.hooks) == 0 {
		if err = plu.check(); err != nil {
			return 0, err
		}
		affected, err = plu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PortalLegalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = plu.check(); err != nil {
				return 0, err
			}
			plu.mutation = mutation
			affected, err = plu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(plu.hooks) - 1; i >= 0; i-- {
			if plu.hooks[i] == nil {
				return 0, fmt.Errorf("store: uninitialized hook (forgotten import store/runtime?)")
			}
			mut = plu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, plu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (plu *PortalLegalUpdate) SaveX(ctx context.Context) int {
	affected, err := plu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (plu *PortalLegalUpdate) Exec(ctx context.Context) error {
	_, err := plu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (plu *PortalLegalUpdate) ExecX(ctx context.Context) {
	if err := plu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (plu *PortalLegalUpdate) check() error {
	if _, ok := plu.mutation.PortalID(); plu.mutation.PortalCleared() && !ok {
		return errors.New(`store: clearing a required unique edge "PortalLegal.portal"`)
	}
	return nil
}

func (plu *PortalLegalUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   portallegal.Table,
			Columns: portallegal.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: portallegal.FieldID,
			},
		},
	}
	if ps := plu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := plu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: portallegal.FieldUpdatedAt,
		})
	}
	if value, ok := plu.mutation.PrivacyPolicy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: portallegal.FieldPrivacyPolicy,
		})
	}
	if value, ok := plu.mutation.TermOfService(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: portallegal.FieldTermOfService,
		})
	}
	if value, ok := plu.mutation.Copyright(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: portallegal.FieldCopyright,
		})
	}
	if value, ok := plu.mutation.OnlineTrainingAgreement(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: portallegal.FieldOnlineTrainingAgreement,
		})
	}
	if plu.mutation.PortalCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   portallegal.PortalTable,
			Columns: []string{portallegal.PortalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint32,
					Column: portal.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := plu.mutation.PortalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   portallegal.PortalTable,
			Columns: []string{portallegal.PortalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint32,
					Column: portal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, plu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{portallegal.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PortalLegalUpdateOne is the builder for updating a single PortalLegal entity.
type PortalLegalUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PortalLegalMutation
}

// SetUpdatedAt sets the "updatedAt" field.
func (pluo *PortalLegalUpdateOne) SetUpdatedAt(t time.Time) *PortalLegalUpdateOne {
	pluo.mutation.SetUpdatedAt(t)
	return pluo
}

// SetPrivacyPolicy sets the "privacyPolicy" field.
func (pluo *PortalLegalUpdateOne) SetPrivacyPolicy(s string) *PortalLegalUpdateOne {
	pluo.mutation.SetPrivacyPolicy(s)
	return pluo
}

// SetTermOfService sets the "termOfService" field.
func (pluo *PortalLegalUpdateOne) SetTermOfService(s string) *PortalLegalUpdateOne {
	pluo.mutation.SetTermOfService(s)
	return pluo
}

// SetCopyright sets the "copyright" field.
func (pluo *PortalLegalUpdateOne) SetCopyright(s string) *PortalLegalUpdateOne {
	pluo.mutation.SetCopyright(s)
	return pluo
}

// SetOnlineTrainingAgreement sets the "onlineTrainingAgreement" field.
func (pluo *PortalLegalUpdateOne) SetOnlineTrainingAgreement(s string) *PortalLegalUpdateOne {
	pluo.mutation.SetOnlineTrainingAgreement(s)
	return pluo
}

// SetPortalID sets the "portal" edge to the Portal entity by ID.
func (pluo *PortalLegalUpdateOne) SetPortalID(id uint32) *PortalLegalUpdateOne {
	pluo.mutation.SetPortalID(id)
	return pluo
}

// SetPortal sets the "portal" edge to the Portal entity.
func (pluo *PortalLegalUpdateOne) SetPortal(p *Portal) *PortalLegalUpdateOne {
	return pluo.SetPortalID(p.ID)
}

// Mutation returns the PortalLegalMutation object of the builder.
func (pluo *PortalLegalUpdateOne) Mutation() *PortalLegalMutation {
	return pluo.mutation
}

// ClearPortal clears the "portal" edge to the Portal entity.
func (pluo *PortalLegalUpdateOne) ClearPortal() *PortalLegalUpdateOne {
	pluo.mutation.ClearPortal()
	return pluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pluo *PortalLegalUpdateOne) Select(field string, fields ...string) *PortalLegalUpdateOne {
	pluo.fields = append([]string{field}, fields...)
	return pluo
}

// Save executes the query and returns the updated PortalLegal entity.
func (pluo *PortalLegalUpdateOne) Save(ctx context.Context) (*PortalLegal, error) {
	var (
		err  error
		node *PortalLegal
	)
	if len(pluo.hooks) == 0 {
		if err = pluo.check(); err != nil {
			return nil, err
		}
		node, err = pluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PortalLegalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pluo.check(); err != nil {
				return nil, err
			}
			pluo.mutation = mutation
			node, err = pluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pluo.hooks) - 1; i >= 0; i-- {
			if pluo.hooks[i] == nil {
				return nil, fmt.Errorf("store: uninitialized hook (forgotten import store/runtime?)")
			}
			mut = pluo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pluo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*PortalLegal)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PortalLegalMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pluo *PortalLegalUpdateOne) SaveX(ctx context.Context) *PortalLegal {
	node, err := pluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pluo *PortalLegalUpdateOne) Exec(ctx context.Context) error {
	_, err := pluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pluo *PortalLegalUpdateOne) ExecX(ctx context.Context) {
	if err := pluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pluo *PortalLegalUpdateOne) check() error {
	if _, ok := pluo.mutation.PortalID(); pluo.mutation.PortalCleared() && !ok {
		return errors.New(`store: clearing a required unique edge "PortalLegal.portal"`)
	}
	return nil
}

func (pluo *PortalLegalUpdateOne) sqlSave(ctx context.Context) (_node *PortalLegal, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   portallegal.Table,
			Columns: portallegal.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: portallegal.FieldID,
			},
		},
	}
	id, ok := pluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`store: missing "PortalLegal.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, portallegal.FieldID)
		for _, f := range fields {
			if !portallegal.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("store: invalid field %q for query", f)}
			}
			if f != portallegal.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pluo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: portallegal.FieldUpdatedAt,
		})
	}
	if value, ok := pluo.mutation.PrivacyPolicy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: portallegal.FieldPrivacyPolicy,
		})
	}
	if value, ok := pluo.mutation.TermOfService(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: portallegal.FieldTermOfService,
		})
	}
	if value, ok := pluo.mutation.Copyright(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: portallegal.FieldCopyright,
		})
	}
	if value, ok := pluo.mutation.OnlineTrainingAgreement(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: portallegal.FieldOnlineTrainingAgreement,
		})
	}
	if pluo.mutation.PortalCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   portallegal.PortalTable,
			Columns: []string{portallegal.PortalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint32,
					Column: portal.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pluo.mutation.PortalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   portallegal.PortalTable,
			Columns: []string{portallegal.PortalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint32,
					Column: portal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PortalLegal{config: pluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{portallegal.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}