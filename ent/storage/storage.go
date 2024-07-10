// Code generated by ent, DO NOT EDIT.

package storage

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the storage type in the database.
	Label = "storage"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldImplementation holds the string denoting the implementation field in the database.
	FieldImplementation = "implementation"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldUser holds the string denoting the user field in the database.
	FieldUser = "user"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldDeleted holds the string denoting the deleted field in the database.
	FieldDeleted = "deleted"
	// FieldDefault holds the string denoting the default field in the database.
	FieldDefault = "default"
	// Table holds the table name of the storage in the database.
	Table = "storages"
)

// Columns holds all SQL columns for storage fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldImplementation,
	FieldPath,
	FieldUser,
	FieldPassword,
	FieldDeleted,
	FieldDefault,
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

var (
	// DefaultDeleted holds the default value on creation for the "deleted" field.
	DefaultDeleted bool
	// DefaultDefault holds the default value on creation for the "default" field.
	DefaultDefault bool
)

// OrderOption defines the ordering options for the Storage queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByImplementation orders the results by the implementation field.
func ByImplementation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImplementation, opts...).ToFunc()
}

// ByPath orders the results by the path field.
func ByPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPath, opts...).ToFunc()
}

// ByUser orders the results by the user field.
func ByUser(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUser, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByDeleted orders the results by the deleted field.
func ByDeleted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeleted, opts...).ToFunc()
}

// ByDefault orders the results by the default field.
func ByDefault(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDefault, opts...).ToFunc()
}
