// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"tammy/ent/account"
	"tammy/ent/portal"
	"tammy/ent/portallegal"
	"tammy/ent/portalmetadata"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PortalCreate is the builder for creating a Portal entity.
type PortalCreate struct {
	config
	mutation *PortalMutation
	hooks    []Hook
}

// SetCreatedAt sets the "createdAt" field.
func (pc *PortalCreate) SetCreatedAt(t time.Time) *PortalCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (pc *PortalCreate) SetNillableCreatedAt(t *time.Time) *PortalCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updatedAt" field.
func (pc *PortalCreate) SetUpdatedAt(t time.Time) *PortalCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (pc *PortalCreate) SetNillableUpdatedAt(t *time.Time) *PortalCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetIsActive sets the "isActive" field.
func (pc *PortalCreate) SetIsActive(b bool) *PortalCreate {
	pc.mutation.SetIsActive(b)
	return pc
}

// SetNillableIsActive sets the "isActive" field if the given value is not nil.
func (pc *PortalCreate) SetNillableIsActive(b *bool) *PortalCreate {
	if b != nil {
		pc.SetIsActive(*b)
	}
	return pc
}

// SetSlug sets the "slug" field.
func (pc *PortalCreate) SetSlug(s string) *PortalCreate {
	pc.mutation.SetSlug(s)
	return pc
}

// SetID sets the "id" field.
func (pc *PortalCreate) SetID(i int) *PortalCreate {
	pc.mutation.SetID(i)
	return pc
}

// AddMemberIDs adds the "members" edge to the Account entity by IDs.
func (pc *PortalCreate) AddMemberIDs(ids ...int) *PortalCreate {
	pc.mutation.AddMemberIDs(ids...)
	return pc
}

// AddMembers adds the "members" edges to the Account entity.
func (pc *PortalCreate) AddMembers(a ...*Account) *PortalCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pc.AddMemberIDs(ids...)
}

// AddMetadatumIDs adds the "metadata" edge to the PortalMetadata entity by IDs.
func (pc *PortalCreate) AddMetadatumIDs(ids ...int) *PortalCreate {
	pc.mutation.AddMetadatumIDs(ids...)
	return pc
}

// AddMetadata adds the "metadata" edges to the PortalMetadata entity.
func (pc *PortalCreate) AddMetadata(p ...*PortalMetadata) *PortalCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddMetadatumIDs(ids...)
}

// AddLegalIDs adds the "legal" edge to the PortalLegal entity by IDs.
func (pc *PortalCreate) AddLegalIDs(ids ...int) *PortalCreate {
	pc.mutation.AddLegalIDs(ids...)
	return pc
}

// AddLegal adds the "legal" edges to the PortalLegal entity.
func (pc *PortalCreate) AddLegal(p ...*PortalLegal) *PortalCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddLegalIDs(ids...)
}

// Mutation returns the PortalMutation object of the builder.
func (pc *PortalCreate) Mutation() *PortalMutation {
	return pc.mutation
}

// Save creates the Portal in the database.
func (pc *PortalCreate) Save(ctx context.Context) (*Portal, error) {
	var (
		err  error
		node *Portal
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PortalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Portal)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PortalMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PortalCreate) SaveX(ctx context.Context) *Portal {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PortalCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PortalCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PortalCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := portal.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := portal.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.IsActive(); !ok {
		v := portal.DefaultIsActive
		pc.mutation.SetIsActive(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PortalCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "Portal.createdAt"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "Portal.updatedAt"`)}
	}
	if _, ok := pc.mutation.IsActive(); !ok {
		return &ValidationError{Name: "isActive", err: errors.New(`ent: missing required field "Portal.isActive"`)}
	}
	if _, ok := pc.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "Portal.slug"`)}
	}
	return nil
}

func (pc *PortalCreate) sqlSave(ctx context.Context) (*Portal, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
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

func (pc *PortalCreate) createSpec() (*Portal, *sqlgraph.CreateSpec) {
	var (
		_node = &Portal{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: portal.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: portal.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: portal.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: portal.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.IsActive(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: portal.FieldIsActive,
		})
		_node.IsActive = value
	}
	if value, ok := pc.mutation.Slug(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: portal.FieldSlug,
		})
		_node.Slug = value
	}
	if nodes := pc.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   portal.MembersTable,
			Columns: portal.MembersPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   portal.MetadataTable,
			Columns: []string{portal.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: portalmetadata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.LegalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   portal.LegalTable,
			Columns: []string{portal.LegalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: portallegal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PortalCreateBulk is the builder for creating many Portal entities in bulk.
type PortalCreateBulk struct {
	config
	builders []*PortalCreate
}

// Save creates the Portal entities in the database.
func (pcb *PortalCreateBulk) Save(ctx context.Context) ([]*Portal, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Portal, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PortalMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PortalCreateBulk) SaveX(ctx context.Context) []*Portal {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PortalCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PortalCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}