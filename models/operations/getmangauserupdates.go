package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetMangaUserUpdatesRequest struct {
	ID   int64  `pathParam:"style=simple,explode=false,name=id"`
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
}

func (o *GetMangaUserUpdatesRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *GetMangaUserUpdatesRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

type GetMangaUserUpdatesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns manga user updates
	MangaUserupdates *components.MangaUserupdates
}

func (o *GetMangaUserUpdatesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetMangaUserUpdatesResponse) GetMangaUserupdates() *components.MangaUserupdates {
	if o == nil {
		return nil
	}
	return o.MangaUserupdates
}
