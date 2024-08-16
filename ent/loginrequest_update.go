// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"wishlist-wrangler-api/ent/loginrequest"
	"wishlist-wrangler-api/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// LoginRequestUpdate is the builder for updating LoginRequest entities.
type LoginRequestUpdate struct {
	config
	hooks    []Hook
	mutation *LoginRequestMutation
}

// Where appends a list predicates to the LoginRequestUpdate builder.
func (lru *LoginRequestUpdate) Where(ps ...predicate.LoginRequest) *LoginRequestUpdate {
	lru.mutation.Where(ps...)
	return lru
}

// SetUserId sets the "userId" field.
func (lru *LoginRequestUpdate) SetUserId(u uuid.UUID) *LoginRequestUpdate {
	lru.mutation.SetUserId(u)
	return lru
}

// SetNillableUserId sets the "userId" field if the given value is not nil.
func (lru *LoginRequestUpdate) SetNillableUserId(u *uuid.UUID) *LoginRequestUpdate {
	if u != nil {
		lru.SetUserId(*u)
	}
	return lru
}

// SetEmail sets the "email" field.
func (lru *LoginRequestUpdate) SetEmail(s string) *LoginRequestUpdate {
	lru.mutation.SetEmail(s)
	return lru
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (lru *LoginRequestUpdate) SetNillableEmail(s *string) *LoginRequestUpdate {
	if s != nil {
		lru.SetEmail(*s)
	}
	return lru
}

// SetLoginRequestCode sets the "loginRequestCode" field.
func (lru *LoginRequestUpdate) SetLoginRequestCode(s string) *LoginRequestUpdate {
	lru.mutation.SetLoginRequestCode(s)
	return lru
}

// SetNillableLoginRequestCode sets the "loginRequestCode" field if the given value is not nil.
func (lru *LoginRequestUpdate) SetNillableLoginRequestCode(s *string) *LoginRequestUpdate {
	if s != nil {
		lru.SetLoginRequestCode(*s)
	}
	return lru
}

// SetStatus sets the "status" field.
func (lru *LoginRequestUpdate) SetStatus(l loginrequest.Status) *LoginRequestUpdate {
	lru.mutation.SetStatus(l)
	return lru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (lru *LoginRequestUpdate) SetNillableStatus(l *loginrequest.Status) *LoginRequestUpdate {
	if l != nil {
		lru.SetStatus(*l)
	}
	return lru
}

// Mutation returns the LoginRequestMutation object of the builder.
func (lru *LoginRequestUpdate) Mutation() *LoginRequestMutation {
	return lru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lru *LoginRequestUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lru.sqlSave, lru.mutation, lru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lru *LoginRequestUpdate) SaveX(ctx context.Context) int {
	affected, err := lru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lru *LoginRequestUpdate) Exec(ctx context.Context) error {
	_, err := lru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lru *LoginRequestUpdate) ExecX(ctx context.Context) {
	if err := lru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lru *LoginRequestUpdate) check() error {
	if v, ok := lru.mutation.Status(); ok {
		if err := loginrequest.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "LoginRequest.status": %w`, err)}
		}
	}
	return nil
}

func (lru *LoginRequestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(loginrequest.Table, loginrequest.Columns, sqlgraph.NewFieldSpec(loginrequest.FieldID, field.TypeUUID))
	if ps := lru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lru.mutation.UserId(); ok {
		_spec.SetField(loginrequest.FieldUserId, field.TypeUUID, value)
	}
	if value, ok := lru.mutation.Email(); ok {
		_spec.SetField(loginrequest.FieldEmail, field.TypeString, value)
	}
	if value, ok := lru.mutation.LoginRequestCode(); ok {
		_spec.SetField(loginrequest.FieldLoginRequestCode, field.TypeString, value)
	}
	if value, ok := lru.mutation.Status(); ok {
		_spec.SetField(loginrequest.FieldStatus, field.TypeEnum, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{loginrequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lru.mutation.done = true
	return n, nil
}

// LoginRequestUpdateOne is the builder for updating a single LoginRequest entity.
type LoginRequestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LoginRequestMutation
}

// SetUserId sets the "userId" field.
func (lruo *LoginRequestUpdateOne) SetUserId(u uuid.UUID) *LoginRequestUpdateOne {
	lruo.mutation.SetUserId(u)
	return lruo
}

// SetNillableUserId sets the "userId" field if the given value is not nil.
func (lruo *LoginRequestUpdateOne) SetNillableUserId(u *uuid.UUID) *LoginRequestUpdateOne {
	if u != nil {
		lruo.SetUserId(*u)
	}
	return lruo
}

// SetEmail sets the "email" field.
func (lruo *LoginRequestUpdateOne) SetEmail(s string) *LoginRequestUpdateOne {
	lruo.mutation.SetEmail(s)
	return lruo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (lruo *LoginRequestUpdateOne) SetNillableEmail(s *string) *LoginRequestUpdateOne {
	if s != nil {
		lruo.SetEmail(*s)
	}
	return lruo
}

// SetLoginRequestCode sets the "loginRequestCode" field.
func (lruo *LoginRequestUpdateOne) SetLoginRequestCode(s string) *LoginRequestUpdateOne {
	lruo.mutation.SetLoginRequestCode(s)
	return lruo
}

// SetNillableLoginRequestCode sets the "loginRequestCode" field if the given value is not nil.
func (lruo *LoginRequestUpdateOne) SetNillableLoginRequestCode(s *string) *LoginRequestUpdateOne {
	if s != nil {
		lruo.SetLoginRequestCode(*s)
	}
	return lruo
}

// SetStatus sets the "status" field.
func (lruo *LoginRequestUpdateOne) SetStatus(l loginrequest.Status) *LoginRequestUpdateOne {
	lruo.mutation.SetStatus(l)
	return lruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (lruo *LoginRequestUpdateOne) SetNillableStatus(l *loginrequest.Status) *LoginRequestUpdateOne {
	if l != nil {
		lruo.SetStatus(*l)
	}
	return lruo
}

// Mutation returns the LoginRequestMutation object of the builder.
func (lruo *LoginRequestUpdateOne) Mutation() *LoginRequestMutation {
	return lruo.mutation
}

// Where appends a list predicates to the LoginRequestUpdate builder.
func (lruo *LoginRequestUpdateOne) Where(ps ...predicate.LoginRequest) *LoginRequestUpdateOne {
	lruo.mutation.Where(ps...)
	return lruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lruo *LoginRequestUpdateOne) Select(field string, fields ...string) *LoginRequestUpdateOne {
	lruo.fields = append([]string{field}, fields...)
	return lruo
}

// Save executes the query and returns the updated LoginRequest entity.
func (lruo *LoginRequestUpdateOne) Save(ctx context.Context) (*LoginRequest, error) {
	return withHooks(ctx, lruo.sqlSave, lruo.mutation, lruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lruo *LoginRequestUpdateOne) SaveX(ctx context.Context) *LoginRequest {
	node, err := lruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lruo *LoginRequestUpdateOne) Exec(ctx context.Context) error {
	_, err := lruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lruo *LoginRequestUpdateOne) ExecX(ctx context.Context) {
	if err := lruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lruo *LoginRequestUpdateOne) check() error {
	if v, ok := lruo.mutation.Status(); ok {
		if err := loginrequest.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "LoginRequest.status": %w`, err)}
		}
	}
	return nil
}

func (lruo *LoginRequestUpdateOne) sqlSave(ctx context.Context) (_node *LoginRequest, err error) {
	if err := lruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(loginrequest.Table, loginrequest.Columns, sqlgraph.NewFieldSpec(loginrequest.FieldID, field.TypeUUID))
	id, ok := lruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LoginRequest.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, loginrequest.FieldID)
		for _, f := range fields {
			if !loginrequest.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != loginrequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lruo.mutation.UserId(); ok {
		_spec.SetField(loginrequest.FieldUserId, field.TypeUUID, value)
	}
	if value, ok := lruo.mutation.Email(); ok {
		_spec.SetField(loginrequest.FieldEmail, field.TypeString, value)
	}
	if value, ok := lruo.mutation.LoginRequestCode(); ok {
		_spec.SetField(loginrequest.FieldLoginRequestCode, field.TypeString, value)
	}
	if value, ok := lruo.mutation.Status(); ok {
		_spec.SetField(loginrequest.FieldStatus, field.TypeEnum, value)
	}
	_node = &LoginRequest{config: lruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{loginrequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	lruo.mutation.done = true
	return _node, nil
}
