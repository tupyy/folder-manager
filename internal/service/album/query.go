package album

import (
	"context"

	"github.com/tupyy/folder-manager/internal/entity"
	"github.com/tupyy/folder-manager/internal/service/media"
	"go.uber.org/zap"
)

type Filter interface {
	Resolve(album entity.Album) (bool, error)
}

type Query struct {
	size int
	page int
	// get personal albums.
	personalAlbums bool
	// get shared albums.
	sharedAlbums bool
	// filter
	filter Filter
	// album repo
	albumRepo AlbumRepository
	// media service
	mediaService *media.Service
	//album sorter
	sorter *albumSorter
}

func (s *Service) Query() *Query {
	return &Query{
		albumRepo:    s.albumRepo,
		mediaService: s.mediaService,
	}
}

func (q *Query) Filter(filter Filter) *Query {
	q.filter = filter

	return q
}

func (q *Query) Size(size int) *Query {
	q.size = size
	return q
}

func (q *Query) Page(page int) *Query {
	q.page = page

	return q
}

func (q *Query) OwnAlbums(b bool) *Query {
	q.personalAlbums = b

	return q
}

func (q *Query) Sort(name SortType, order SortOrder) *Query {
	as := newSorter(name, order)
	q.sorter = as

	return q
}

func (q *Query) SharedAlbums(b bool) *Query {
	q.sharedAlbums = b

	return q
}

// All returns a list of albums sliced if offset & limit are set and the total number of albums.
func (q *Query) All(ctx context.Context, user entity.User) ([]entity.Album, int, error) {
	albums := make(map[string]entity.Album)

	if q.personalAlbums {
		// fetch personal albums
		pa, err := q.albumRepo.GetByOwner(ctx, user.Username)
		if err != nil {
			return []entity.Album{}, 0, err
		}

		for _, a := range pa {
			albums[a.ID] = a
		}
	}

	if q.sharedAlbums {
		// if the user is an admin, get all albums regardless of permissions
		if user.Role == entity.RoleAdmin {
			sa, err := q.albumRepo.Get(ctx)
			if err != nil {
				return []entity.Album{}, 0, err
			}

			for _, a := range sa {
				albums[a.ID] = a
			}
		} else if user.CanShare {
			sharedAlbums, err := q.albumRepo.GetByUser(ctx, user.Username)
			if err != nil {
				return []entity.Album{}, 0, err
			}

			// get albums shared by the user's groups but filter out the ones owns by the user
			// groupSharedAlbum, err := q.albumRepo.GetByGroups(ctx, groupsToList(user.Groups))
			// if err != nil {
			// 	return []entity.Album{}, 0, fmt.Errorf("%w shared albums by group: %v", services.ErrGetAlbums, err)
			// }
			groupSharedAlbum := []entity.Album{}

			for i := 0; i < len(sharedAlbums)+len(groupSharedAlbum); i++ {
				found := false
				if i < len(sharedAlbums) {
					albums[sharedAlbums[i].ID] = sharedAlbums[i]
					found = true
				}

				if i < len(groupSharedAlbum) {
					albums[groupSharedAlbum[i].ID] = groupSharedAlbum[i]
					found = true
				}

				if !found {
					break
				}
			}
		}
	}

	// put all the albums into a list and return them
	albs := make([]entity.Album, 0, len(albums))
	for _, a := range albums {
		if q.filter != nil {
			resolved, err := q.filter.Resolve(a)
			if err != nil {
				zap.S().Errorw("failed to resolve album", "error", err, "album", a)
				continue
			}

			if resolved {
				albs = append(albs, a)
			}
		} else {
			albs = append(albs, a)
		}
	}

	if q.sorter != nil {
		q.sorter.Sort(albs)
	}

	return q.paginate(albs), len(albs), nil
}

func (q *Query) First(ctx context.Context, id string) (entity.Album, error) {
	album, err := q.albumRepo.GetByID(ctx, id)
	if err != nil {
		return entity.Album{}, err
	}

	medias, err := q.mediaService.ListBucket(ctx, album.Bucket)
	if err != nil {
		return entity.Album{}, err
	}

	photos := make([]entity.Media, 0, len(medias))
	videos := make([]entity.Media, 0, len(medias))

	for _, m := range medias {
		switch m.MediaType {
		case entity.Photo:
			photos = append(photos, m)
		case entity.Video:
			videos = append(videos, m)
		}
	}

	album.Photos = photos
	album.Videos = videos

	return album, nil
}

func (q *Query) paginate(albums []entity.Album) []entity.Album {
	// pagination
	var page []entity.Album

	if q.page <= 0 || q.size <= 0 {
		return albums
	}

	offset := (q.page - 1) * q.size
	limit := offset + q.size

	if limit >= len(albums) {
		limit = len(albums)
	}

	if offset >= len(albums) {
		return []entity.Album{}
	}

	page = append(page, albums[offset:limit]...)

	return page
}

func groupsToList(groups []entity.Group) []string {
	l := make([]string, 0, len(groups))

	for _, g := range groups {
		l = append(l, g.Name)
	}

	return l
}
