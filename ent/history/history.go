// Code generated by ent, DO NOT EDIT.

package history

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the history type in the database.
	Label = "history"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSeriesID holds the string denoting the series_id field in the database.
	FieldSeriesID = "series_id"
	// FieldEpisodeID holds the string denoting the episode_id field in the database.
	FieldEpisodeID = "episode_id"
	// FieldSourceTitle holds the string denoting the source_title field in the database.
	FieldSourceTitle = "source_title"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldTargetDir holds the string denoting the target_dir field in the database.
	FieldTargetDir = "target_dir"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldSaved holds the string denoting the saved field in the database.
	FieldSaved = "saved"
	// Table holds the table name of the history in the database.
	Table = "histories"
)

// Columns holds all SQL columns for history fields.
var Columns = []string{
	FieldID,
	FieldSeriesID,
	FieldEpisodeID,
	FieldSourceTitle,
	FieldDate,
	FieldTargetDir,
	FieldStatus,
	FieldSaved,
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

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusRunning   Status = "running"
	StatusSuccess   Status = "success"
	StatusFail      Status = "fail"
	StatusUploading Status = "uploading"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusRunning, StatusSuccess, StatusFail, StatusUploading:
		return nil
	default:
		return fmt.Errorf("history: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the History queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySeriesID orders the results by the series_id field.
func BySeriesID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSeriesID, opts...).ToFunc()
}

// ByEpisodeID orders the results by the episode_id field.
func ByEpisodeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEpisodeID, opts...).ToFunc()
}

// BySourceTitle orders the results by the source_title field.
func BySourceTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSourceTitle, opts...).ToFunc()
}

// ByDate orders the results by the date field.
func ByDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDate, opts...).ToFunc()
}

// ByTargetDir orders the results by the target_dir field.
func ByTargetDir(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTargetDir, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// BySaved orders the results by the saved field.
func BySaved(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSaved, opts...).ToFunc()
}
