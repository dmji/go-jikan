package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetAnimeMoreInfoRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetAnimeMoreInfoRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetAnimeMoreInfoResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns anime statistics
	Moreinfo *components.Moreinfo
}

func (o *GetAnimeMoreInfoResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAnimeMoreInfoResponse) GetMoreinfo() *components.Moreinfo {
	if o == nil {
		return nil
	}
	return o.Moreinfo
}
