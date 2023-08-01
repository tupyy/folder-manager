package entity

import (
	"time"
)

type ObjectMeta struct {
	// ID - id of the object
	ID string
	// Name - name of the object
	Name string
	// CreateAt - creation date
	CreatedAt time.Time
	// Labels
	Labels map[string]string
}

type FolderMeta struct {
	// ResourceID -- id of the resource from the resource server
	ResourceID string
}

type Folder struct {
	ObjectMeta
	FolderMeta
	// Owner - owner of the folder
	Owner User
	// Bucket - name of bucket in the store
	Bucket string
	// Files - list of files
	Files []File
	// Description -- description
	Description string
}
