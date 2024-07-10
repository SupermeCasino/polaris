// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DownloadClientsColumns holds the columns for the "download_clients" table.
	DownloadClientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "enable", Type: field.TypeBool},
		{Name: "name", Type: field.TypeString},
		{Name: "implementation", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
		{Name: "user", Type: field.TypeString, Default: ""},
		{Name: "password", Type: field.TypeString, Default: ""},
		{Name: "settings", Type: field.TypeString, Default: ""},
		{Name: "priority", Type: field.TypeString, Default: ""},
		{Name: "remove_completed_downloads", Type: field.TypeBool, Default: true},
		{Name: "remove_failed_downloads", Type: field.TypeBool, Default: true},
		{Name: "tags", Type: field.TypeString, Default: ""},
	}
	// DownloadClientsTable holds the schema information for the "download_clients" table.
	DownloadClientsTable = &schema.Table{
		Name:       "download_clients",
		Columns:    DownloadClientsColumns,
		PrimaryKey: []*schema.Column{DownloadClientsColumns[0]},
	}
	// EpisodesColumns holds the columns for the "episodes" table.
	EpisodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "season_number", Type: field.TypeInt},
		{Name: "episode_number", Type: field.TypeInt},
		{Name: "title", Type: field.TypeString},
		{Name: "overview", Type: field.TypeString},
		{Name: "air_date", Type: field.TypeString},
		{Name: "series_id", Type: field.TypeInt, Nullable: true},
	}
	// EpisodesTable holds the schema information for the "episodes" table.
	EpisodesTable = &schema.Table{
		Name:       "episodes",
		Columns:    EpisodesColumns,
		PrimaryKey: []*schema.Column{EpisodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "episodes_series_episodes",
				Columns:    []*schema.Column{EpisodesColumns[6]},
				RefColumns: []*schema.Column{SeriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// HistoriesColumns holds the columns for the "histories" table.
	HistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "series_id", Type: field.TypeInt},
		{Name: "episode_id", Type: field.TypeInt},
		{Name: "source_title", Type: field.TypeString},
		{Name: "date", Type: field.TypeTime},
		{Name: "target_dir", Type: field.TypeString},
		{Name: "completed", Type: field.TypeBool, Default: false},
		{Name: "saved", Type: field.TypeString, Nullable: true},
	}
	// HistoriesTable holds the schema information for the "histories" table.
	HistoriesTable = &schema.Table{
		Name:       "histories",
		Columns:    HistoriesColumns,
		PrimaryKey: []*schema.Column{HistoriesColumns[0]},
	}
	// IndexersColumns holds the columns for the "indexers" table.
	IndexersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "implementation", Type: field.TypeString},
		{Name: "settings", Type: field.TypeString},
		{Name: "enable_rss", Type: field.TypeBool, Default: true},
		{Name: "priority", Type: field.TypeInt},
	}
	// IndexersTable holds the schema information for the "indexers" table.
	IndexersTable = &schema.Table{
		Name:       "indexers",
		Columns:    IndexersColumns,
		PrimaryKey: []*schema.Column{IndexersColumns[0]},
	}
	// SeriesColumns holds the columns for the "series" table.
	SeriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "tmdb_id", Type: field.TypeInt},
		{Name: "imdb_id", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "name_en", Type: field.TypeString, Nullable: true},
		{Name: "original_name", Type: field.TypeString},
		{Name: "overview", Type: field.TypeString},
		{Name: "poster_path", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "air_date", Type: field.TypeString, Default: ""},
		{Name: "resolution", Type: field.TypeString, Default: ""},
		{Name: "storage_id", Type: field.TypeInt, Nullable: true},
	}
	// SeriesTable holds the schema information for the "series" table.
	SeriesTable = &schema.Table{
		Name:       "series",
		Columns:    SeriesColumns,
		PrimaryKey: []*schema.Column{SeriesColumns[0]},
	}
	// SettingsColumns holds the columns for the "settings" table.
	SettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "key", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
	}
	// SettingsTable holds the schema information for the "settings" table.
	SettingsTable = &schema.Table{
		Name:       "settings",
		Columns:    SettingsColumns,
		PrimaryKey: []*schema.Column{SettingsColumns[0]},
	}
	// StoragesColumns holds the columns for the "storages" table.
	StoragesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "implementation", Type: field.TypeString},
		{Name: "path", Type: field.TypeString},
		{Name: "user", Type: field.TypeString, Nullable: true},
		{Name: "password", Type: field.TypeString, Nullable: true},
		{Name: "deleted", Type: field.TypeBool, Default: false},
		{Name: "default", Type: field.TypeBool, Default: false},
	}
	// StoragesTable holds the schema information for the "storages" table.
	StoragesTable = &schema.Table{
		Name:       "storages",
		Columns:    StoragesColumns,
		PrimaryKey: []*schema.Column{StoragesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DownloadClientsTable,
		EpisodesTable,
		HistoriesTable,
		IndexersTable,
		SeriesTable,
		SettingsTable,
		StoragesTable,
	}
)

func init() {
	EpisodesTable.ForeignKeys[0].RefTable = SeriesTable
}
