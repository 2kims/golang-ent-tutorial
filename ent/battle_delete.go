// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yowmark/go-ent-pokemon/ent/battle"
	"github.com/yowmark/go-ent-pokemon/ent/predicate"
)

// BattleDelete is the builder for deleting a Battle entity.
type BattleDelete struct {
	config
	hooks    []Hook
	mutation *BattleMutation
}

// Where appends a list predicates to the BattleDelete builder.
func (bd *BattleDelete) Where(ps ...predicate.Battle) *BattleDelete {
	bd.mutation.Where(ps...)
	return bd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bd *BattleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, BattleMutation](ctx, bd.sqlExec, bd.mutation, bd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bd *BattleDelete) ExecX(ctx context.Context) int {
	n, err := bd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bd *BattleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(battle.Table, sqlgraph.NewFieldSpec(battle.FieldID, field.TypeInt))
	if ps := bd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bd.mutation.done = true
	return affected, err
}

// BattleDeleteOne is the builder for deleting a single Battle entity.
type BattleDeleteOne struct {
	bd *BattleDelete
}

// Where appends a list predicates to the BattleDelete builder.
func (bdo *BattleDeleteOne) Where(ps ...predicate.Battle) *BattleDeleteOne {
	bdo.bd.mutation.Where(ps...)
	return bdo
}

// Exec executes the deletion query.
func (bdo *BattleDeleteOne) Exec(ctx context.Context) error {
	n, err := bdo.bd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{battle.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bdo *BattleDeleteOne) ExecX(ctx context.Context) {
	if err := bdo.Exec(ctx); err != nil {
		panic(err)
	}
}
