// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yowmark/go-ent-pokemon/ent/battle"
	"github.com/yowmark/go-ent-pokemon/ent/pokemon"
)

// PokemonCreate is the builder for creating a Pokemon entity.
type PokemonCreate struct {
	config
	mutation *PokemonMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *PokemonCreate) SetName(s string) *PokemonCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *PokemonCreate) SetDescription(s string) *PokemonCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetWeight sets the "weight" field.
func (pc *PokemonCreate) SetWeight(f float64) *PokemonCreate {
	pc.mutation.SetWeight(f)
	return pc
}

// SetHeight sets the "height" field.
func (pc *PokemonCreate) SetHeight(f float64) *PokemonCreate {
	pc.mutation.SetHeight(f)
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PokemonCreate) SetCreatedAt(t time.Time) *PokemonCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PokemonCreate) SetNillableCreatedAt(t *time.Time) *PokemonCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PokemonCreate) SetUpdatedAt(t time.Time) *PokemonCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PokemonCreate) SetNillableUpdatedAt(t *time.Time) *PokemonCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PokemonCreate) SetID(i int) *PokemonCreate {
	pc.mutation.SetID(i)
	return pc
}

// AddFightIDs adds the "fights" edge to the Battle entity by IDs.
func (pc *PokemonCreate) AddFightIDs(ids ...int) *PokemonCreate {
	pc.mutation.AddFightIDs(ids...)
	return pc
}

// AddFights adds the "fights" edges to the Battle entity.
func (pc *PokemonCreate) AddFights(b ...*Battle) *PokemonCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pc.AddFightIDs(ids...)
}

// AddOpponentIDs adds the "opponents" edge to the Battle entity by IDs.
func (pc *PokemonCreate) AddOpponentIDs(ids ...int) *PokemonCreate {
	pc.mutation.AddOpponentIDs(ids...)
	return pc
}

// AddOpponents adds the "opponents" edges to the Battle entity.
func (pc *PokemonCreate) AddOpponents(b ...*Battle) *PokemonCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pc.AddOpponentIDs(ids...)
}

// Mutation returns the PokemonMutation object of the builder.
func (pc *PokemonCreate) Mutation() *PokemonMutation {
	return pc.mutation
}

// Save creates the Pokemon in the database.
func (pc *PokemonCreate) Save(ctx context.Context) (*Pokemon, error) {
	pc.defaults()
	return withHooks[*Pokemon, PokemonMutation](ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PokemonCreate) SaveX(ctx context.Context) *Pokemon {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PokemonCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PokemonCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PokemonCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := pokemon.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := pokemon.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PokemonCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Pokemon.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := pokemon.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Pokemon.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Pokemon.description"`)}
	}
	if v, ok := pc.mutation.Description(); ok {
		if err := pokemon.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Pokemon.description": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Weight(); !ok {
		return &ValidationError{Name: "weight", err: errors.New(`ent: missing required field "Pokemon.weight"`)}
	}
	if _, ok := pc.mutation.Height(); !ok {
		return &ValidationError{Name: "height", err: errors.New(`ent: missing required field "Pokemon.height"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Pokemon.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Pokemon.updated_at"`)}
	}
	return nil
}

func (pc *PokemonCreate) sqlSave(ctx context.Context) (*Pokemon, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PokemonCreate) createSpec() (*Pokemon, *sqlgraph.CreateSpec) {
	var (
		_node = &Pokemon{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(pokemon.Table, sqlgraph.NewFieldSpec(pokemon.FieldID, field.TypeInt))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(pokemon.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(pokemon.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.Weight(); ok {
		_spec.SetField(pokemon.FieldWeight, field.TypeFloat64, value)
		_node.Weight = value
	}
	if value, ok := pc.mutation.Height(); ok {
		_spec.SetField(pokemon.FieldHeight, field.TypeFloat64, value)
		_node.Height = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(pokemon.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(pokemon.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := pc.mutation.FightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pokemon.FightsTable,
			Columns: []string{pokemon.FightsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.OpponentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pokemon.OpponentsTable,
			Columns: []string{pokemon.OpponentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PokemonCreateBulk is the builder for creating many Pokemon entities in bulk.
type PokemonCreateBulk struct {
	config
	builders []*PokemonCreate
}

// Save creates the Pokemon entities in the database.
func (pcb *PokemonCreateBulk) Save(ctx context.Context) ([]*Pokemon, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Pokemon, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PokemonMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PokemonCreateBulk) SaveX(ctx context.Context) []*Pokemon {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PokemonCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PokemonCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}