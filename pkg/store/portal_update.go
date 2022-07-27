// Code generated by ent, DO NOT EDIT.

package store

import (
	"context"
	"errors"
	"fmt"
	"tammy/pkg/store/account"
	"tammy/pkg/store/portal"
	"tammy/pkg/store/portallegal"
	"tammy/pkg/store/portalmetadata"
	"tammy/pkg/store/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PortalUpdate is the builder for updating Portal entities.
type PortalUpdate struct {
	config
	hooks    []Hook
	mutation *PortalMutation
}

// Where appends a list predicates to the PortalUpdate builder.
func (pu *PortalUpdate) Where(ps ...predicate.Portal) *PortalUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetIsActive sets the "isActive" field.
func (pu *PortalUpdate) SetIsActive(b bool) *PortalUpdate {
	pu.mutation.SetIsActive(b)
	return pu
}

// SetNillableIsActive sets the "isActive" field if the given value is not nil.
func (pu *PortalUpdate) SetNillableIsActive(b *bool) *PortalUpdate {
	if b != nil {
		pu.SetIsActive(*b)
	}
	return pu
}

// AddMemberIDs adds the "members" edge to the Account entity by IDs.
func (pu *PortalUpdate) AddMemberIDs(ids ...uint32) *PortalUpdate {
	pu.mutation.AddMemberIDs(ids...)
	return pu
}

// AddMembers adds the "members" edges to the Account entity.
func (pu *PortalUpdate) AddMembers(a ...*Account) *PortalUpdate {
	ids := make([]uint32, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pu.AddMemberIDs(ids...)
}

// AddMetadatumIDs adds the "metadata" edge to the PortalMetadata entity by IDs.
func (pu *PortalUpdate) AddMetadatumIDs(ids ...uint32) *PortalUpdate {
	pu.mutation.AddMetadatumIDs(ids...)
	return pu
}

// AddMetadata adds the "metadata" edges to the PortalMetadata entity.
func (pu *PortalUpdate) AddMetadata(p ...*PortalMetadata) *PortalUpdate {
	ids := make([]uint32, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddMetadatumIDs(ids...)
}

// AddLegalIDs adds the "legal" edge to the PortalLegal entity by IDs.
func (pu *PortalUpdate) AddLegalIDs(ids ...uint32) *PortalUpdate {
	pu.mutation.AddLegalIDs(ids...)
	return pu
}

// AddLegal adds the "legal" edges to the PortalLegal entity.
func (pu *PortalUpdate) AddLegal(p ...*PortalLegal) *PortalUpdate {
	ids := make([]uint32, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddLegalIDs(ids...)
}

// Mutation returns the PortalMutation object of the builder.
func (pu *PortalUpdate) Mutation() *PortalMutation {
	return pu.mutation
}

// ClearMembers clears all "members" edges to the Account entity.
func (pu *PortalUpdate) ClearMembers() *PortalUpdate {
	pu.mutation.ClearMembers()
	return pu
}

// RemoveMemberIDs removes the "members" edge to Account entities by IDs.
func (pu *PortalUpdate) RemoveMemberIDs(ids ...uint32) *PortalUpdate {
	pu.mutation.RemoveMemberIDs(ids...)
	return pu
}

// RemoveMembers removes "members" edges to Account entities.
func (pu *PortalUpdate) RemoveMembers(a ...*Account) *PortalUpdate {
	ids := make([]uint32, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pu.RemoveMemberIDs(ids...)
}

// ClearMetadata clears all "metadata" edges to the PortalMetadata entity.
func (pu *PortalUpdate) ClearMetadata() *PortalUpdate {
	pu.mutation.ClearMetadata()
	return pu
}

// RemoveMetadatumIDs removes the "metadata" edge to PortalMetadata entities by IDs.
func (pu *PortalUpdate) RemoveMetadatumIDs(ids ...uint32) *PortalUpdate {
	pu.mutation.RemoveMetadatumIDs(ids...)
	return pu
}

// RemoveMetadata removes "metadata" edges to PortalMetadata entities.
func (pu *PortalUpdate) RemoveMetadata(p ...*PortalMetadata) *PortalUpdate {
	ids := make([]uint32, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveMetadatumIDs(ids...)
}

// ClearLegal clears all "legal" edges to the PortalLegal entity.
func (pu *PortalUpdate) ClearLegal() *PortalUpdate {
	pu.mutation.ClearLegal()
	return pu
}

// RemoveLegalIDs removes the "legal" edge to PortalLegal entities by IDs.
func (pu *PortalUpdate) RemoveLegalIDs(ids ...uint32) *PortalUpdate {
	pu.mutation.RemoveLegalIDs(ids...)
	return pu
}

// RemoveLegal removes "legal" edges to PortalLegal entities.
func (pu *PortalUpdate) RemoveLegal(p ...*PortalLegal) *PortalUpdate {
	ids := make([]uint32, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveLegalIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PortalUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PortalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("store: uninitialized hook (forgotten import store/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PortalUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PortalUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PortalUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PortalUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   portal.Table,
			Columns: portal.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: portal.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.IsActive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: portal.FieldIsActive,
		})
	}
	if pu.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   portal.MembersTable,
			Columns: portal.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint32,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedMembersIDs(); len(nodes) > 0 && !pu.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   portal.MembersTable,
			Columns: portal.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint32,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   portal.MembersTable,
			Columns: portal.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint32,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   portal.MetadataTable,
			Columns: []string{portal.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint32,
					Column: portalmetadata.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedMetadataIDs(); len(nodes) > 0 && !pu.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   portal.MetadataTable,
			Columns: []string{portal.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint32,
					Column: portalmetadata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   portal.MetadataTable,
			Columns: []string{portal.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint32,
					Column: portalmetadata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.LegalCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   portal.LegalTable,
			Columns: []string{portal.LegalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint32,
					Column: portallegal.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedLegalIDs(); len(nodes) > 0 && !pu.mutation.LegalCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   portal.LegalTable,
			Columns: []string{portal.LegalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint32,
					Column: portallegal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.LegalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   portal.LegalTable,
			Columns: []string{portal.LegalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint32,
					Column: portallegal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{portal.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PortalUpdateOne is the builder for updating a single Portal entity.
type PortalUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PortalMutation
}

// SetIsActive sets the "isActive" field.
func (puo *PortalUpdateOne) SetIsActive(b bool) *PortalUpdateOne {
	puo.mutation.SetIsActive(b)
	return puo
}

// SetNillableIsActive sets the "isActive" field if the given value is not nil.
func (puo *PortalUpdateOne) SetNillableIsActive(b *bool) *PortalUpdateOne {
	if b != nil {
		puo.SetIsActive(*b)
	}
	return puo
}

// AddMemberIDs adds the "members" edge to the Account entity by IDs.
func (puo *PortalUpdateOne) AddMemberIDs(ids ...uint32) *PortalUpdateOne {
	puo.mutation.AddMemberIDs(ids...)
	return puo
}

// AddMembers adds the "members" edges to the Account entity.
func (puo *PortalUpdateOne) AddMembers(a ...*Account) *PortalUpdateOne {
	ids := make([]uint32, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return puo.AddMemberIDs(ids...)
}

// AddMetadatumIDs adds the "metadata" edge to the PortalMetadata entity by IDs.
func (puo *PortalUpdateOne) AddMetadatumIDs(ids ...uint32) *PortalUpdateOne {
	puo.mutation.AddMetadatumIDs(ids...)
	return puo
}

// AddMetadata adds the "metadata" edges to the PortalMetadata entity.
func (puo *PortalUpdateOne) AddMetadata(p ...*PortalMetadata) *PortalUpdateOne {
	ids := make([]uint32, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddMetadatumIDs(ids...)
}

// AddLegalIDs adds the "legal" edge to the PortalLegal entity by IDs.
func (puo *PortalUpdateOne) AddLegalIDs(ids ...uint32) *PortalUpdateOne {
	puo.mutation.AddLegalIDs(ids...)
	return puo
}

// AddLegal adds the "legal" edges to the PortalLegal entity.
func (puo *PortalUpdateOne) AddLegal(p ...*PortalLegal) *PortalUpdateOne {
	ids := make([]uint32, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddLegalIDs(ids...)
}

// Mutation returns the PortalMutation object of the builder.
func (puo *PortalUpdateOne) Mutation() *PortalMutation {
	return puo.mutation
}

// ClearMembers clears all "members" edges to the Account entity.
func (puo *PortalUpdateOne) ClearMembers() *PortalUpdateOne {
	puo.mutation.ClearMembers()
	return puo
}

// RemoveMemberIDs removes the "members" edge to Account entities by IDs.
func (puo *PortalUpdateOne) RemoveMemberIDs(ids ...uint32) *PortalUpdateOne {
	puo.mutation.RemoveMemberIDs(ids...)
	return puo
}

// RemoveMembers removes "members" edges to Account entities.
func (puo *PortalUpdateOne) RemoveMembers(a ...*Account) *PortalUpdateOne {
	ids := make([]uint32, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return puo.RemoveMemberIDs(ids...)
}

// ClearMetadata clears all "metadata" edges to the PortalMetadata entity.
func (puo *PortalUpdateOne) ClearMetadata() *PortalUpdateOne {
	puo.mutation.ClearMetadata()
	return puo
}

// RemoveMetadatumIDs removes the "metadata" edge to PortalMetadata entities by IDs.
func (puo *PortalUpdateOne) RemoveMetadatumIDs(ids ...uint32) *PortalUpdateOne {
	puo.mutation.RemoveMetadatumIDs(ids...)
	return puo
}

// RemoveMetadata removes "metadata" edges to PortalMetadata entities.
func (puo *PortalUpdateOne) RemoveMetadata(p ...*PortalMetadata) *PortalUpdateOne {
	ids := make([]uint32, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveMetadatumIDs(ids...)
}

// ClearLegal clears all "legal" edges to the PortalLegal entity.
func (puo *PortalUpdateOne) ClearLegal() *PortalUpdateOne {
	puo.mutation.ClearLegal()
	return puo
}

// RemoveLegalIDs removes the "legal" edge to PortalLegal entities by IDs.
func (puo *PortalUpdateOne) RemoveLegalIDs(ids ...uint32) *PortalUpdateOne {
	puo.mutation.RemoveLegalIDs(ids...)
	return puo
}

// RemoveLegal removes "legal" edges to PortalLegal entities.
func (puo *PortalUpdateOne) RemoveLegal(p ...*PortalLegal) *PortalUpdateOne {
	ids := make([]uint32, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveLegalIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PortalUpdateOne) Select(field string, fields ...string) *PortalUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Portal entity.
func (puo *PortalUpdateOne) Save(ctx context.Context) (*Portal, error) {
	var (
		err  error
		node *Portal
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PortalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("store: uninitialized hook (forgotten import store/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, puo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (puo *PortalUpdateOne) SaveX(ctx context.Context) *Portal {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PortalUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PortalUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PortalUpdateOne) sqlSave(ctx context.Context) (_node *Portal, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   portal.Table,
			Columns: portal.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: portal.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`store: missing "Portal.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, portal.FieldID)
		for _, f := range fields {
			if !portal.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("store: invalid field %q for query", f)}
			}
			if f != portal.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.IsActive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: portal.FieldIsActive,
		})
	}
	if puo.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   portal.MembersTable,
			Columns: portal.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint32,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedMembersIDs(); len(nodes) > 0 && !puo.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   portal.MembersTable,
			Columns: portal.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint32,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   portal.MembersTable,
			Columns: portal.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint32,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   portal.MetadataTable,
			Columns: []string{portal.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint32,
					Column: portalmetadata.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedMetadataIDs(); len(nodes) > 0 && !puo.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   portal.MetadataTable,
			Columns: []string{portal.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint32,
					Column: portalmetadata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   portal.MetadataTable,
			Columns: []string{portal.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint32,
					Column: portalmetadata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.LegalCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   portal.LegalTable,
			Columns: []string{portal.LegalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint32,
					Column: portallegal.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedLegalIDs(); len(nodes) > 0 && !puo.mutation.LegalCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   portal.LegalTable,
			Columns: []string{portal.LegalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint32,
					Column: portallegal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.LegalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   portal.LegalTable,
			Columns: []string{portal.LegalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint32,
					Column: portallegal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Portal{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{portal.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
