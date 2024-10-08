// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"wishlist-wrangler-api/ent/predicate"
	"wishlist-wrangler-api/ent/wishlist"

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

// SetTemplateID sets the "template_id" field.
func (wu *WishlistUpdate) SetTemplateID(u uuid.UUID) *WishlistUpdate {
	wu.mutation.SetTemplateID(u)
	return wu
}

// SetNillableTemplateID sets the "template_id" field if the given value is not nil.
func (wu *WishlistUpdate) SetNillableTemplateID(u *uuid.UUID) *WishlistUpdate {
	if u != nil {
		wu.SetTemplateID(*u)
	}
	return wu
}

// SetCreatorID sets the "creator_id" field.
func (wu *WishlistUpdate) SetCreatorID(u uuid.UUID) *WishlistUpdate {
	wu.mutation.SetCreatorID(u)
	return wu
}

// SetNillableCreatorID sets the "creator_id" field if the given value is not nil.
func (wu *WishlistUpdate) SetNillableCreatorID(u *uuid.UUID) *WishlistUpdate {
	if u != nil {
		wu.SetCreatorID(*u)
	}
	return wu
}

// Mutation returns the WishlistMutation object of the builder.
func (wu *WishlistUpdate) Mutation() *WishlistMutation {
	return wu.mutation
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
	if value, ok := wu.mutation.TemplateID(); ok {
		_spec.SetField(wishlist.FieldTemplateID, field.TypeUUID, value)
	}
	if value, ok := wu.mutation.CreatorID(); ok {
		_spec.SetField(wishlist.FieldCreatorID, field.TypeUUID, value)
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

// SetTemplateID sets the "template_id" field.
func (wuo *WishlistUpdateOne) SetTemplateID(u uuid.UUID) *WishlistUpdateOne {
	wuo.mutation.SetTemplateID(u)
	return wuo
}

// SetNillableTemplateID sets the "template_id" field if the given value is not nil.
func (wuo *WishlistUpdateOne) SetNillableTemplateID(u *uuid.UUID) *WishlistUpdateOne {
	if u != nil {
		wuo.SetTemplateID(*u)
	}
	return wuo
}

// SetCreatorID sets the "creator_id" field.
func (wuo *WishlistUpdateOne) SetCreatorID(u uuid.UUID) *WishlistUpdateOne {
	wuo.mutation.SetCreatorID(u)
	return wuo
}

// SetNillableCreatorID sets the "creator_id" field if the given value is not nil.
func (wuo *WishlistUpdateOne) SetNillableCreatorID(u *uuid.UUID) *WishlistUpdateOne {
	if u != nil {
		wuo.SetCreatorID(*u)
	}
	return wuo
}

// Mutation returns the WishlistMutation object of the builder.
func (wuo *WishlistUpdateOne) Mutation() *WishlistMutation {
	return wuo.mutation
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
	if value, ok := wuo.mutation.TemplateID(); ok {
		_spec.SetField(wishlist.FieldTemplateID, field.TypeUUID, value)
	}
	if value, ok := wuo.mutation.CreatorID(); ok {
		_spec.SetField(wishlist.FieldCreatorID, field.TypeUUID, value)
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
