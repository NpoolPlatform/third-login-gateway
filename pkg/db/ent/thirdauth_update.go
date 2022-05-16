// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/third-login-gateway/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/third-login-gateway/pkg/db/ent/thirdauth"
	"github.com/google/uuid"
)

// ThirdAuthUpdate is the builder for updating ThirdAuth entities.
type ThirdAuthUpdate struct {
	config
	hooks    []Hook
	mutation *ThirdAuthMutation
}

// Where appends a list predicates to the ThirdAuthUpdate builder.
func (tau *ThirdAuthUpdate) Where(ps ...predicate.ThirdAuth) *ThirdAuthUpdate {
	tau.mutation.Where(ps...)
	return tau
}

// SetCreatedAt sets the "created_at" field.
func (tau *ThirdAuthUpdate) SetCreatedAt(u uint32) *ThirdAuthUpdate {
	tau.mutation.ResetCreatedAt()
	tau.mutation.SetCreatedAt(u)
	return tau
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tau *ThirdAuthUpdate) SetNillableCreatedAt(u *uint32) *ThirdAuthUpdate {
	if u != nil {
		tau.SetCreatedAt(*u)
	}
	return tau
}

// AddCreatedAt adds u to the "created_at" field.
func (tau *ThirdAuthUpdate) AddCreatedAt(u int32) *ThirdAuthUpdate {
	tau.mutation.AddCreatedAt(u)
	return tau
}

// SetUpdatedAt sets the "updated_at" field.
func (tau *ThirdAuthUpdate) SetUpdatedAt(u uint32) *ThirdAuthUpdate {
	tau.mutation.ResetUpdatedAt()
	tau.mutation.SetUpdatedAt(u)
	return tau
}

// AddUpdatedAt adds u to the "updated_at" field.
func (tau *ThirdAuthUpdate) AddUpdatedAt(u int32) *ThirdAuthUpdate {
	tau.mutation.AddUpdatedAt(u)
	return tau
}

// SetDeletedAt sets the "deleted_at" field.
func (tau *ThirdAuthUpdate) SetDeletedAt(u uint32) *ThirdAuthUpdate {
	tau.mutation.ResetDeletedAt()
	tau.mutation.SetDeletedAt(u)
	return tau
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tau *ThirdAuthUpdate) SetNillableDeletedAt(u *uint32) *ThirdAuthUpdate {
	if u != nil {
		tau.SetDeletedAt(*u)
	}
	return tau
}

// AddDeletedAt adds u to the "deleted_at" field.
func (tau *ThirdAuthUpdate) AddDeletedAt(u int32) *ThirdAuthUpdate {
	tau.mutation.AddDeletedAt(u)
	return tau
}

// SetAppID sets the "app_id" field.
func (tau *ThirdAuthUpdate) SetAppID(u uuid.UUID) *ThirdAuthUpdate {
	tau.mutation.SetAppID(u)
	return tau
}

// SetThird sets the "third" field.
func (tau *ThirdAuthUpdate) SetThird(s string) *ThirdAuthUpdate {
	tau.mutation.SetThird(s)
	return tau
}

// SetLogoURL sets the "logo_url" field.
func (tau *ThirdAuthUpdate) SetLogoURL(s string) *ThirdAuthUpdate {
	tau.mutation.SetLogoURL(s)
	return tau
}

// SetThirdAppKey sets the "third_app_key" field.
func (tau *ThirdAuthUpdate) SetThirdAppKey(s string) *ThirdAuthUpdate {
	tau.mutation.SetThirdAppKey(s)
	return tau
}

// SetThirdAppSecret sets the "third_app_secret" field.
func (tau *ThirdAuthUpdate) SetThirdAppSecret(s string) *ThirdAuthUpdate {
	tau.mutation.SetThirdAppSecret(s)
	return tau
}

// SetRedirectURL sets the "redirect_url" field.
func (tau *ThirdAuthUpdate) SetRedirectURL(s string) *ThirdAuthUpdate {
	tau.mutation.SetRedirectURL(s)
	return tau
}

// Mutation returns the ThirdAuthMutation object of the builder.
func (tau *ThirdAuthUpdate) Mutation() *ThirdAuthMutation {
	return tau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tau *ThirdAuthUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tau.defaults()
	if len(tau.hooks) == 0 {
		affected, err = tau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ThirdAuthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tau.mutation = mutation
			affected, err = tau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tau.hooks) - 1; i >= 0; i-- {
			if tau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tau *ThirdAuthUpdate) SaveX(ctx context.Context) int {
	affected, err := tau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tau *ThirdAuthUpdate) Exec(ctx context.Context) error {
	_, err := tau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tau *ThirdAuthUpdate) ExecX(ctx context.Context) {
	if err := tau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tau *ThirdAuthUpdate) defaults() {
	if _, ok := tau.mutation.UpdatedAt(); !ok {
		v := thirdauth.UpdateDefaultUpdatedAt()
		tau.mutation.SetUpdatedAt(v)
	}
}

func (tau *ThirdAuthUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   thirdauth.Table,
			Columns: thirdauth.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: thirdauth.FieldID,
			},
		},
	}
	if ps := tau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tau.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: thirdauth.FieldCreatedAt,
		})
	}
	if value, ok := tau.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: thirdauth.FieldCreatedAt,
		})
	}
	if value, ok := tau.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: thirdauth.FieldUpdatedAt,
		})
	}
	if value, ok := tau.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: thirdauth.FieldUpdatedAt,
		})
	}
	if value, ok := tau.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: thirdauth.FieldDeletedAt,
		})
	}
	if value, ok := tau.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: thirdauth.FieldDeletedAt,
		})
	}
	if value, ok := tau.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: thirdauth.FieldAppID,
		})
	}
	if value, ok := tau.mutation.Third(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: thirdauth.FieldThird,
		})
	}
	if value, ok := tau.mutation.LogoURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: thirdauth.FieldLogoURL,
		})
	}
	if value, ok := tau.mutation.ThirdAppKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: thirdauth.FieldThirdAppKey,
		})
	}
	if value, ok := tau.mutation.ThirdAppSecret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: thirdauth.FieldThirdAppSecret,
		})
	}
	if value, ok := tau.mutation.RedirectURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: thirdauth.FieldRedirectURL,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{thirdauth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ThirdAuthUpdateOne is the builder for updating a single ThirdAuth entity.
type ThirdAuthUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ThirdAuthMutation
}

// SetCreatedAt sets the "created_at" field.
func (tauo *ThirdAuthUpdateOne) SetCreatedAt(u uint32) *ThirdAuthUpdateOne {
	tauo.mutation.ResetCreatedAt()
	tauo.mutation.SetCreatedAt(u)
	return tauo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tauo *ThirdAuthUpdateOne) SetNillableCreatedAt(u *uint32) *ThirdAuthUpdateOne {
	if u != nil {
		tauo.SetCreatedAt(*u)
	}
	return tauo
}

// AddCreatedAt adds u to the "created_at" field.
func (tauo *ThirdAuthUpdateOne) AddCreatedAt(u int32) *ThirdAuthUpdateOne {
	tauo.mutation.AddCreatedAt(u)
	return tauo
}

// SetUpdatedAt sets the "updated_at" field.
func (tauo *ThirdAuthUpdateOne) SetUpdatedAt(u uint32) *ThirdAuthUpdateOne {
	tauo.mutation.ResetUpdatedAt()
	tauo.mutation.SetUpdatedAt(u)
	return tauo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (tauo *ThirdAuthUpdateOne) AddUpdatedAt(u int32) *ThirdAuthUpdateOne {
	tauo.mutation.AddUpdatedAt(u)
	return tauo
}

// SetDeletedAt sets the "deleted_at" field.
func (tauo *ThirdAuthUpdateOne) SetDeletedAt(u uint32) *ThirdAuthUpdateOne {
	tauo.mutation.ResetDeletedAt()
	tauo.mutation.SetDeletedAt(u)
	return tauo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tauo *ThirdAuthUpdateOne) SetNillableDeletedAt(u *uint32) *ThirdAuthUpdateOne {
	if u != nil {
		tauo.SetDeletedAt(*u)
	}
	return tauo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (tauo *ThirdAuthUpdateOne) AddDeletedAt(u int32) *ThirdAuthUpdateOne {
	tauo.mutation.AddDeletedAt(u)
	return tauo
}

// SetAppID sets the "app_id" field.
func (tauo *ThirdAuthUpdateOne) SetAppID(u uuid.UUID) *ThirdAuthUpdateOne {
	tauo.mutation.SetAppID(u)
	return tauo
}

// SetThird sets the "third" field.
func (tauo *ThirdAuthUpdateOne) SetThird(s string) *ThirdAuthUpdateOne {
	tauo.mutation.SetThird(s)
	return tauo
}

// SetLogoURL sets the "logo_url" field.
func (tauo *ThirdAuthUpdateOne) SetLogoURL(s string) *ThirdAuthUpdateOne {
	tauo.mutation.SetLogoURL(s)
	return tauo
}

// SetThirdAppKey sets the "third_app_key" field.
func (tauo *ThirdAuthUpdateOne) SetThirdAppKey(s string) *ThirdAuthUpdateOne {
	tauo.mutation.SetThirdAppKey(s)
	return tauo
}

// SetThirdAppSecret sets the "third_app_secret" field.
func (tauo *ThirdAuthUpdateOne) SetThirdAppSecret(s string) *ThirdAuthUpdateOne {
	tauo.mutation.SetThirdAppSecret(s)
	return tauo
}

// SetRedirectURL sets the "redirect_url" field.
func (tauo *ThirdAuthUpdateOne) SetRedirectURL(s string) *ThirdAuthUpdateOne {
	tauo.mutation.SetRedirectURL(s)
	return tauo
}

// Mutation returns the ThirdAuthMutation object of the builder.
func (tauo *ThirdAuthUpdateOne) Mutation() *ThirdAuthMutation {
	return tauo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tauo *ThirdAuthUpdateOne) Select(field string, fields ...string) *ThirdAuthUpdateOne {
	tauo.fields = append([]string{field}, fields...)
	return tauo
}

// Save executes the query and returns the updated ThirdAuth entity.
func (tauo *ThirdAuthUpdateOne) Save(ctx context.Context) (*ThirdAuth, error) {
	var (
		err  error
		node *ThirdAuth
	)
	tauo.defaults()
	if len(tauo.hooks) == 0 {
		node, err = tauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ThirdAuthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tauo.mutation = mutation
			node, err = tauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tauo.hooks) - 1; i >= 0; i-- {
			if tauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tauo *ThirdAuthUpdateOne) SaveX(ctx context.Context) *ThirdAuth {
	node, err := tauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tauo *ThirdAuthUpdateOne) Exec(ctx context.Context) error {
	_, err := tauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tauo *ThirdAuthUpdateOne) ExecX(ctx context.Context) {
	if err := tauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tauo *ThirdAuthUpdateOne) defaults() {
	if _, ok := tauo.mutation.UpdatedAt(); !ok {
		v := thirdauth.UpdateDefaultUpdatedAt()
		tauo.mutation.SetUpdatedAt(v)
	}
}

func (tauo *ThirdAuthUpdateOne) sqlSave(ctx context.Context) (_node *ThirdAuth, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   thirdauth.Table,
			Columns: thirdauth.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: thirdauth.FieldID,
			},
		},
	}
	id, ok := tauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ThirdAuth.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, thirdauth.FieldID)
		for _, f := range fields {
			if !thirdauth.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != thirdauth.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tauo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: thirdauth.FieldCreatedAt,
		})
	}
	if value, ok := tauo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: thirdauth.FieldCreatedAt,
		})
	}
	if value, ok := tauo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: thirdauth.FieldUpdatedAt,
		})
	}
	if value, ok := tauo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: thirdauth.FieldUpdatedAt,
		})
	}
	if value, ok := tauo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: thirdauth.FieldDeletedAt,
		})
	}
	if value, ok := tauo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: thirdauth.FieldDeletedAt,
		})
	}
	if value, ok := tauo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: thirdauth.FieldAppID,
		})
	}
	if value, ok := tauo.mutation.Third(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: thirdauth.FieldThird,
		})
	}
	if value, ok := tauo.mutation.LogoURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: thirdauth.FieldLogoURL,
		})
	}
	if value, ok := tauo.mutation.ThirdAppKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: thirdauth.FieldThirdAppKey,
		})
	}
	if value, ok := tauo.mutation.ThirdAppSecret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: thirdauth.FieldThirdAppSecret,
		})
	}
	if value, ok := tauo.mutation.RedirectURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: thirdauth.FieldRedirectURL,
		})
	}
	_node = &ThirdAuth{config: tauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{thirdauth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
