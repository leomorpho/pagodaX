// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/filestorage"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// FileStorageDelete is the builder for deleting a FileStorage entity.
type FileStorageDelete struct {
	config
	hooks    []Hook
	mutation *FileStorageMutation
}

// Where appends a list predicates to the FileStorageDelete builder.
func (fsd *FileStorageDelete) Where(ps ...predicate.FileStorage) *FileStorageDelete {
	fsd.mutation.Where(ps...)
	return fsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fsd *FileStorageDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, fsd.sqlExec, fsd.mutation, fsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fsd *FileStorageDelete) ExecX(ctx context.Context) int {
	n, err := fsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fsd *FileStorageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(filestorage.Table, sqlgraph.NewFieldSpec(filestorage.FieldID, field.TypeInt))
	if ps := fsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fsd.mutation.done = true
	return affected, err
}

// FileStorageDeleteOne is the builder for deleting a single FileStorage entity.
type FileStorageDeleteOne struct {
	fsd *FileStorageDelete
}

// Where appends a list predicates to the FileStorageDelete builder.
func (fsdo *FileStorageDeleteOne) Where(ps ...predicate.FileStorage) *FileStorageDeleteOne {
	fsdo.fsd.mutation.Where(ps...)
	return fsdo
}

// Exec executes the deletion query.
func (fsdo *FileStorageDeleteOne) Exec(ctx context.Context) error {
	n, err := fsdo.fsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{filestorage.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fsdo *FileStorageDeleteOne) ExecX(ctx context.Context) {
	if err := fsdo.Exec(ctx); err != nil {
		panic(err)
	}
}
