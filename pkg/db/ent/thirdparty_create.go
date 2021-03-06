// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/third-login-gateway/pkg/db/ent/thirdparty"
	"github.com/google/uuid"
)

// ThirdPartyCreate is the builder for creating a ThirdParty entity.
type ThirdPartyCreate struct {
	config
	mutation *ThirdPartyMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (tpc *ThirdPartyCreate) SetCreatedAt(u uint32) *ThirdPartyCreate {
	tpc.mutation.SetCreatedAt(u)
	return tpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tpc *ThirdPartyCreate) SetNillableCreatedAt(u *uint32) *ThirdPartyCreate {
	if u != nil {
		tpc.SetCreatedAt(*u)
	}
	return tpc
}

// SetUpdatedAt sets the "updated_at" field.
func (tpc *ThirdPartyCreate) SetUpdatedAt(u uint32) *ThirdPartyCreate {
	tpc.mutation.SetUpdatedAt(u)
	return tpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tpc *ThirdPartyCreate) SetNillableUpdatedAt(u *uint32) *ThirdPartyCreate {
	if u != nil {
		tpc.SetUpdatedAt(*u)
	}
	return tpc
}

// SetDeletedAt sets the "deleted_at" field.
func (tpc *ThirdPartyCreate) SetDeletedAt(u uint32) *ThirdPartyCreate {
	tpc.mutation.SetDeletedAt(u)
	return tpc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tpc *ThirdPartyCreate) SetNillableDeletedAt(u *uint32) *ThirdPartyCreate {
	if u != nil {
		tpc.SetDeletedAt(*u)
	}
	return tpc
}

// SetBrandName sets the "brand_name" field.
func (tpc *ThirdPartyCreate) SetBrandName(s string) *ThirdPartyCreate {
	tpc.mutation.SetBrandName(s)
	return tpc
}

// SetLogo sets the "logo" field.
func (tpc *ThirdPartyCreate) SetLogo(s string) *ThirdPartyCreate {
	tpc.mutation.SetLogo(s)
	return tpc
}

// SetDomain sets the "domain" field.
func (tpc *ThirdPartyCreate) SetDomain(s string) *ThirdPartyCreate {
	tpc.mutation.SetDomain(s)
	return tpc
}

// SetID sets the "id" field.
func (tpc *ThirdPartyCreate) SetID(u uuid.UUID) *ThirdPartyCreate {
	tpc.mutation.SetID(u)
	return tpc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tpc *ThirdPartyCreate) SetNillableID(u *uuid.UUID) *ThirdPartyCreate {
	if u != nil {
		tpc.SetID(*u)
	}
	return tpc
}

// Mutation returns the ThirdPartyMutation object of the builder.
func (tpc *ThirdPartyCreate) Mutation() *ThirdPartyMutation {
	return tpc.mutation
}

// Save creates the ThirdParty in the database.
func (tpc *ThirdPartyCreate) Save(ctx context.Context) (*ThirdParty, error) {
	var (
		err  error
		node *ThirdParty
	)
	if err := tpc.defaults(); err != nil {
		return nil, err
	}
	if len(tpc.hooks) == 0 {
		if err = tpc.check(); err != nil {
			return nil, err
		}
		node, err = tpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ThirdPartyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tpc.check(); err != nil {
				return nil, err
			}
			tpc.mutation = mutation
			if node, err = tpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tpc.hooks) - 1; i >= 0; i-- {
			if tpc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tpc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tpc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tpc *ThirdPartyCreate) SaveX(ctx context.Context) *ThirdParty {
	v, err := tpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tpc *ThirdPartyCreate) Exec(ctx context.Context) error {
	_, err := tpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpc *ThirdPartyCreate) ExecX(ctx context.Context) {
	if err := tpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tpc *ThirdPartyCreate) defaults() error {
	if _, ok := tpc.mutation.CreatedAt(); !ok {
		if thirdparty.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized thirdparty.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := thirdparty.DefaultCreatedAt()
		tpc.mutation.SetCreatedAt(v)
	}
	if _, ok := tpc.mutation.UpdatedAt(); !ok {
		if thirdparty.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized thirdparty.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := thirdparty.DefaultUpdatedAt()
		tpc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tpc.mutation.DeletedAt(); !ok {
		if thirdparty.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized thirdparty.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := thirdparty.DefaultDeletedAt()
		tpc.mutation.SetDeletedAt(v)
	}
	if _, ok := tpc.mutation.ID(); !ok {
		if thirdparty.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized thirdparty.DefaultID (forgotten import ent/runtime?)")
		}
		v := thirdparty.DefaultID()
		tpc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tpc *ThirdPartyCreate) check() error {
	if _, ok := tpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ThirdParty.created_at"`)}
	}
	if _, ok := tpc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ThirdParty.updated_at"`)}
	}
	if _, ok := tpc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "ThirdParty.deleted_at"`)}
	}
	if _, ok := tpc.mutation.BrandName(); !ok {
		return &ValidationError{Name: "brand_name", err: errors.New(`ent: missing required field "ThirdParty.brand_name"`)}
	}
	if _, ok := tpc.mutation.Logo(); !ok {
		return &ValidationError{Name: "logo", err: errors.New(`ent: missing required field "ThirdParty.logo"`)}
	}
	if _, ok := tpc.mutation.Domain(); !ok {
		return &ValidationError{Name: "domain", err: errors.New(`ent: missing required field "ThirdParty.domain"`)}
	}
	return nil
}

func (tpc *ThirdPartyCreate) sqlSave(ctx context.Context) (*ThirdParty, error) {
	_node, _spec := tpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (tpc *ThirdPartyCreate) createSpec() (*ThirdParty, *sqlgraph.CreateSpec) {
	var (
		_node = &ThirdParty{config: tpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: thirdparty.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: thirdparty.FieldID,
			},
		}
	)
	_spec.OnConflict = tpc.conflict
	if id, ok := tpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tpc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: thirdparty.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tpc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: thirdparty.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := tpc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: thirdparty.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := tpc.mutation.BrandName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: thirdparty.FieldBrandName,
		})
		_node.BrandName = value
	}
	if value, ok := tpc.mutation.Logo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: thirdparty.FieldLogo,
		})
		_node.Logo = value
	}
	if value, ok := tpc.mutation.Domain(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: thirdparty.FieldDomain,
		})
		_node.Domain = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ThirdParty.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ThirdPartyUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (tpc *ThirdPartyCreate) OnConflict(opts ...sql.ConflictOption) *ThirdPartyUpsertOne {
	tpc.conflict = opts
	return &ThirdPartyUpsertOne{
		create: tpc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ThirdParty.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tpc *ThirdPartyCreate) OnConflictColumns(columns ...string) *ThirdPartyUpsertOne {
	tpc.conflict = append(tpc.conflict, sql.ConflictColumns(columns...))
	return &ThirdPartyUpsertOne{
		create: tpc,
	}
}

type (
	// ThirdPartyUpsertOne is the builder for "upsert"-ing
	//  one ThirdParty node.
	ThirdPartyUpsertOne struct {
		create *ThirdPartyCreate
	}

	// ThirdPartyUpsert is the "OnConflict" setter.
	ThirdPartyUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *ThirdPartyUpsert) SetCreatedAt(v uint32) *ThirdPartyUpsert {
	u.Set(thirdparty.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ThirdPartyUpsert) UpdateCreatedAt() *ThirdPartyUpsert {
	u.SetExcluded(thirdparty.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *ThirdPartyUpsert) AddCreatedAt(v uint32) *ThirdPartyUpsert {
	u.Add(thirdparty.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ThirdPartyUpsert) SetUpdatedAt(v uint32) *ThirdPartyUpsert {
	u.Set(thirdparty.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ThirdPartyUpsert) UpdateUpdatedAt() *ThirdPartyUpsert {
	u.SetExcluded(thirdparty.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *ThirdPartyUpsert) AddUpdatedAt(v uint32) *ThirdPartyUpsert {
	u.Add(thirdparty.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ThirdPartyUpsert) SetDeletedAt(v uint32) *ThirdPartyUpsert {
	u.Set(thirdparty.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ThirdPartyUpsert) UpdateDeletedAt() *ThirdPartyUpsert {
	u.SetExcluded(thirdparty.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *ThirdPartyUpsert) AddDeletedAt(v uint32) *ThirdPartyUpsert {
	u.Add(thirdparty.FieldDeletedAt, v)
	return u
}

// SetBrandName sets the "brand_name" field.
func (u *ThirdPartyUpsert) SetBrandName(v string) *ThirdPartyUpsert {
	u.Set(thirdparty.FieldBrandName, v)
	return u
}

// UpdateBrandName sets the "brand_name" field to the value that was provided on create.
func (u *ThirdPartyUpsert) UpdateBrandName() *ThirdPartyUpsert {
	u.SetExcluded(thirdparty.FieldBrandName)
	return u
}

// SetLogo sets the "logo" field.
func (u *ThirdPartyUpsert) SetLogo(v string) *ThirdPartyUpsert {
	u.Set(thirdparty.FieldLogo, v)
	return u
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *ThirdPartyUpsert) UpdateLogo() *ThirdPartyUpsert {
	u.SetExcluded(thirdparty.FieldLogo)
	return u
}

// SetDomain sets the "domain" field.
func (u *ThirdPartyUpsert) SetDomain(v string) *ThirdPartyUpsert {
	u.Set(thirdparty.FieldDomain, v)
	return u
}

// UpdateDomain sets the "domain" field to the value that was provided on create.
func (u *ThirdPartyUpsert) UpdateDomain() *ThirdPartyUpsert {
	u.SetExcluded(thirdparty.FieldDomain)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ThirdParty.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(thirdparty.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ThirdPartyUpsertOne) UpdateNewValues() *ThirdPartyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(thirdparty.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.ThirdParty.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ThirdPartyUpsertOne) Ignore() *ThirdPartyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ThirdPartyUpsertOne) DoNothing() *ThirdPartyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ThirdPartyCreate.OnConflict
// documentation for more info.
func (u *ThirdPartyUpsertOne) Update(set func(*ThirdPartyUpsert)) *ThirdPartyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ThirdPartyUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ThirdPartyUpsertOne) SetCreatedAt(v uint32) *ThirdPartyUpsertOne {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *ThirdPartyUpsertOne) AddCreatedAt(v uint32) *ThirdPartyUpsertOne {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ThirdPartyUpsertOne) UpdateCreatedAt() *ThirdPartyUpsertOne {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ThirdPartyUpsertOne) SetUpdatedAt(v uint32) *ThirdPartyUpsertOne {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *ThirdPartyUpsertOne) AddUpdatedAt(v uint32) *ThirdPartyUpsertOne {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ThirdPartyUpsertOne) UpdateUpdatedAt() *ThirdPartyUpsertOne {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ThirdPartyUpsertOne) SetDeletedAt(v uint32) *ThirdPartyUpsertOne {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *ThirdPartyUpsertOne) AddDeletedAt(v uint32) *ThirdPartyUpsertOne {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ThirdPartyUpsertOne) UpdateDeletedAt() *ThirdPartyUpsertOne {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetBrandName sets the "brand_name" field.
func (u *ThirdPartyUpsertOne) SetBrandName(v string) *ThirdPartyUpsertOne {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.SetBrandName(v)
	})
}

// UpdateBrandName sets the "brand_name" field to the value that was provided on create.
func (u *ThirdPartyUpsertOne) UpdateBrandName() *ThirdPartyUpsertOne {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.UpdateBrandName()
	})
}

// SetLogo sets the "logo" field.
func (u *ThirdPartyUpsertOne) SetLogo(v string) *ThirdPartyUpsertOne {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.SetLogo(v)
	})
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *ThirdPartyUpsertOne) UpdateLogo() *ThirdPartyUpsertOne {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.UpdateLogo()
	})
}

// SetDomain sets the "domain" field.
func (u *ThirdPartyUpsertOne) SetDomain(v string) *ThirdPartyUpsertOne {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.SetDomain(v)
	})
}

// UpdateDomain sets the "domain" field to the value that was provided on create.
func (u *ThirdPartyUpsertOne) UpdateDomain() *ThirdPartyUpsertOne {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.UpdateDomain()
	})
}

// Exec executes the query.
func (u *ThirdPartyUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ThirdPartyCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ThirdPartyUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ThirdPartyUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ThirdPartyUpsertOne.ID is not supported by MySQL driver. Use ThirdPartyUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ThirdPartyUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ThirdPartyCreateBulk is the builder for creating many ThirdParty entities in bulk.
type ThirdPartyCreateBulk struct {
	config
	builders []*ThirdPartyCreate
	conflict []sql.ConflictOption
}

// Save creates the ThirdParty entities in the database.
func (tpcb *ThirdPartyCreateBulk) Save(ctx context.Context) ([]*ThirdParty, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tpcb.builders))
	nodes := make([]*ThirdParty, len(tpcb.builders))
	mutators := make([]Mutator, len(tpcb.builders))
	for i := range tpcb.builders {
		func(i int, root context.Context) {
			builder := tpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ThirdPartyMutation)
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
					_, err = mutators[i+1].Mutate(root, tpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tpcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tpcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, tpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tpcb *ThirdPartyCreateBulk) SaveX(ctx context.Context) []*ThirdParty {
	v, err := tpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tpcb *ThirdPartyCreateBulk) Exec(ctx context.Context) error {
	_, err := tpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpcb *ThirdPartyCreateBulk) ExecX(ctx context.Context) {
	if err := tpcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ThirdParty.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ThirdPartyUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (tpcb *ThirdPartyCreateBulk) OnConflict(opts ...sql.ConflictOption) *ThirdPartyUpsertBulk {
	tpcb.conflict = opts
	return &ThirdPartyUpsertBulk{
		create: tpcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ThirdParty.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tpcb *ThirdPartyCreateBulk) OnConflictColumns(columns ...string) *ThirdPartyUpsertBulk {
	tpcb.conflict = append(tpcb.conflict, sql.ConflictColumns(columns...))
	return &ThirdPartyUpsertBulk{
		create: tpcb,
	}
}

// ThirdPartyUpsertBulk is the builder for "upsert"-ing
// a bulk of ThirdParty nodes.
type ThirdPartyUpsertBulk struct {
	create *ThirdPartyCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ThirdParty.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(thirdparty.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ThirdPartyUpsertBulk) UpdateNewValues() *ThirdPartyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(thirdparty.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ThirdParty.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ThirdPartyUpsertBulk) Ignore() *ThirdPartyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ThirdPartyUpsertBulk) DoNothing() *ThirdPartyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ThirdPartyCreateBulk.OnConflict
// documentation for more info.
func (u *ThirdPartyUpsertBulk) Update(set func(*ThirdPartyUpsert)) *ThirdPartyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ThirdPartyUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ThirdPartyUpsertBulk) SetCreatedAt(v uint32) *ThirdPartyUpsertBulk {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *ThirdPartyUpsertBulk) AddCreatedAt(v uint32) *ThirdPartyUpsertBulk {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ThirdPartyUpsertBulk) UpdateCreatedAt() *ThirdPartyUpsertBulk {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ThirdPartyUpsertBulk) SetUpdatedAt(v uint32) *ThirdPartyUpsertBulk {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *ThirdPartyUpsertBulk) AddUpdatedAt(v uint32) *ThirdPartyUpsertBulk {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ThirdPartyUpsertBulk) UpdateUpdatedAt() *ThirdPartyUpsertBulk {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ThirdPartyUpsertBulk) SetDeletedAt(v uint32) *ThirdPartyUpsertBulk {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *ThirdPartyUpsertBulk) AddDeletedAt(v uint32) *ThirdPartyUpsertBulk {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ThirdPartyUpsertBulk) UpdateDeletedAt() *ThirdPartyUpsertBulk {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetBrandName sets the "brand_name" field.
func (u *ThirdPartyUpsertBulk) SetBrandName(v string) *ThirdPartyUpsertBulk {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.SetBrandName(v)
	})
}

// UpdateBrandName sets the "brand_name" field to the value that was provided on create.
func (u *ThirdPartyUpsertBulk) UpdateBrandName() *ThirdPartyUpsertBulk {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.UpdateBrandName()
	})
}

// SetLogo sets the "logo" field.
func (u *ThirdPartyUpsertBulk) SetLogo(v string) *ThirdPartyUpsertBulk {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.SetLogo(v)
	})
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *ThirdPartyUpsertBulk) UpdateLogo() *ThirdPartyUpsertBulk {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.UpdateLogo()
	})
}

// SetDomain sets the "domain" field.
func (u *ThirdPartyUpsertBulk) SetDomain(v string) *ThirdPartyUpsertBulk {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.SetDomain(v)
	})
}

// UpdateDomain sets the "domain" field to the value that was provided on create.
func (u *ThirdPartyUpsertBulk) UpdateDomain() *ThirdPartyUpsertBulk {
	return u.Update(func(s *ThirdPartyUpsert) {
		s.UpdateDomain()
	})
}

// Exec executes the query.
func (u *ThirdPartyUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ThirdPartyCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ThirdPartyCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ThirdPartyUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
