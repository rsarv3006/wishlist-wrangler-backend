// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"wishlist-wrangler-api/ent/predicate"
	"wishlist-wrangler-api/ent/user"
	"wishlist-wrangler-api/ent/wishlist"
	"wishlist-wrangler-api/ent/wishlistsection"
	"wishlist-wrangler-api/ent/wishlisttemplate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WishlistUpdate is the builder for updating Wishlist entities.
type WishlistUpdate struct {
	config
	hooks    []Hook
	mutation *WishlistMutation
}

// Where appends a list predicates to the WishlistUpdate builder.
func (wu *WishlistUpdate) Where(ps ...predicate.Wishlist) *WishlistUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetTitle sets the "title" field.
func (wu *WishlistUpdate) SetTitle(s string) *WishlistUpdate {
	wu.mutation.SetTitle(s)
	return wu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (wu *WishlistUpdate) SetNillableTitle(s *string) *WishlistUpdate {
	if s != nil {
		wu.SetTitle(*s)
	}
	return wu
}

// SetStatus sets the "status" field.
func (wu *WishlistUpdate) SetStatus(w wishlist.Status) *WishlistUpdate {
	wu.mutation.SetStatus(w)
	return wu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wu *WishlistUpdate) SetNillableStatus(w *wishlist.Status) *WishlistUpdate {
	if w != nil {
		wu.SetStatus(*w)
	}
	return wu
}

// AddCreatorIDs adds the "creator" edge to the User entity by IDs.
func (wu *WishlistUpdate) AddCreatorIDs(ids ...uuid.UUID) *WishlistUpdate {
	wu.mutation.AddCreatorIDs(ids...)
	return wu
}

// AddCreator adds the "creator" edges to the User entity.
func (wu *WishlistUpdate) AddCreator(u ...*User) *WishlistUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return wu.AddCreatorIDs(ids...)
}

// AddTemplateIDs adds the "template" edge to the WishlistTemplate entity by IDs.
func (wu *WishlistUpdate) AddTemplateIDs(ids ...uuid.UUID) *WishlistUpdate {
	wu.mutation.AddTemplateIDs(ids...)
	return wu
}

// AddTemplate adds the "template" edges to the WishlistTemplate entity.
func (wu *WishlistUpdate) AddTemplate(w ...*WishlistTemplate) *WishlistUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.AddTemplateIDs(ids...)
}

// AddSectionIDs adds the "sections" edge to the WishlistSection entity by IDs.
func (wu *WishlistUpdate) AddSectionIDs(ids ...uuid.UUID) *WishlistUpdate {
	wu.mutation.AddSectionIDs(ids...)
	return wu
}

// AddSections adds the "sections" edges to the WishlistSection entity.
func (wu *WishlistUpdate) AddSections(w ...*WishlistSection) *WishlistUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.AddSectionIDs(ids...)
}

// Mutation returns the WishlistMutation object of the builder.
func (wu *WishlistUpdate) Mutation() *WishlistMutation {
	return wu.mutation
}

// ClearCreator clears all "creator" edges to the User entity.
func (wu *WishlistUpdate) ClearCreator() *WishlistUpdate {
	wu.mutation.ClearCreator()
	return wu
}

// RemoveCreatorIDs removes the "creator" edge to User entities by IDs.
func (wu *WishlistUpdate) RemoveCreatorIDs(ids ...uuid.UUID) *WishlistUpdate {
	wu.mutation.RemoveCreatorIDs(ids...)
	return wu
}

// RemoveCreator removes "creator" edges to User entities.
func (wu *WishlistUpdate) RemoveCreator(u ...*User) *WishlistUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return wu.RemoveCreatorIDs(ids...)
}

// ClearTemplate clears all "template" edges to the WishlistTemplate entity.
func (wu *WishlistUpdate) ClearTemplate() *WishlistUpdate {
	wu.mutation.ClearTemplate()
	return wu
}

// RemoveTemplateIDs removes the "template" edge to WishlistTemplate entities by IDs.
func (wu *WishlistUpdate) RemoveTemplateIDs(ids ...uuid.UUID) *WishlistUpdate {
	wu.mutation.RemoveTemplateIDs(ids...)
	return wu
}

// RemoveTemplate removes "template" edges to WishlistTemplate entities.
func (wu *WishlistUpdate) RemoveTemplate(w ...*WishlistTemplate) *WishlistUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.RemoveTemplateIDs(ids...)
}

// ClearSections clears all "sections" edges to the WishlistSection entity.
func (wu *WishlistUpdate) ClearSections() *WishlistUpdate {
	wu.mutation.ClearSections()
	return wu
}

// RemoveSectionIDs removes the "sections" edge to WishlistSection entities by IDs.
func (wu *WishlistUpdate) RemoveSectionIDs(ids ...uuid.UUID) *WishlistUpdate {
	wu.mutation.RemoveSectionIDs(ids...)
	return wu
}

// RemoveSections removes "sections" edges to WishlistSection entities.
func (wu *WishlistUpdate) RemoveSections(w ...*WishlistSection) *WishlistUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.RemoveSectionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WishlistUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, wu.sqlSave, wu.mutation, wu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WishlistUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WishlistUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WishlistUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WishlistUpdate) check() error {
	if v, ok := wu.mutation.Title(); ok {
		if err := wishlist.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Wishlist.title": %w`, err)}
		}
	}
	if v, ok := wu.mutation.Status(); ok {
		if err := wishlist.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Wishlist.status": %w`, err)}
		}
	}
	return nil
}

func (wu *WishlistUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(wishlist.Table, wishlist.Columns, sqlgraph.NewFieldSpec(wishlist.FieldID, field.TypeUUID))
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.Title(); ok {
		_spec.SetField(wishlist.FieldTitle, field.TypeString, value)
	}
	if value, ok := wu.mutation.Status(); ok {
		_spec.SetField(wishlist.FieldStatus, field.TypeEnum, value)
	}
	if wu.mutation.CreatorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlist.CreatorTable,
			Columns: []string{wishlist.CreatorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedCreatorIDs(); len(nodes) > 0 && !wu.mutation.CreatorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlist.CreatorTable,
			Columns: []string{wishlist.CreatorColumn},
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
	if nodes := wu.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlist.CreatorTable,
			Columns: []string{wishlist.CreatorColumn},
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
	if wu.mutation.TemplateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlist.TemplateTable,
			Columns: []string{wishlist.TemplateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wishlisttemplate.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedTemplateIDs(); len(nodes) > 0 && !wu.mutation.TemplateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlist.TemplateTable,
			Columns: []string{wishlist.TemplateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wishlisttemplate.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.TemplateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlist.TemplateTable,
			Columns: []string{wishlist.TemplateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wishlisttemplate.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wu.mutation.SectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlist.SectionsTable,
			Columns: []string{wishlist.SectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wishlistsection.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedSectionsIDs(); len(nodes) > 0 && !wu.mutation.SectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlist.SectionsTable,
			Columns: []string{wishlist.SectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wishlistsection.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.SectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlist.SectionsTable,
			Columns: []string{wishlist.SectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wishlistsection.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wishlist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wu.mutation.done = true
	return n, nil
}

// WishlistUpdateOne is the builder for updating a single Wishlist entity.
type WishlistUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WishlistMutation
}

// SetTitle sets the "title" field.
func (wuo *WishlistUpdateOne) SetTitle(s string) *WishlistUpdateOne {
	wuo.mutation.SetTitle(s)
	return wuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (wuo *WishlistUpdateOne) SetNillableTitle(s *string) *WishlistUpdateOne {
	if s != nil {
		wuo.SetTitle(*s)
	}
	return wuo
}

// SetStatus sets the "status" field.
func (wuo *WishlistUpdateOne) SetStatus(w wishlist.Status) *WishlistUpdateOne {
	wuo.mutation.SetStatus(w)
	return wuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wuo *WishlistUpdateOne) SetNillableStatus(w *wishlist.Status) *WishlistUpdateOne {
	if w != nil {
		wuo.SetStatus(*w)
	}
	return wuo
}

// AddCreatorIDs adds the "creator" edge to the User entity by IDs.
func (wuo *WishlistUpdateOne) AddCreatorIDs(ids ...uuid.UUID) *WishlistUpdateOne {
	wuo.mutation.AddCreatorIDs(ids...)
	return wuo
}

// AddCreator adds the "creator" edges to the User entity.
func (wuo *WishlistUpdateOne) AddCreator(u ...*User) *WishlistUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return wuo.AddCreatorIDs(ids...)
}

// AddTemplateIDs adds the "template" edge to the WishlistTemplate entity by IDs.
func (wuo *WishlistUpdateOne) AddTemplateIDs(ids ...uuid.UUID) *WishlistUpdateOne {
	wuo.mutation.AddTemplateIDs(ids...)
	return wuo
}

// AddTemplate adds the "template" edges to the WishlistTemplate entity.
func (wuo *WishlistUpdateOne) AddTemplate(w ...*WishlistTemplate) *WishlistUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.AddTemplateIDs(ids...)
}

// AddSectionIDs adds the "sections" edge to the WishlistSection entity by IDs.
func (wuo *WishlistUpdateOne) AddSectionIDs(ids ...uuid.UUID) *WishlistUpdateOne {
	wuo.mutation.AddSectionIDs(ids...)
	return wuo
}

// AddSections adds the "sections" edges to the WishlistSection entity.
func (wuo *WishlistUpdateOne) AddSections(w ...*WishlistSection) *WishlistUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.AddSectionIDs(ids...)
}

// Mutation returns the WishlistMutation object of the builder.
func (wuo *WishlistUpdateOne) Mutation() *WishlistMutation {
	return wuo.mutation
}

// ClearCreator clears all "creator" edges to the User entity.
func (wuo *WishlistUpdateOne) ClearCreator() *WishlistUpdateOne {
	wuo.mutation.ClearCreator()
	return wuo
}

// RemoveCreatorIDs removes the "creator" edge to User entities by IDs.
func (wuo *WishlistUpdateOne) RemoveCreatorIDs(ids ...uuid.UUID) *WishlistUpdateOne {
	wuo.mutation.RemoveCreatorIDs(ids...)
	return wuo
}

// RemoveCreator removes "creator" edges to User entities.
func (wuo *WishlistUpdateOne) RemoveCreator(u ...*User) *WishlistUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return wuo.RemoveCreatorIDs(ids...)
}

// ClearTemplate clears all "template" edges to the WishlistTemplate entity.
func (wuo *WishlistUpdateOne) ClearTemplate() *WishlistUpdateOne {
	wuo.mutation.ClearTemplate()
	return wuo
}

// RemoveTemplateIDs removes the "template" edge to WishlistTemplate entities by IDs.
func (wuo *WishlistUpdateOne) RemoveTemplateIDs(ids ...uuid.UUID) *WishlistUpdateOne {
	wuo.mutation.RemoveTemplateIDs(ids...)
	return wuo
}

// RemoveTemplate removes "template" edges to WishlistTemplate entities.
func (wuo *WishlistUpdateOne) RemoveTemplate(w ...*WishlistTemplate) *WishlistUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.RemoveTemplateIDs(ids...)
}

// ClearSections clears all "sections" edges to the WishlistSection entity.
func (wuo *WishlistUpdateOne) ClearSections() *WishlistUpdateOne {
	wuo.mutation.ClearSections()
	return wuo
}

// RemoveSectionIDs removes the "sections" edge to WishlistSection entities by IDs.
func (wuo *WishlistUpdateOne) RemoveSectionIDs(ids ...uuid.UUID) *WishlistUpdateOne {
	wuo.mutation.RemoveSectionIDs(ids...)
	return wuo
}

// RemoveSections removes "sections" edges to WishlistSection entities.
func (wuo *WishlistUpdateOne) RemoveSections(w ...*WishlistSection) *WishlistUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.RemoveSectionIDs(ids...)
}

// Where appends a list predicates to the WishlistUpdate builder.
func (wuo *WishlistUpdateOne) Where(ps ...predicate.Wishlist) *WishlistUpdateOne {
	wuo.mutation.Where(ps...)
	return wuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WishlistUpdateOne) Select(field string, fields ...string) *WishlistUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Wishlist entity.
func (wuo *WishlistUpdateOne) Save(ctx context.Context) (*Wishlist, error) {
	return withHooks(ctx, wuo.sqlSave, wuo.mutation, wuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WishlistUpdateOne) SaveX(ctx context.Context) *Wishlist {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WishlistUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WishlistUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WishlistUpdateOne) check() error {
	if v, ok := wuo.mutation.Title(); ok {
		if err := wishlist.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Wishlist.title": %w`, err)}
		}
	}
	if v, ok := wuo.mutation.Status(); ok {
		if err := wishlist.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Wishlist.status": %w`, err)}
		}
	}
	return nil
}

func (wuo *WishlistUpdateOne) sqlSave(ctx context.Context) (_node *Wishlist, err error) {
	if err := wuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(wishlist.Table, wishlist.Columns, sqlgraph.NewFieldSpec(wishlist.FieldID, field.TypeUUID))
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Wishlist.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wishlist.FieldID)
		for _, f := range fields {
			if !wishlist.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != wishlist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.Title(); ok {
		_spec.SetField(wishlist.FieldTitle, field.TypeString, value)
	}
	if value, ok := wuo.mutation.Status(); ok {
		_spec.SetField(wishlist.FieldStatus, field.TypeEnum, value)
	}
	if wuo.mutation.CreatorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlist.CreatorTable,
			Columns: []string{wishlist.CreatorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedCreatorIDs(); len(nodes) > 0 && !wuo.mutation.CreatorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlist.CreatorTable,
			Columns: []string{wishlist.CreatorColumn},
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
	if nodes := wuo.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlist.CreatorTable,
			Columns: []string{wishlist.CreatorColumn},
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
	if wuo.mutation.TemplateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlist.TemplateTable,
			Columns: []string{wishlist.TemplateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wishlisttemplate.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedTemplateIDs(); len(nodes) > 0 && !wuo.mutation.TemplateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlist.TemplateTable,
			Columns: []string{wishlist.TemplateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wishlisttemplate.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.TemplateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlist.TemplateTable,
			Columns: []string{wishlist.TemplateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wishlisttemplate.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wuo.mutation.SectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlist.SectionsTable,
			Columns: []string{wishlist.SectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wishlistsection.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedSectionsIDs(); len(nodes) > 0 && !wuo.mutation.SectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlist.SectionsTable,
			Columns: []string{wishlist.SectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wishlistsection.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.SectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wishlist.SectionsTable,
			Columns: []string{wishlist.SectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wishlistsection.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Wishlist{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wishlist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wuo.mutation.done = true
	return _node, nil
}
