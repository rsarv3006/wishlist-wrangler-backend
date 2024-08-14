// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"wishlist-wrangler-api/ent/predicate"
	"wishlist-wrangler-api/ent/user"
	"wishlist-wrangler-api/ent/wishlisttemplate"
	"wishlist-wrangler-api/ent/wishlisttemplatesection"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WishlistTemplateUpdate is the builder for updating WishlistTemplate entities.
type WishlistTemplateUpdate struct {
	config
	hooks    []Hook
	mutation *WishlistTemplateMutation
}

// Where appends a list predicates to the WishlistTemplateUpdate builder.
func (wtu *WishlistTemplateUpdate) Where(ps ...predicate.WishlistTemplate) *WishlistTemplateUpdate {
	wtu.mutation.Where(ps...)
	return wtu
}

// SetTitle sets the "title" field.
func (wtu *WishlistTemplateUpdate) SetTitle(s string) *WishlistTemplateUpdate {
	wtu.mutation.SetTitle(s)
	return wtu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (wtu *WishlistTemplateUpdate) SetNillableTitle(s *string) *WishlistTemplateUpdate {
	if s != nil {
		wtu.SetTitle(*s)
	}
	return wtu
}

// SetDescription sets the "description" field.
func (wtu *WishlistTemplateUpdate) SetDescription(s string) *WishlistTemplateUpdate {
	wtu.mutation.SetDescription(s)
	return wtu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wtu *WishlistTemplateUpdate) SetNillableDescription(s *string) *WishlistTemplateUpdate {
	if s != nil {
		wtu.SetDescription(*s)
	}
	return wtu
}

// AddCreatorIdIDs adds the "creatorId" edge to the User entity by IDs.
func (wtu *WishlistTemplateUpdate) AddCreatorIdIDs(ids ...uuid.UUID) *WishlistTemplateUpdate {
	wtu.mutation.AddCreatorIdIDs(ids...)
	return wtu
}

// AddCreatorId adds the "creatorId" edges to the User entity.
func (wtu *WishlistTemplateUpdate) AddCreatorId(u ...*User) *WishlistTemplateUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return wtu.AddCreatorIdIDs(ids...)
}

// AddSectionIDs adds the "sections" edge to the WishlistTemplateSection entity by IDs.
func (wtu *WishlistTemplateUpdate) AddSectionIDs(ids ...uuid.UUID) *WishlistTemplateUpdate {
	wtu.mutation.AddSectionIDs(ids...)
	return wtu
}

// AddSections adds the "sections" edges to the WishlistTemplateSection entity.
func (wtu *WishlistTemplateUpdate) AddSections(w ...*WishlistTemplateSection) *WishlistTemplateUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wtu.AddSectionIDs(ids...)
}

// Mutation returns the WishlistTemplateMutation object of the builder.
func (wtu *WishlistTemplateUpdate) Mutation() *WishlistTemplateMutation {
	return wtu.mutation
}

// ClearCreatorId clears all "creatorId" edges to the User entity.
func (wtu *WishlistTemplateUpdate) ClearCreatorId() *WishlistTemplateUpdate {
	wtu.mutation.ClearCreatorId()
	return wtu
}

// RemoveCreatorIdIDs removes the "creatorId" edge to User entities by IDs.
func (wtu *WishlistTemplateUpdate) RemoveCreatorIdIDs(ids ...uuid.UUID) *WishlistTemplateUpdate {
	wtu.mutation.RemoveCreatorIdIDs(ids...)
	return wtu
}

// RemoveCreatorId removes "creatorId" edges to User entities.
func (wtu *WishlistTemplateUpdate) RemoveCreatorId(u ...*User) *WishlistTemplateUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return wtu.RemoveCreatorIdIDs(ids...)
}

// ClearSections clears all "sections" edges to the WishlistTemplateSection entity.
func (wtu *WishlistTemplateUpdate) ClearSections() *WishlistTemplateUpdate {
	wtu.mutation.ClearSections()
	return wtu
}

// RemoveSectionIDs removes the "sections" edge to WishlistTemplateSection entities by IDs.
func (wtu *WishlistTemplateUpdate) RemoveSectionIDs(ids ...uuid.UUID) *WishlistTemplateUpdate {
	wtu.mutation.RemoveSectionIDs(ids...)
	return wtu
}

// RemoveSections removes "sections" edges to WishlistTemplateSection entities.
func (wtu *WishlistTemplateUpdate) RemoveSections(w ...*WishlistTemplateSection) *WishlistTemplateUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wtu.RemoveSectionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wtu *WishlistTemplateUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, wtu.sqlSave, wtu.mutation, wtu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wtu *WishlistTemplateUpdate) SaveX(ctx context.Context) int {
	affected, err := wtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wtu *WishlistTemplateUpdate) Exec(ctx context.Context) error {
	_, err := wtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wtu *WishlistTemplateUpdate) ExecX(ctx context.Context) {
	if err := wtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wtu *WishlistTemplateUpdate) check() error {
	if v, ok := wtu.mutation.Title(); ok {
		if err := wishlisttemplate.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "WishlistTemplate.title": %w`, err)}
		}
	}
	if v, ok := wtu.mutation.Description(); ok {
		if err := wishlisttemplate.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "WishlistTemplate.description": %w`, err)}
		}
	}
	return nil
}

func (wtu *WishlistTemplateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wtu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(wishlisttemplate.Table, wishlisttemplate.Columns, sqlgraph.NewFieldSpec(wishlisttemplate.FieldID, field.TypeUUID))
	if ps := wtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wtu.mutation.Title(); ok {
		_spec.SetField(wishlisttemplate.FieldTitle, field.TypeString, value)
	}
	if value, ok := wtu.mutation.Description(); ok {
		_spec.SetField(wishlisttemplate.FieldDescription, field.TypeString, value)
	}
	if wtu.mutation.CreatorIdCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlisttemplate.CreatorIdTable,
			Columns: []string{wishlisttemplate.CreatorIdColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wtu.mutation.RemovedCreatorIdIDs(); len(nodes) > 0 && !wtu.mutation.CreatorIdCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlisttemplate.CreatorIdTable,
			Columns: []string{wishlisttemplate.CreatorIdColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wtu.mutation.CreatorIdIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlisttemplate.CreatorIdTable,
			Columns: []string{wishlisttemplate.CreatorIdColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wtu.mutation.SectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlisttemplate.SectionsTable,
			Columns: []string{wishlisttemplate.SectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wishlisttemplatesection.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wtu.mutation.RemovedSectionsIDs(); len(nodes) > 0 && !wtu.mutation.SectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlisttemplate.SectionsTable,
			Columns: []string{wishlisttemplate.SectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wishlisttemplatesection.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wtu.mutation.SectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlisttemplate.SectionsTable,
			Columns: []string{wishlisttemplate.SectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wishlisttemplatesection.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wishlisttemplate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wtu.mutation.done = true
	return n, nil
}

// WishlistTemplateUpdateOne is the builder for updating a single WishlistTemplate entity.
type WishlistTemplateUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WishlistTemplateMutation
}

// SetTitle sets the "title" field.
func (wtuo *WishlistTemplateUpdateOne) SetTitle(s string) *WishlistTemplateUpdateOne {
	wtuo.mutation.SetTitle(s)
	return wtuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (wtuo *WishlistTemplateUpdateOne) SetNillableTitle(s *string) *WishlistTemplateUpdateOne {
	if s != nil {
		wtuo.SetTitle(*s)
	}
	return wtuo
}

// SetDescription sets the "description" field.
func (wtuo *WishlistTemplateUpdateOne) SetDescription(s string) *WishlistTemplateUpdateOne {
	wtuo.mutation.SetDescription(s)
	return wtuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wtuo *WishlistTemplateUpdateOne) SetNillableDescription(s *string) *WishlistTemplateUpdateOne {
	if s != nil {
		wtuo.SetDescription(*s)
	}
	return wtuo
}

// AddCreatorIdIDs adds the "creatorId" edge to the User entity by IDs.
func (wtuo *WishlistTemplateUpdateOne) AddCreatorIdIDs(ids ...uuid.UUID) *WishlistTemplateUpdateOne {
	wtuo.mutation.AddCreatorIdIDs(ids...)
	return wtuo
}

// AddCreatorId adds the "creatorId" edges to the User entity.
func (wtuo *WishlistTemplateUpdateOne) AddCreatorId(u ...*User) *WishlistTemplateUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return wtuo.AddCreatorIdIDs(ids...)
}

// AddSectionIDs adds the "sections" edge to the WishlistTemplateSection entity by IDs.
func (wtuo *WishlistTemplateUpdateOne) AddSectionIDs(ids ...uuid.UUID) *WishlistTemplateUpdateOne {
	wtuo.mutation.AddSectionIDs(ids...)
	return wtuo
}

// AddSections adds the "sections" edges to the WishlistTemplateSection entity.
func (wtuo *WishlistTemplateUpdateOne) AddSections(w ...*WishlistTemplateSection) *WishlistTemplateUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wtuo.AddSectionIDs(ids...)
}

// Mutation returns the WishlistTemplateMutation object of the builder.
func (wtuo *WishlistTemplateUpdateOne) Mutation() *WishlistTemplateMutation {
	return wtuo.mutation
}

// ClearCreatorId clears all "creatorId" edges to the User entity.
func (wtuo *WishlistTemplateUpdateOne) ClearCreatorId() *WishlistTemplateUpdateOne {
	wtuo.mutation.ClearCreatorId()
	return wtuo
}

// RemoveCreatorIdIDs removes the "creatorId" edge to User entities by IDs.
func (wtuo *WishlistTemplateUpdateOne) RemoveCreatorIdIDs(ids ...uuid.UUID) *WishlistTemplateUpdateOne {
	wtuo.mutation.RemoveCreatorIdIDs(ids...)
	return wtuo
}

// RemoveCreatorId removes "creatorId" edges to User entities.
func (wtuo *WishlistTemplateUpdateOne) RemoveCreatorId(u ...*User) *WishlistTemplateUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return wtuo.RemoveCreatorIdIDs(ids...)
}

// ClearSections clears all "sections" edges to the WishlistTemplateSection entity.
func (wtuo *WishlistTemplateUpdateOne) ClearSections() *WishlistTemplateUpdateOne {
	wtuo.mutation.ClearSections()
	return wtuo
}

// RemoveSectionIDs removes the "sections" edge to WishlistTemplateSection entities by IDs.
func (wtuo *WishlistTemplateUpdateOne) RemoveSectionIDs(ids ...uuid.UUID) *WishlistTemplateUpdateOne {
	wtuo.mutation.RemoveSectionIDs(ids...)
	return wtuo
}

// RemoveSections removes "sections" edges to WishlistTemplateSection entities.
func (wtuo *WishlistTemplateUpdateOne) RemoveSections(w ...*WishlistTemplateSection) *WishlistTemplateUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wtuo.RemoveSectionIDs(ids...)
}

// Where appends a list predicates to the WishlistTemplateUpdate builder.
func (wtuo *WishlistTemplateUpdateOne) Where(ps ...predicate.WishlistTemplate) *WishlistTemplateUpdateOne {
	wtuo.mutation.Where(ps...)
	return wtuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wtuo *WishlistTemplateUpdateOne) Select(field string, fields ...string) *WishlistTemplateUpdateOne {
	wtuo.fields = append([]string{field}, fields...)
	return wtuo
}

// Save executes the query and returns the updated WishlistTemplate entity.
func (wtuo *WishlistTemplateUpdateOne) Save(ctx context.Context) (*WishlistTemplate, error) {
	return withHooks(ctx, wtuo.sqlSave, wtuo.mutation, wtuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wtuo *WishlistTemplateUpdateOne) SaveX(ctx context.Context) *WishlistTemplate {
	node, err := wtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wtuo *WishlistTemplateUpdateOne) Exec(ctx context.Context) error {
	_, err := wtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wtuo *WishlistTemplateUpdateOne) ExecX(ctx context.Context) {
	if err := wtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wtuo *WishlistTemplateUpdateOne) check() error {
	if v, ok := wtuo.mutation.Title(); ok {
		if err := wishlisttemplate.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "WishlistTemplate.title": %w`, err)}
		}
	}
	if v, ok := wtuo.mutation.Description(); ok {
		if err := wishlisttemplate.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "WishlistTemplate.description": %w`, err)}
		}
	}
	return nil
}

func (wtuo *WishlistTemplateUpdateOne) sqlSave(ctx context.Context) (_node *WishlistTemplate, err error) {
	if err := wtuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(wishlisttemplate.Table, wishlisttemplate.Columns, sqlgraph.NewFieldSpec(wishlisttemplate.FieldID, field.TypeUUID))
	id, ok := wtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WishlistTemplate.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wishlisttemplate.FieldID)
		for _, f := range fields {
			if !wishlisttemplate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != wishlisttemplate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wtuo.mutation.Title(); ok {
		_spec.SetField(wishlisttemplate.FieldTitle, field.TypeString, value)
	}
	if value, ok := wtuo.mutation.Description(); ok {
		_spec.SetField(wishlisttemplate.FieldDescription, field.TypeString, value)
	}
	if wtuo.mutation.CreatorIdCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlisttemplate.CreatorIdTable,
			Columns: []string{wishlisttemplate.CreatorIdColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wtuo.mutation.RemovedCreatorIdIDs(); len(nodes) > 0 && !wtuo.mutation.CreatorIdCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlisttemplate.CreatorIdTable,
			Columns: []string{wishlisttemplate.CreatorIdColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wtuo.mutation.CreatorIdIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlisttemplate.CreatorIdTable,
			Columns: []string{wishlisttemplate.CreatorIdColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wtuo.mutation.SectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlisttemplate.SectionsTable,
			Columns: []string{wishlisttemplate.SectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wishlisttemplatesection.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wtuo.mutation.RemovedSectionsIDs(); len(nodes) > 0 && !wtuo.mutation.SectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlisttemplate.SectionsTable,
			Columns: []string{wishlisttemplate.SectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wishlisttemplatesection.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wtuo.mutation.SectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlisttemplate.SectionsTable,
			Columns: []string{wishlisttemplate.SectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wishlisttemplatesection.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &WishlistTemplate{config: wtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wishlisttemplate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wtuo.mutation.done = true
	return _node, nil
}
