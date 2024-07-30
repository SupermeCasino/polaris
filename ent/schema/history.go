package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// History holds the schema definition for the History entity.
type History struct {
	ent.Schema
}

// Fields of the History.
func (History) Fields() []ent.Field {
	return []ent.Field{
		field.Int("media_id"),
		field.Int("episode_id").Optional(),
		field.String("source_title"),
		field.Time("date"),
		field.String("target_dir"),
		field.Int("size").Default(0),
		field.Int("download_client_id").Optional(),
		field.Enum("status").Values("running", "success", "fail", "uploading"),
		field.String("saved").Optional(),
	}
}

// Edges of the History.
func (History) Edges() []ent.Edge {
	return nil
}
