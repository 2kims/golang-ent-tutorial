// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/yowmark/go-ent-pokemon/ent/pokemon"
)

// Pokemon is the model entity for the Pokemon schema.
type Pokemon struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"oid,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Weight holds the value of the "weight" field.
	Weight float64 `json:"weight,omitempty"`
	// Height holds the value of the "height" field.
	Height float64 `json:"height,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PokemonQuery when eager-loading is set.
	Edges PokemonEdges `json:"edges"`
}

// PokemonEdges holds the relations/edges for other nodes in the graph.
type PokemonEdges struct {
	// Fights holds the value of the fights edge.
	Fights []*Battle `json:"fights,omitempty"`
	// Opponents holds the value of the opponents edge.
	Opponents []*Battle `json:"opponents,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// FightsOrErr returns the Fights value or an error if the edge
// was not loaded in eager-loading.
func (e PokemonEdges) FightsOrErr() ([]*Battle, error) {
	if e.loadedTypes[0] {
		return e.Fights, nil
	}
	return nil, &NotLoadedError{edge: "fights"}
}

// OpponentsOrErr returns the Opponents value or an error if the edge
// was not loaded in eager-loading.
func (e PokemonEdges) OpponentsOrErr() ([]*Battle, error) {
	if e.loadedTypes[1] {
		return e.Opponents, nil
	}
	return nil, &NotLoadedError{edge: "opponents"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pokemon) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case pokemon.FieldWeight, pokemon.FieldHeight:
			values[i] = new(sql.NullFloat64)
		case pokemon.FieldID:
			values[i] = new(sql.NullInt64)
		case pokemon.FieldName, pokemon.FieldDescription:
			values[i] = new(sql.NullString)
		case pokemon.FieldCreatedAt, pokemon.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Pokemon", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pokemon fields.
func (po *Pokemon) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pokemon.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			po.ID = int(value.Int64)
		case pokemon.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				po.Name = value.String
			}
		case pokemon.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				po.Description = value.String
			}
		case pokemon.FieldWeight:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				po.Weight = value.Float64
			}
		case pokemon.FieldHeight:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field height", values[i])
			} else if value.Valid {
				po.Height = value.Float64
			}
		case pokemon.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				po.CreatedAt = value.Time
			}
		case pokemon.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				po.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryFights queries the "fights" edge of the Pokemon entity.
func (po *Pokemon) QueryFights() *BattleQuery {
	return NewPokemonClient(po.config).QueryFights(po)
}

// QueryOpponents queries the "opponents" edge of the Pokemon entity.
func (po *Pokemon) QueryOpponents() *BattleQuery {
	return NewPokemonClient(po.config).QueryOpponents(po)
}

// Update returns a builder for updating this Pokemon.
// Note that you need to call Pokemon.Unwrap() before calling this method if this Pokemon
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Pokemon) Update() *PokemonUpdateOne {
	return NewPokemonClient(po.config).UpdateOne(po)
}

// Unwrap unwraps the Pokemon entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Pokemon) Unwrap() *Pokemon {
	_tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pokemon is not a transactional entity")
	}
	po.config.driver = _tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Pokemon) String() string {
	var builder strings.Builder
	builder.WriteString("Pokemon(")
	builder.WriteString(fmt.Sprintf("id=%v, ", po.ID))
	builder.WriteString("name=")
	builder.WriteString(po.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(po.Description)
	builder.WriteString(", ")
	builder.WriteString("weight=")
	builder.WriteString(fmt.Sprintf("%v", po.Weight))
	builder.WriteString(", ")
	builder.WriteString("height=")
	builder.WriteString(fmt.Sprintf("%v", po.Height))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(po.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(po.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Pokemons is a parsable slice of Pokemon.
type Pokemons []*Pokemon
