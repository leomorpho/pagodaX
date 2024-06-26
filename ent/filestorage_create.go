// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/filestorage"
)

// FileStorageCreate is the builder for creating a FileStorage entity.
type FileStorageCreate struct {
	config
	mutation *FileStorageMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (fsc *FileStorageCreate) SetCreatedAt(t time.Time) *FileStorageCreate {
	fsc.mutation.SetCreatedAt(t)
	return fsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fsc *FileStorageCreate) SetNillableCreatedAt(t *time.Time) *FileStorageCreate {
	if t != nil {
		fsc.SetCreatedAt(*t)
	}
	return fsc
}

// SetUpdatedAt sets the "updated_at" field.
func (fsc *FileStorageCreate) SetUpdatedAt(t time.Time) *FileStorageCreate {
	fsc.mutation.SetUpdatedAt(t)
	return fsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fsc *FileStorageCreate) SetNillableUpdatedAt(t *time.Time) *FileStorageCreate {
	if t != nil {
		fsc.SetUpdatedAt(*t)
	}
	return fsc
}

// SetBucketName sets the "bucket_name" field.
func (fsc *FileStorageCreate) SetBucketName(s string) *FileStorageCreate {
	fsc.mutation.SetBucketName(s)
	return fsc
}

// SetObjectKey sets the "object_key" field.
func (fsc *FileStorageCreate) SetObjectKey(s string) *FileStorageCreate {
	fsc.mutation.SetObjectKey(s)
	return fsc
}

// SetOriginalFileName sets the "original_file_name" field.
func (fsc *FileStorageCreate) SetOriginalFileName(s string) *FileStorageCreate {
	fsc.mutation.SetOriginalFileName(s)
	return fsc
}

// SetNillableOriginalFileName sets the "original_file_name" field if the given value is not nil.
func (fsc *FileStorageCreate) SetNillableOriginalFileName(s *string) *FileStorageCreate {
	if s != nil {
		fsc.SetOriginalFileName(*s)
	}
	return fsc
}

// SetFileSize sets the "file_size" field.
func (fsc *FileStorageCreate) SetFileSize(i int64) *FileStorageCreate {
	fsc.mutation.SetFileSize(i)
	return fsc
}

// SetNillableFileSize sets the "file_size" field if the given value is not nil.
func (fsc *FileStorageCreate) SetNillableFileSize(i *int64) *FileStorageCreate {
	if i != nil {
		fsc.SetFileSize(*i)
	}
	return fsc
}

// SetContentType sets the "content_type" field.
func (fsc *FileStorageCreate) SetContentType(s string) *FileStorageCreate {
	fsc.mutation.SetContentType(s)
	return fsc
}

// SetNillableContentType sets the "content_type" field if the given value is not nil.
func (fsc *FileStorageCreate) SetNillableContentType(s *string) *FileStorageCreate {
	if s != nil {
		fsc.SetContentType(*s)
	}
	return fsc
}

// SetFileHash sets the "file_hash" field.
func (fsc *FileStorageCreate) SetFileHash(s string) *FileStorageCreate {
	fsc.mutation.SetFileHash(s)
	return fsc
}

// SetNillableFileHash sets the "file_hash" field if the given value is not nil.
func (fsc *FileStorageCreate) SetNillableFileHash(s *string) *FileStorageCreate {
	if s != nil {
		fsc.SetFileHash(*s)
	}
	return fsc
}

// Mutation returns the FileStorageMutation object of the builder.
func (fsc *FileStorageCreate) Mutation() *FileStorageMutation {
	return fsc.mutation
}

// Save creates the FileStorage in the database.
func (fsc *FileStorageCreate) Save(ctx context.Context) (*FileStorage, error) {
	fsc.defaults()
	return withHooks(ctx, fsc.sqlSave, fsc.mutation, fsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fsc *FileStorageCreate) SaveX(ctx context.Context) *FileStorage {
	v, err := fsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fsc *FileStorageCreate) Exec(ctx context.Context) error {
	_, err := fsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fsc *FileStorageCreate) ExecX(ctx context.Context) {
	if err := fsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fsc *FileStorageCreate) defaults() {
	if _, ok := fsc.mutation.CreatedAt(); !ok {
		v := filestorage.DefaultCreatedAt()
		fsc.mutation.SetCreatedAt(v)
	}
	if _, ok := fsc.mutation.UpdatedAt(); !ok {
		v := filestorage.DefaultUpdatedAt()
		fsc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fsc *FileStorageCreate) check() error {
	if _, ok := fsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "FileStorage.created_at"`)}
	}
	if _, ok := fsc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "FileStorage.updated_at"`)}
	}
	if _, ok := fsc.mutation.BucketName(); !ok {
		return &ValidationError{Name: "bucket_name", err: errors.New(`ent: missing required field "FileStorage.bucket_name"`)}
	}
	if v, ok := fsc.mutation.BucketName(); ok {
		if err := filestorage.BucketNameValidator(v); err != nil {
			return &ValidationError{Name: "bucket_name", err: fmt.Errorf(`ent: validator failed for field "FileStorage.bucket_name": %w`, err)}
		}
	}
	if _, ok := fsc.mutation.ObjectKey(); !ok {
		return &ValidationError{Name: "object_key", err: errors.New(`ent: missing required field "FileStorage.object_key"`)}
	}
	if v, ok := fsc.mutation.ObjectKey(); ok {
		if err := filestorage.ObjectKeyValidator(v); err != nil {
			return &ValidationError{Name: "object_key", err: fmt.Errorf(`ent: validator failed for field "FileStorage.object_key": %w`, err)}
		}
	}
	return nil
}

func (fsc *FileStorageCreate) sqlSave(ctx context.Context) (*FileStorage, error) {
	if err := fsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	fsc.mutation.id = &_node.ID
	fsc.mutation.done = true
	return _node, nil
}

func (fsc *FileStorageCreate) createSpec() (*FileStorage, *sqlgraph.CreateSpec) {
	var (
		_node = &FileStorage{config: fsc.config}
		_spec = sqlgraph.NewCreateSpec(filestorage.Table, sqlgraph.NewFieldSpec(filestorage.FieldID, field.TypeInt))
	)
	_spec.OnConflict = fsc.conflict
	if value, ok := fsc.mutation.CreatedAt(); ok {
		_spec.SetField(filestorage.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := fsc.mutation.UpdatedAt(); ok {
		_spec.SetField(filestorage.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := fsc.mutation.BucketName(); ok {
		_spec.SetField(filestorage.FieldBucketName, field.TypeString, value)
		_node.BucketName = value
	}
	if value, ok := fsc.mutation.ObjectKey(); ok {
		_spec.SetField(filestorage.FieldObjectKey, field.TypeString, value)
		_node.ObjectKey = value
	}
	if value, ok := fsc.mutation.OriginalFileName(); ok {
		_spec.SetField(filestorage.FieldOriginalFileName, field.TypeString, value)
		_node.OriginalFileName = value
	}
	if value, ok := fsc.mutation.FileSize(); ok {
		_spec.SetField(filestorage.FieldFileSize, field.TypeInt64, value)
		_node.FileSize = value
	}
	if value, ok := fsc.mutation.ContentType(); ok {
		_spec.SetField(filestorage.FieldContentType, field.TypeString, value)
		_node.ContentType = value
	}
	if value, ok := fsc.mutation.FileHash(); ok {
		_spec.SetField(filestorage.FieldFileHash, field.TypeString, value)
		_node.FileHash = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FileStorage.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FileStorageUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (fsc *FileStorageCreate) OnConflict(opts ...sql.ConflictOption) *FileStorageUpsertOne {
	fsc.conflict = opts
	return &FileStorageUpsertOne{
		create: fsc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FileStorage.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fsc *FileStorageCreate) OnConflictColumns(columns ...string) *FileStorageUpsertOne {
	fsc.conflict = append(fsc.conflict, sql.ConflictColumns(columns...))
	return &FileStorageUpsertOne{
		create: fsc,
	}
}

type (
	// FileStorageUpsertOne is the builder for "upsert"-ing
	//  one FileStorage node.
	FileStorageUpsertOne struct {
		create *FileStorageCreate
	}

	// FileStorageUpsert is the "OnConflict" setter.
	FileStorageUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *FileStorageUpsert) SetUpdatedAt(v time.Time) *FileStorageUpsert {
	u.Set(filestorage.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FileStorageUpsert) UpdateUpdatedAt() *FileStorageUpsert {
	u.SetExcluded(filestorage.FieldUpdatedAt)
	return u
}

// SetBucketName sets the "bucket_name" field.
func (u *FileStorageUpsert) SetBucketName(v string) *FileStorageUpsert {
	u.Set(filestorage.FieldBucketName, v)
	return u
}

// UpdateBucketName sets the "bucket_name" field to the value that was provided on create.
func (u *FileStorageUpsert) UpdateBucketName() *FileStorageUpsert {
	u.SetExcluded(filestorage.FieldBucketName)
	return u
}

// SetObjectKey sets the "object_key" field.
func (u *FileStorageUpsert) SetObjectKey(v string) *FileStorageUpsert {
	u.Set(filestorage.FieldObjectKey, v)
	return u
}

// UpdateObjectKey sets the "object_key" field to the value that was provided on create.
func (u *FileStorageUpsert) UpdateObjectKey() *FileStorageUpsert {
	u.SetExcluded(filestorage.FieldObjectKey)
	return u
}

// SetOriginalFileName sets the "original_file_name" field.
func (u *FileStorageUpsert) SetOriginalFileName(v string) *FileStorageUpsert {
	u.Set(filestorage.FieldOriginalFileName, v)
	return u
}

// UpdateOriginalFileName sets the "original_file_name" field to the value that was provided on create.
func (u *FileStorageUpsert) UpdateOriginalFileName() *FileStorageUpsert {
	u.SetExcluded(filestorage.FieldOriginalFileName)
	return u
}

// ClearOriginalFileName clears the value of the "original_file_name" field.
func (u *FileStorageUpsert) ClearOriginalFileName() *FileStorageUpsert {
	u.SetNull(filestorage.FieldOriginalFileName)
	return u
}

// SetFileSize sets the "file_size" field.
func (u *FileStorageUpsert) SetFileSize(v int64) *FileStorageUpsert {
	u.Set(filestorage.FieldFileSize, v)
	return u
}

// UpdateFileSize sets the "file_size" field to the value that was provided on create.
func (u *FileStorageUpsert) UpdateFileSize() *FileStorageUpsert {
	u.SetExcluded(filestorage.FieldFileSize)
	return u
}

// AddFileSize adds v to the "file_size" field.
func (u *FileStorageUpsert) AddFileSize(v int64) *FileStorageUpsert {
	u.Add(filestorage.FieldFileSize, v)
	return u
}

// ClearFileSize clears the value of the "file_size" field.
func (u *FileStorageUpsert) ClearFileSize() *FileStorageUpsert {
	u.SetNull(filestorage.FieldFileSize)
	return u
}

// SetContentType sets the "content_type" field.
func (u *FileStorageUpsert) SetContentType(v string) *FileStorageUpsert {
	u.Set(filestorage.FieldContentType, v)
	return u
}

// UpdateContentType sets the "content_type" field to the value that was provided on create.
func (u *FileStorageUpsert) UpdateContentType() *FileStorageUpsert {
	u.SetExcluded(filestorage.FieldContentType)
	return u
}

// ClearContentType clears the value of the "content_type" field.
func (u *FileStorageUpsert) ClearContentType() *FileStorageUpsert {
	u.SetNull(filestorage.FieldContentType)
	return u
}

// SetFileHash sets the "file_hash" field.
func (u *FileStorageUpsert) SetFileHash(v string) *FileStorageUpsert {
	u.Set(filestorage.FieldFileHash, v)
	return u
}

// UpdateFileHash sets the "file_hash" field to the value that was provided on create.
func (u *FileStorageUpsert) UpdateFileHash() *FileStorageUpsert {
	u.SetExcluded(filestorage.FieldFileHash)
	return u
}

// ClearFileHash clears the value of the "file_hash" field.
func (u *FileStorageUpsert) ClearFileHash() *FileStorageUpsert {
	u.SetNull(filestorage.FieldFileHash)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.FileStorage.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *FileStorageUpsertOne) UpdateNewValues() *FileStorageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(filestorage.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FileStorage.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FileStorageUpsertOne) Ignore() *FileStorageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FileStorageUpsertOne) DoNothing() *FileStorageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FileStorageCreate.OnConflict
// documentation for more info.
func (u *FileStorageUpsertOne) Update(set func(*FileStorageUpsert)) *FileStorageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FileStorageUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FileStorageUpsertOne) SetUpdatedAt(v time.Time) *FileStorageUpsertOne {
	return u.Update(func(s *FileStorageUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FileStorageUpsertOne) UpdateUpdatedAt() *FileStorageUpsertOne {
	return u.Update(func(s *FileStorageUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetBucketName sets the "bucket_name" field.
func (u *FileStorageUpsertOne) SetBucketName(v string) *FileStorageUpsertOne {
	return u.Update(func(s *FileStorageUpsert) {
		s.SetBucketName(v)
	})
}

// UpdateBucketName sets the "bucket_name" field to the value that was provided on create.
func (u *FileStorageUpsertOne) UpdateBucketName() *FileStorageUpsertOne {
	return u.Update(func(s *FileStorageUpsert) {
		s.UpdateBucketName()
	})
}

// SetObjectKey sets the "object_key" field.
func (u *FileStorageUpsertOne) SetObjectKey(v string) *FileStorageUpsertOne {
	return u.Update(func(s *FileStorageUpsert) {
		s.SetObjectKey(v)
	})
}

// UpdateObjectKey sets the "object_key" field to the value that was provided on create.
func (u *FileStorageUpsertOne) UpdateObjectKey() *FileStorageUpsertOne {
	return u.Update(func(s *FileStorageUpsert) {
		s.UpdateObjectKey()
	})
}

// SetOriginalFileName sets the "original_file_name" field.
func (u *FileStorageUpsertOne) SetOriginalFileName(v string) *FileStorageUpsertOne {
	return u.Update(func(s *FileStorageUpsert) {
		s.SetOriginalFileName(v)
	})
}

// UpdateOriginalFileName sets the "original_file_name" field to the value that was provided on create.
func (u *FileStorageUpsertOne) UpdateOriginalFileName() *FileStorageUpsertOne {
	return u.Update(func(s *FileStorageUpsert) {
		s.UpdateOriginalFileName()
	})
}

// ClearOriginalFileName clears the value of the "original_file_name" field.
func (u *FileStorageUpsertOne) ClearOriginalFileName() *FileStorageUpsertOne {
	return u.Update(func(s *FileStorageUpsert) {
		s.ClearOriginalFileName()
	})
}

// SetFileSize sets the "file_size" field.
func (u *FileStorageUpsertOne) SetFileSize(v int64) *FileStorageUpsertOne {
	return u.Update(func(s *FileStorageUpsert) {
		s.SetFileSize(v)
	})
}

// AddFileSize adds v to the "file_size" field.
func (u *FileStorageUpsertOne) AddFileSize(v int64) *FileStorageUpsertOne {
	return u.Update(func(s *FileStorageUpsert) {
		s.AddFileSize(v)
	})
}

// UpdateFileSize sets the "file_size" field to the value that was provided on create.
func (u *FileStorageUpsertOne) UpdateFileSize() *FileStorageUpsertOne {
	return u.Update(func(s *FileStorageUpsert) {
		s.UpdateFileSize()
	})
}

// ClearFileSize clears the value of the "file_size" field.
func (u *FileStorageUpsertOne) ClearFileSize() *FileStorageUpsertOne {
	return u.Update(func(s *FileStorageUpsert) {
		s.ClearFileSize()
	})
}

// SetContentType sets the "content_type" field.
func (u *FileStorageUpsertOne) SetContentType(v string) *FileStorageUpsertOne {
	return u.Update(func(s *FileStorageUpsert) {
		s.SetContentType(v)
	})
}

// UpdateContentType sets the "content_type" field to the value that was provided on create.
func (u *FileStorageUpsertOne) UpdateContentType() *FileStorageUpsertOne {
	return u.Update(func(s *FileStorageUpsert) {
		s.UpdateContentType()
	})
}

// ClearContentType clears the value of the "content_type" field.
func (u *FileStorageUpsertOne) ClearContentType() *FileStorageUpsertOne {
	return u.Update(func(s *FileStorageUpsert) {
		s.ClearContentType()
	})
}

// SetFileHash sets the "file_hash" field.
func (u *FileStorageUpsertOne) SetFileHash(v string) *FileStorageUpsertOne {
	return u.Update(func(s *FileStorageUpsert) {
		s.SetFileHash(v)
	})
}

// UpdateFileHash sets the "file_hash" field to the value that was provided on create.
func (u *FileStorageUpsertOne) UpdateFileHash() *FileStorageUpsertOne {
	return u.Update(func(s *FileStorageUpsert) {
		s.UpdateFileHash()
	})
}

// ClearFileHash clears the value of the "file_hash" field.
func (u *FileStorageUpsertOne) ClearFileHash() *FileStorageUpsertOne {
	return u.Update(func(s *FileStorageUpsert) {
		s.ClearFileHash()
	})
}

// Exec executes the query.
func (u *FileStorageUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FileStorageCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FileStorageUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FileStorageUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FileStorageUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FileStorageCreateBulk is the builder for creating many FileStorage entities in bulk.
type FileStorageCreateBulk struct {
	config
	err      error
	builders []*FileStorageCreate
	conflict []sql.ConflictOption
}

// Save creates the FileStorage entities in the database.
func (fscb *FileStorageCreateBulk) Save(ctx context.Context) ([]*FileStorage, error) {
	if fscb.err != nil {
		return nil, fscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(fscb.builders))
	nodes := make([]*FileStorage, len(fscb.builders))
	mutators := make([]Mutator, len(fscb.builders))
	for i := range fscb.builders {
		func(i int, root context.Context) {
			builder := fscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FileStorageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
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
		if _, err := mutators[0].Mutate(ctx, fscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fscb *FileStorageCreateBulk) SaveX(ctx context.Context) []*FileStorage {
	v, err := fscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fscb *FileStorageCreateBulk) Exec(ctx context.Context) error {
	_, err := fscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fscb *FileStorageCreateBulk) ExecX(ctx context.Context) {
	if err := fscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FileStorage.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FileStorageUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (fscb *FileStorageCreateBulk) OnConflict(opts ...sql.ConflictOption) *FileStorageUpsertBulk {
	fscb.conflict = opts
	return &FileStorageUpsertBulk{
		create: fscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FileStorage.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fscb *FileStorageCreateBulk) OnConflictColumns(columns ...string) *FileStorageUpsertBulk {
	fscb.conflict = append(fscb.conflict, sql.ConflictColumns(columns...))
	return &FileStorageUpsertBulk{
		create: fscb,
	}
}

// FileStorageUpsertBulk is the builder for "upsert"-ing
// a bulk of FileStorage nodes.
type FileStorageUpsertBulk struct {
	create *FileStorageCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.FileStorage.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *FileStorageUpsertBulk) UpdateNewValues() *FileStorageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(filestorage.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FileStorage.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FileStorageUpsertBulk) Ignore() *FileStorageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FileStorageUpsertBulk) DoNothing() *FileStorageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FileStorageCreateBulk.OnConflict
// documentation for more info.
func (u *FileStorageUpsertBulk) Update(set func(*FileStorageUpsert)) *FileStorageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FileStorageUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FileStorageUpsertBulk) SetUpdatedAt(v time.Time) *FileStorageUpsertBulk {
	return u.Update(func(s *FileStorageUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FileStorageUpsertBulk) UpdateUpdatedAt() *FileStorageUpsertBulk {
	return u.Update(func(s *FileStorageUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetBucketName sets the "bucket_name" field.
func (u *FileStorageUpsertBulk) SetBucketName(v string) *FileStorageUpsertBulk {
	return u.Update(func(s *FileStorageUpsert) {
		s.SetBucketName(v)
	})
}

// UpdateBucketName sets the "bucket_name" field to the value that was provided on create.
func (u *FileStorageUpsertBulk) UpdateBucketName() *FileStorageUpsertBulk {
	return u.Update(func(s *FileStorageUpsert) {
		s.UpdateBucketName()
	})
}

// SetObjectKey sets the "object_key" field.
func (u *FileStorageUpsertBulk) SetObjectKey(v string) *FileStorageUpsertBulk {
	return u.Update(func(s *FileStorageUpsert) {
		s.SetObjectKey(v)
	})
}

// UpdateObjectKey sets the "object_key" field to the value that was provided on create.
func (u *FileStorageUpsertBulk) UpdateObjectKey() *FileStorageUpsertBulk {
	return u.Update(func(s *FileStorageUpsert) {
		s.UpdateObjectKey()
	})
}

// SetOriginalFileName sets the "original_file_name" field.
func (u *FileStorageUpsertBulk) SetOriginalFileName(v string) *FileStorageUpsertBulk {
	return u.Update(func(s *FileStorageUpsert) {
		s.SetOriginalFileName(v)
	})
}

// UpdateOriginalFileName sets the "original_file_name" field to the value that was provided on create.
func (u *FileStorageUpsertBulk) UpdateOriginalFileName() *FileStorageUpsertBulk {
	return u.Update(func(s *FileStorageUpsert) {
		s.UpdateOriginalFileName()
	})
}

// ClearOriginalFileName clears the value of the "original_file_name" field.
func (u *FileStorageUpsertBulk) ClearOriginalFileName() *FileStorageUpsertBulk {
	return u.Update(func(s *FileStorageUpsert) {
		s.ClearOriginalFileName()
	})
}

// SetFileSize sets the "file_size" field.
func (u *FileStorageUpsertBulk) SetFileSize(v int64) *FileStorageUpsertBulk {
	return u.Update(func(s *FileStorageUpsert) {
		s.SetFileSize(v)
	})
}

// AddFileSize adds v to the "file_size" field.
func (u *FileStorageUpsertBulk) AddFileSize(v int64) *FileStorageUpsertBulk {
	return u.Update(func(s *FileStorageUpsert) {
		s.AddFileSize(v)
	})
}

// UpdateFileSize sets the "file_size" field to the value that was provided on create.
func (u *FileStorageUpsertBulk) UpdateFileSize() *FileStorageUpsertBulk {
	return u.Update(func(s *FileStorageUpsert) {
		s.UpdateFileSize()
	})
}

// ClearFileSize clears the value of the "file_size" field.
func (u *FileStorageUpsertBulk) ClearFileSize() *FileStorageUpsertBulk {
	return u.Update(func(s *FileStorageUpsert) {
		s.ClearFileSize()
	})
}

// SetContentType sets the "content_type" field.
func (u *FileStorageUpsertBulk) SetContentType(v string) *FileStorageUpsertBulk {
	return u.Update(func(s *FileStorageUpsert) {
		s.SetContentType(v)
	})
}

// UpdateContentType sets the "content_type" field to the value that was provided on create.
func (u *FileStorageUpsertBulk) UpdateContentType() *FileStorageUpsertBulk {
	return u.Update(func(s *FileStorageUpsert) {
		s.UpdateContentType()
	})
}

// ClearContentType clears the value of the "content_type" field.
func (u *FileStorageUpsertBulk) ClearContentType() *FileStorageUpsertBulk {
	return u.Update(func(s *FileStorageUpsert) {
		s.ClearContentType()
	})
}

// SetFileHash sets the "file_hash" field.
func (u *FileStorageUpsertBulk) SetFileHash(v string) *FileStorageUpsertBulk {
	return u.Update(func(s *FileStorageUpsert) {
		s.SetFileHash(v)
	})
}

// UpdateFileHash sets the "file_hash" field to the value that was provided on create.
func (u *FileStorageUpsertBulk) UpdateFileHash() *FileStorageUpsertBulk {
	return u.Update(func(s *FileStorageUpsert) {
		s.UpdateFileHash()
	})
}

// ClearFileHash clears the value of the "file_hash" field.
func (u *FileStorageUpsertBulk) ClearFileHash() *FileStorageUpsertBulk {
	return u.Update(func(s *FileStorageUpsert) {
		s.ClearFileHash()
	})
}

// Exec executes the query.
func (u *FileStorageUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FileStorageCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FileStorageCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FileStorageUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
