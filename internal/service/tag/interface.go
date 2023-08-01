package tag

import (
	"context"

	"github.com/tupyy/folder-manager/internal/entity"
)

type TagRepository interface {
	// Create -- create the tag.
	Create(ctx context.Context, tag entity.Tag) (string, error)
	// Update -- update the tag.
	Update(ctx context.Context, tag entity.Tag) error
	// Delete -- delete the tag. it does not cascade.
	Delete(ctx context.Context, id string) error
	// GetByUser -- fetch all user's tags
	GetByUser(ctx context.Context, userID string) ([]entity.Tag, error)
	// GetByName -- fetch the tag by name and user id.
	GetByName(ctx context.Context, userID, name string) (entity.Tag, error)
	// GetByID -- fetch the tag by id
	GetByID(ctx context.Context, userID string, id string) (entity.Tag, error)
	// GetByAlbum -- fetch all user's tag for the album
	GetByAlbum(ctx context.Context, albumID string) ([]entity.Tag, error)
	// AssociateTag -- associates a tag with an album.
	Associate(ctx context.Context, albumID, tagID string) error
	// Dissociate -- removes a tag from an album.
	Dissociate(ctx context.Context, albumID, tagID string) error
}
