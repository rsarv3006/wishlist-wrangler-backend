// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"wishlist-wrangler-api/ent/predicate"
	"wishlist-wrangler-api/ent/wishlisttemplate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WishlistTemplateDelete is the builder for deleting a WishlistTemplate entity.
type WishlistTemplateDelete struct {
	config
	hooks    []Hook
	mutation *WishlistTemplateMutation
}

// Where appends a list predicates to the WishlistTemplateDelete builder.
func (wtd *WishlistTemplateDelete) Where(ps ...predicate.WishlistTemplate) *WishlistTemplateDelete {
	wtd.mutation.Where(ps...)
	return wtd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wtd *WishlistTemplateDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wtd.sqlExec, wtd.mutation, wtd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wtd *WishlistTemplateDelete) ExecX(ctx context.Context) int {
	n, err := wtd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wtd *WishlistTemplateDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(wishlisttemplate.Table, sqlgraph.NewFieldSpec(wishlisttemplate.FieldID, field.TypeUUID))
	if ps := wtd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wtd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wtd.mutation.done = true
	return affected, err
}

// WishlistTemplateDeleteOne is the builder for deleting a single WishlistTemplate entity.
type WishlistTemplateDeleteOne struct {
	wtd *WishlistTemplateDelete
}

// Where appends a list predicates to the WishlistTemplateDelete builder.
func (wtdo *WishlistTemplateDeleteOne) Where(ps ...predicate.WishlistTemplate) *WishlistTemplateDeleteOne {
	wtdo.wtd.mutation.Where(ps...)
	return wtdo
}

// Exec executes the deletion query.
func (wtdo *WishlistTemplateDeleteOne) Exec(ctx context.Context) error {
	n, err := wtdo.wtd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{wishlisttemplate.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wtdo *WishlistTemplateDeleteOne) ExecX(ctx context.Context) {
	if err := wtdo.Exec(ctx); err != nil {
		panic(err)
	}
}
