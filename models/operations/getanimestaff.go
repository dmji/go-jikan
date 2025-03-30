package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetAnimeStaffRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetAnimeStaffRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetAnimeStaffResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns anime staff resource
	AnimeStaff *components.AnimeStaff
}

func (o *GetAnimeStaffResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAnimeStaffResponse) GetAnimeStaff() *components.AnimeStaff {
	if o == nil {
		return nil
	}
	return o.AnimeStaff
}
