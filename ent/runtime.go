// Code generated by ent, DO NOT EDIT.

package ent

import (
	"polaris/ent/downloadclients"
	"polaris/ent/history"
	"polaris/ent/indexers"
	"polaris/ent/schema"
	"polaris/ent/series"
	"polaris/ent/storage"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	downloadclientsFields := schema.DownloadClients{}.Fields()
	_ = downloadclientsFields
	// downloadclientsDescUser is the schema descriptor for user field.
	downloadclientsDescUser := downloadclientsFields[4].Descriptor()
	// downloadclients.DefaultUser holds the default value on creation for the user field.
	downloadclients.DefaultUser = downloadclientsDescUser.Default.(string)
	// downloadclientsDescPassword is the schema descriptor for password field.
	downloadclientsDescPassword := downloadclientsFields[5].Descriptor()
	// downloadclients.DefaultPassword holds the default value on creation for the password field.
	downloadclients.DefaultPassword = downloadclientsDescPassword.Default.(string)
	// downloadclientsDescSettings is the schema descriptor for settings field.
	downloadclientsDescSettings := downloadclientsFields[6].Descriptor()
	// downloadclients.DefaultSettings holds the default value on creation for the settings field.
	downloadclients.DefaultSettings = downloadclientsDescSettings.Default.(string)
	// downloadclientsDescPriority is the schema descriptor for priority field.
	downloadclientsDescPriority := downloadclientsFields[7].Descriptor()
	// downloadclients.DefaultPriority holds the default value on creation for the priority field.
	downloadclients.DefaultPriority = downloadclientsDescPriority.Default.(string)
	// downloadclientsDescRemoveCompletedDownloads is the schema descriptor for remove_completed_downloads field.
	downloadclientsDescRemoveCompletedDownloads := downloadclientsFields[8].Descriptor()
	// downloadclients.DefaultRemoveCompletedDownloads holds the default value on creation for the remove_completed_downloads field.
	downloadclients.DefaultRemoveCompletedDownloads = downloadclientsDescRemoveCompletedDownloads.Default.(bool)
	// downloadclientsDescRemoveFailedDownloads is the schema descriptor for remove_failed_downloads field.
	downloadclientsDescRemoveFailedDownloads := downloadclientsFields[9].Descriptor()
	// downloadclients.DefaultRemoveFailedDownloads holds the default value on creation for the remove_failed_downloads field.
	downloadclients.DefaultRemoveFailedDownloads = downloadclientsDescRemoveFailedDownloads.Default.(bool)
	// downloadclientsDescTags is the schema descriptor for tags field.
	downloadclientsDescTags := downloadclientsFields[10].Descriptor()
	// downloadclients.DefaultTags holds the default value on creation for the tags field.
	downloadclients.DefaultTags = downloadclientsDescTags.Default.(string)
	historyFields := schema.History{}.Fields()
	_ = historyFields
	// historyDescCompleted is the schema descriptor for completed field.
	historyDescCompleted := historyFields[5].Descriptor()
	// history.DefaultCompleted holds the default value on creation for the completed field.
	history.DefaultCompleted = historyDescCompleted.Default.(bool)
	indexersFields := schema.Indexers{}.Fields()
	_ = indexersFields
	// indexersDescEnableRss is the schema descriptor for enable_rss field.
	indexersDescEnableRss := indexersFields[3].Descriptor()
	// indexers.DefaultEnableRss holds the default value on creation for the enable_rss field.
	indexers.DefaultEnableRss = indexersDescEnableRss.Default.(bool)
	seriesFields := schema.Series{}.Fields()
	_ = seriesFields
	// seriesDescCreatedAt is the schema descriptor for created_at field.
	seriesDescCreatedAt := seriesFields[7].Descriptor()
	// series.DefaultCreatedAt holds the default value on creation for the created_at field.
	series.DefaultCreatedAt = seriesDescCreatedAt.Default.(time.Time)
	// seriesDescAirDate is the schema descriptor for air_date field.
	seriesDescAirDate := seriesFields[8].Descriptor()
	// series.DefaultAirDate holds the default value on creation for the air_date field.
	series.DefaultAirDate = seriesDescAirDate.Default.(string)
	// seriesDescResolution is the schema descriptor for resolution field.
	seriesDescResolution := seriesFields[9].Descriptor()
	// series.DefaultResolution holds the default value on creation for the resolution field.
	series.DefaultResolution = seriesDescResolution.Default.(string)
	storageFields := schema.Storage{}.Fields()
	_ = storageFields
	// storageDescDeleted is the schema descriptor for deleted field.
	storageDescDeleted := storageFields[5].Descriptor()
	// storage.DefaultDeleted holds the default value on creation for the deleted field.
	storage.DefaultDeleted = storageDescDeleted.Default.(bool)
	// storageDescDefault is the schema descriptor for default field.
	storageDescDefault := storageFields[6].Descriptor()
	// storage.DefaultDefault holds the default value on creation for the default field.
	storage.DefaultDefault = storageDescDefault.Default.(bool)
}
