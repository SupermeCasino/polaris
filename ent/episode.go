// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"polaris/ent/episode"
	"polaris/ent/series"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Episode is the model entity for the Episode schema.
type Episode struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SeriesID holds the value of the "series_id" field.
	SeriesID int `json:"series_id,omitempty"`
	// SeasonNumber holds the value of the "season_number" field.
	SeasonNumber int `json:"season_number"`
	// EpisodeNumber holds the value of the "episode_number" field.
	EpisodeNumber int `json:"episode_number"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Overview holds the value of the "overview" field.
	Overview string `json:"overview,omitempty"`
	// AirDate holds the value of the "air_date" field.
	AirDate string `json:"air_date,omitempty"`
	// FileInStorage holds the value of the "file_in_storage" field.
	FileInStorage string `json:"file_in_storage,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EpisodeQuery when eager-loading is set.
	Edges        EpisodeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// EpisodeEdges holds the relations/edges for other nodes in the graph.
type EpisodeEdges struct {
	// Series holds the value of the series edge.
	Series *Series `json:"series,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SeriesOrErr returns the Series value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EpisodeEdges) SeriesOrErr() (*Series, error) {
	if e.Series != nil {
		return e.Series, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: series.Label}
	}
	return nil, &NotLoadedError{edge: "series"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Episode) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case episode.FieldID, episode.FieldSeriesID, episode.FieldSeasonNumber, episode.FieldEpisodeNumber:
			values[i] = new(sql.NullInt64)
		case episode.FieldTitle, episode.FieldOverview, episode.FieldAirDate, episode.FieldFileInStorage:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Episode fields.
func (e *Episode) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case episode.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case episode.FieldSeriesID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field series_id", values[i])
			} else if value.Valid {
				e.SeriesID = int(value.Int64)
			}
		case episode.FieldSeasonNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field season_number", values[i])
			} else if value.Valid {
				e.SeasonNumber = int(value.Int64)
			}
		case episode.FieldEpisodeNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field episode_number", values[i])
			} else if value.Valid {
				e.EpisodeNumber = int(value.Int64)
			}
		case episode.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				e.Title = value.String
			}
		case episode.FieldOverview:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field overview", values[i])
			} else if value.Valid {
				e.Overview = value.String
			}
		case episode.FieldAirDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field air_date", values[i])
			} else if value.Valid {
				e.AirDate = value.String
			}
		case episode.FieldFileInStorage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field file_in_storage", values[i])
			} else if value.Valid {
				e.FileInStorage = value.String
			}
		default:
			e.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Episode.
// This includes values selected through modifiers, order, etc.
func (e *Episode) Value(name string) (ent.Value, error) {
	return e.selectValues.Get(name)
}

// QuerySeries queries the "series" edge of the Episode entity.
func (e *Episode) QuerySeries() *SeriesQuery {
	return NewEpisodeClient(e.config).QuerySeries(e)
}

// Update returns a builder for updating this Episode.
// Note that you need to call Episode.Unwrap() before calling this method if this Episode
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Episode) Update() *EpisodeUpdateOne {
	return NewEpisodeClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Episode entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Episode) Unwrap() *Episode {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Episode is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Episode) String() string {
	var builder strings.Builder
	builder.WriteString("Episode(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("series_id=")
	builder.WriteString(fmt.Sprintf("%v", e.SeriesID))
	builder.WriteString(", ")
	builder.WriteString("season_number=")
	builder.WriteString(fmt.Sprintf("%v", e.SeasonNumber))
	builder.WriteString(", ")
	builder.WriteString("episode_number=")
	builder.WriteString(fmt.Sprintf("%v", e.EpisodeNumber))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(e.Title)
	builder.WriteString(", ")
	builder.WriteString("overview=")
	builder.WriteString(e.Overview)
	builder.WriteString(", ")
	builder.WriteString("air_date=")
	builder.WriteString(e.AirDate)
	builder.WriteString(", ")
	builder.WriteString("file_in_storage=")
	builder.WriteString(e.FileInStorage)
	builder.WriteByte(')')
	return builder.String()
}

// Episodes is a parsable slice of Episode.
type Episodes []*Episode
