package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetAnimeUserUpdatesRequest struct {
	ID   int64  `pathParam:"style=simple,explode=false,name=id"`
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
}

func (o *GetAnimeUserUpdatesRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *GetAnimeUserUpdatesRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

type GetAnimeUserUpdatesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns a list of users who have added/updated/removed the entry on their list
	AnimeUserupdates *components.AnimeUserupdates
}

func (o *GetAnimeUserUpdatesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAnimeUserUpdatesResponse) GetAnimeUserupdates() *components.AnimeUserupdates {
	if o == nil {
		return nil
	}
	return o.AnimeUserupdates
}
