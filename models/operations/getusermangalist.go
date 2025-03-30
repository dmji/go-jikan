package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetUserMangaListRequest struct {
	Username string `pathParam:"style=simple,explode=false,name=username"`
	// User's anime list status filter options
	Status *components.UserMangaListStatusFilter `queryParam:"style=form,explode=true,name=status"`
}

func (o *GetUserMangaListRequest) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *GetUserMangaListRequest) GetStatus() *components.UserMangaListStatusFilter {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetUserMangaListResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns user manga list
	Any any
}

func (o *GetUserMangaListResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetUserMangaListResponse) GetAny() any {
	if o == nil {
		return nil
	}
	return o.Any
}
