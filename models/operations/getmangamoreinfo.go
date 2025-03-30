package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetMangaMoreInfoRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetMangaMoreInfoRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetMangaMoreInfoResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns manga moreinfo
	Moreinfo *components.Moreinfo
}

func (o *GetMangaMoreInfoResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetMangaMoreInfoResponse) GetMoreinfo() *components.Moreinfo {
	if o == nil {
		return nil
	}
	return o.Moreinfo
}
