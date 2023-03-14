// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/yowmark/go-ent-pokemon/ent/battle"
	"github.com/yowmark/go-ent-pokemon/ent/pokemon"
)

// Battle is the model entity for the Battle schema.
type Battle struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"oid,omitempty"`
	// Result holds the value of the "result" field.
	Result string `json:"result,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BattleQuery when eager-loading is set.
	Edges             BattleEdges `json:"edges"`
	pokemon_fights    *int
	pokemon_opponents *int
}

// BattleEdges holds the relations/edges for other nodes in the graph.
type BattleEdges struct {
	// Contender holds the value of the contender edge.
	Contender *Pokemon `json:"contender,omitempty"`
	// Oponent holds the value of the oponent edge.
	Oponent *Pokemon `json:"oponent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ContenderOrErr returns the Contender value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BattleEdges) ContenderOrErr() (*Pokemon, error) {
	if e.loadedTypes[0] {
		if e.Contender == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: pokemon.Label}
		}
		return e.Contender, nil
	}
	return nil, &NotLoadedError{edge: "contender"}
}

// OponentOrErr returns the Oponent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BattleEdges) OponentOrErr() (*Pokemon, error) {
	if e.loadedTypes[1] {
		if e.Oponent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: pokemon.Label}
		}
		return e.Oponent, nil
	}
	return nil, &NotLoadedError{edge: "oponent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Battle) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case battle.FieldID:
			values[i] = new(sql.NullInt64)
		case battle.FieldResult:
			values[i] = new(sql.NullString)
		case battle.FieldCreatedAt, battle.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case battle.ForeignKeys[0]: // pokemon_fights
			values[i] = new(sql.NullInt64)
		case battle.ForeignKeys[1]: // pokemon_opponents
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Battle", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Battle fields.
func (b *Battle) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case battle.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case battle.FieldResult:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field result", values[i])
			} else if value.Valid {
				b.Result = value.String
			}
		case battle.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case battle.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		case battle.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field pokemon_fights", value)
			} else if value.Valid {
				b.pokemon_fights = new(int)
				*b.pokemon_fights = int(value.Int64)
			}
		case battle.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field pokemon_opponents", value)
			} else if value.Valid {
				b.pokemon_opponents = new(int)
				*b.pokemon_opponents = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryContender queries the "contender" edge of the Battle entity.
func (b *Battle) QueryContender() *PokemonQuery {
	return NewBattleClient(b.config).QueryContender(b)
}

// QueryOponent queries the "oponent" edge of the Battle entity.
func (b *Battle) QueryOponent() *PokemonQuery {
	return NewBattleClient(b.config).QueryOponent(b)
}

// Update returns a builder for updating this Battle.
// Note that you need to call Battle.Unwrap() before calling this method if this Battle
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Battle) Update() *BattleUpdateOne {
	return NewBattleClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Battle entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Battle) Unwrap() *Battle {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Battle is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Battle) String() string {
	var builder strings.Builder
	builder.WriteString("Battle(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("result=")
	builder.WriteString(b.Result)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Battles is a parsable slice of Battle.
type Battles []*Battle
