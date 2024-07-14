// Code generated by ent, DO NOT EDIT.

package episode

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the episode type in the database.
	Label = "episode"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSeriesID holds the string denoting the series_id field in the database.
	FieldSeriesID = "series_id"
	// FieldSeasonNumber holds the string denoting the season_number field in the database.
	FieldSeasonNumber = "season_number"
	// FieldEpisodeNumber holds the string denoting the episode_number field in the database.
	FieldEpisodeNumber = "episode_number"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldOverview holds the string denoting the overview field in the database.
	FieldOverview = "overview"
	// FieldAirDate holds the string denoting the air_date field in the database.
	FieldAirDate = "air_date"
	// FieldFileInStorage holds the string denoting the file_in_storage field in the database.
	FieldFileInStorage = "file_in_storage"
	// EdgeSeries holds the string denoting the series edge name in mutations.
	EdgeSeries = "series"
	// Table holds the table name of the episode in the database.
	Table = "episodes"
	// SeriesTable is the table that holds the series relation/edge.
	SeriesTable = "episodes"
	// SeriesInverseTable is the table name for the Series entity.
	// It exists in this package in order to avoid circular dependency with the "series" package.
	SeriesInverseTable = "series"
	// SeriesColumn is the table column denoting the series relation/edge.
	SeriesColumn = "series_id"
)

// Columns holds all SQL columns for episode fields.
var Columns = []string{
	FieldID,
	FieldSeriesID,
	FieldSeasonNumber,
	FieldEpisodeNumber,
	FieldTitle,
	FieldOverview,
	FieldAirDate,
	FieldFileInStorage,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Episode queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySeriesID orders the results by the series_id field.
func BySeriesID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSeriesID, opts...).ToFunc()
}

// BySeasonNumber orders the results by the season_number field.
func BySeasonNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSeasonNumber, opts...).ToFunc()
}

// ByEpisodeNumber orders the results by the episode_number field.
func ByEpisodeNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEpisodeNumber, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByOverview orders the results by the overview field.
func ByOverview(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOverview, opts...).ToFunc()
}

// ByAirDate orders the results by the air_date field.
func ByAirDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAirDate, opts...).ToFunc()
}

// ByFileInStorage orders the results by the file_in_storage field.
func ByFileInStorage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFileInStorage, opts...).ToFunc()
}

// BySeriesField orders the results by series field.
func BySeriesField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSeriesStep(), sql.OrderByField(field, opts...))
	}
}
func newSeriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SeriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SeriesTable, SeriesColumn),
	)
}
