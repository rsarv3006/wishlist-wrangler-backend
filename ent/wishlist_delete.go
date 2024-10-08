// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"wishlist-wrangler-api/ent/predicate"
	"wishlist-wrangler-api/ent/wishlist"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WishlistDelete is the builder for deleting a Wishlist entity.
type WishlistDelete struct {
	config
	hooks    []Hook
	mutation *WishlistMutation
}

// Where appends a list predicates to the WishlistDelete builder.
func (wd *WishlistDelete) Where(ps ...predicate.Wishlist) *WishlistDelete {
	wd.mutation.Where(ps...)
	return wd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wd *WishlistDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wd.sqlExec, wd.mutation, wd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wd *WishlistDelete) ExecX(ctx context.Context) int {
	n, err := wd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wd *WishlistDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(wishlist.Table, sqlgraph.NewFieldSpec(wishlist.FieldID, field.TypeUUID))
	if ps := wd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wd.mutation.done = true
	return affected, err
}

// WishlistDeleteOne is the builder for deleting a single Wishlist entity.
type WishlistDeleteOne struct {
	wd *WishlistDelete
}

// Where appends a list predicates to the WishlistDelete builder.
func (wdo *WishlistDeleteOne) Where(ps ...predicate.Wishlist) *WishlistDeleteOne {
	wdo.wd.mutation.Where(ps...)
	return wdo
}

// Exec executes the deletion query.
func (wdo *WishlistDeleteOne) Exec(ctx context.Context) error {
	n, err := wdo.wd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{wishlist.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wdo *WishlistDeleteOne) ExecX(ctx context.Context) {
	if err := wdo.Exec(ctx); err != nil {
		panic(err)
	}
}
