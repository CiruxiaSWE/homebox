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
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/attachment"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/document"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/item"
)

// AttachmentCreate is the builder for creating a Attachment entity.
type AttachmentCreate struct {
	config
	mutation *AttachmentMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ac *AttachmentCreate) SetCreatedAt(t time.Time) *AttachmentCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AttachmentCreate) SetNillableCreatedAt(t *time.Time) *AttachmentCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AttachmentCreate) SetUpdatedAt(t time.Time) *AttachmentCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AttachmentCreate) SetNillableUpdatedAt(t *time.Time) *AttachmentCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetType sets the "type" field.
func (ac *AttachmentCreate) SetType(a attachment.Type) *AttachmentCreate {
	ac.mutation.SetType(a)
	return ac
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ac *AttachmentCreate) SetNillableType(a *attachment.Type) *AttachmentCreate {
	if a != nil {
		ac.SetType(*a)
	}
	return ac
}

// SetPrimary sets the "primary" field.
func (ac *AttachmentCreate) SetPrimary(b bool) *AttachmentCreate {
	ac.mutation.SetPrimary(b)
	return ac
}

// SetNillablePrimary sets the "primary" field if the given value is not nil.
func (ac *AttachmentCreate) SetNillablePrimary(b *bool) *AttachmentCreate {
	if b != nil {
		ac.SetPrimary(*b)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AttachmentCreate) SetID(u uuid.UUID) *AttachmentCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AttachmentCreate) SetNillableID(u *uuid.UUID) *AttachmentCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// SetItemID sets the "item" edge to the Item entity by ID.
func (ac *AttachmentCreate) SetItemID(id uuid.UUID) *AttachmentCreate {
	ac.mutation.SetItemID(id)
	return ac
}

// SetItem sets the "item" edge to the Item entity.
func (ac *AttachmentCreate) SetItem(i *Item) *AttachmentCreate {
	return ac.SetItemID(i.ID)
}

// SetDocumentID sets the "document" edge to the Document entity by ID.
func (ac *AttachmentCreate) SetDocumentID(id uuid.UUID) *AttachmentCreate {
	ac.mutation.SetDocumentID(id)
	return ac
}

// SetDocument sets the "document" edge to the Document entity.
func (ac *AttachmentCreate) SetDocument(d *Document) *AttachmentCreate {
	return ac.SetDocumentID(d.ID)
}

// Mutation returns the AttachmentMutation object of the builder.
func (ac *AttachmentCreate) Mutation() *AttachmentMutation {
	return ac.mutation
}

// Save creates the Attachment in the database.
func (ac *AttachmentCreate) Save(ctx context.Context) (*Attachment, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AttachmentCreate) SaveX(ctx context.Context) *Attachment {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AttachmentCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AttachmentCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AttachmentCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := attachment.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := attachment.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.GetType(); !ok {
		v := attachment.DefaultType
		ac.mutation.SetType(v)
	}
	if _, ok := ac.mutation.Primary(); !ok {
		v := attachment.DefaultPrimary
		ac.mutation.SetPrimary(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := attachment.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AttachmentCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Attachment.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Attachment.updated_at"`)}
	}
	if _, ok := ac.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Attachment.type"`)}
	}
	if v, ok := ac.mutation.GetType(); ok {
		if err := attachment.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Attachment.type": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Primary(); !ok {
		return &ValidationError{Name: "primary", err: errors.New(`ent: missing required field "Attachment.primary"`)}
	}
	if len(ac.mutation.ItemIDs()) == 0 {
		return &ValidationError{Name: "item", err: errors.New(`ent: missing required edge "Attachment.item"`)}
	}
	if len(ac.mutation.DocumentIDs()) == 0 {
		return &ValidationError{Name: "document", err: errors.New(`ent: missing required edge "Attachment.document"`)}
	}
	return nil
}

func (ac *AttachmentCreate) sqlSave(ctx context.Context) (*Attachment, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
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
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AttachmentCreate) createSpec() (*Attachment, *sqlgraph.CreateSpec) {
	var (
		_node = &Attachment{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(attachment.Table, sqlgraph.NewFieldSpec(attachment.FieldID, field.TypeUUID))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(attachment.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(attachment.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.GetType(); ok {
		_spec.SetField(attachment.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := ac.mutation.Primary(); ok {
		_spec.SetField(attachment.FieldPrimary, field.TypeBool, value)
		_node.Primary = value
	}
	if nodes := ac.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.ItemTable,
			Columns: []string{attachment.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.item_attachments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.DocumentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.DocumentTable,
			Columns: []string{attachment.DocumentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(document.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.document_attachments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AttachmentCreateBulk is the builder for creating many Attachment entities in bulk.
type AttachmentCreateBulk struct {
	config
	err      error
	builders []*AttachmentCreate
}

// Save creates the Attachment entities in the database.
func (acb *AttachmentCreateBulk) Save(ctx context.Context) ([]*Attachment, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Attachment, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AttachmentMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AttachmentCreateBulk) SaveX(ctx context.Context) []*Attachment {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AttachmentCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AttachmentCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
