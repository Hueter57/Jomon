// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/traPtitech/Jomon/ent/file"
	"github.com/traPtitech/Jomon/ent/request"
	"github.com/traPtitech/Jomon/ent/user"
)

// FileCreate is the builder for creating a File entity.
type FileCreate struct {
	config
	mutation *FileMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (fc *FileCreate) SetName(s string) *FileCreate {
	fc.mutation.SetName(s)
	return fc
}

// SetMimeType sets the "mime_type" field.
func (fc *FileCreate) SetMimeType(s string) *FileCreate {
	fc.mutation.SetMimeType(s)
	return fc
}

// SetCreatedAt sets the "created_at" field.
func (fc *FileCreate) SetCreatedAt(t time.Time) *FileCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FileCreate) SetNillableCreatedAt(t *time.Time) *FileCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// SetDeletedAt sets the "deleted_at" field.
func (fc *FileCreate) SetDeletedAt(t time.Time) *FileCreate {
	fc.mutation.SetDeletedAt(t)
	return fc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fc *FileCreate) SetNillableDeletedAt(t *time.Time) *FileCreate {
	if t != nil {
		fc.SetDeletedAt(*t)
	}
	return fc
}

// SetID sets the "id" field.
func (fc *FileCreate) SetID(u uuid.UUID) *FileCreate {
	fc.mutation.SetID(u)
	return fc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fc *FileCreate) SetNillableID(u *uuid.UUID) *FileCreate {
	if u != nil {
		fc.SetID(*u)
	}
	return fc
}

// SetRequestID sets the "request" edge to the Request entity by ID.
func (fc *FileCreate) SetRequestID(id uuid.UUID) *FileCreate {
	fc.mutation.SetRequestID(id)
	return fc
}

// SetNillableRequestID sets the "request" edge to the Request entity by ID if the given value is not nil.
func (fc *FileCreate) SetNillableRequestID(id *uuid.UUID) *FileCreate {
	if id != nil {
		fc = fc.SetRequestID(*id)
	}
	return fc
}

// SetRequest sets the "request" edge to the Request entity.
func (fc *FileCreate) SetRequest(r *Request) *FileCreate {
	return fc.SetRequestID(r.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (fc *FileCreate) SetUserID(id uuid.UUID) *FileCreate {
	fc.mutation.SetUserID(id)
	return fc
}

// SetUser sets the "user" edge to the User entity.
func (fc *FileCreate) SetUser(u *User) *FileCreate {
	return fc.SetUserID(u.ID)
}

// Mutation returns the FileMutation object of the builder.
func (fc *FileCreate) Mutation() *FileMutation {
	return fc.mutation
}

// Save creates the File in the database.
func (fc *FileCreate) Save(ctx context.Context) (*File, error) {
	fc.defaults()
	return withHooks[*File, FileMutation](ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FileCreate) SaveX(ctx context.Context) *File {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FileCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FileCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FileCreate) defaults() {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		v := file.DefaultCreatedAt()
		fc.mutation.SetCreatedAt(v)
	}
	if _, ok := fc.mutation.ID(); !ok {
		v := file.DefaultID()
		fc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FileCreate) check() error {
	if _, ok := fc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "File.name"`)}
	}
	if v, ok := fc.mutation.Name(); ok {
		if err := file.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "File.name": %w`, err)}
		}
	}
	if _, ok := fc.mutation.MimeType(); !ok {
		return &ValidationError{Name: "mime_type", err: errors.New(`ent: missing required field "File.mime_type"`)}
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "File.created_at"`)}
	}
	if _, ok := fc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "File.user"`)}
	}
	return nil
}

func (fc *FileCreate) sqlSave(ctx context.Context) (*File, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
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
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FileCreate) createSpec() (*File, *sqlgraph.CreateSpec) {
	var (
		_node = &File{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(file.Table, sqlgraph.NewFieldSpec(file.FieldID, field.TypeUUID))
	)
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := fc.mutation.Name(); ok {
		_spec.SetField(file.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := fc.mutation.MimeType(); ok {
		_spec.SetField(file.FieldMimeType, field.TypeString, value)
		_node.MimeType = value
	}
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.SetField(file.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := fc.mutation.DeletedAt(); ok {
		_spec.SetField(file.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if nodes := fc.mutation.RequestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.RequestTable,
			Columns: []string{file.RequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: request.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.request_file = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   file.UserTable,
			Columns: []string{file.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.file_user = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FileCreateBulk is the builder for creating many File entities in bulk.
type FileCreateBulk struct {
	config
	builders []*FileCreate
}

// Save creates the File entities in the database.
func (fcb *FileCreateBulk) Save(ctx context.Context) ([]*File, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*File, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FileMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FileCreateBulk) SaveX(ctx context.Context) []*File {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FileCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FileCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}
