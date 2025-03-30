package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetAnimeExternalRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetAnimeExternalRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetAnimeExternalResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns anime external links
	ExternalLinks *components.ExternalLinks
}

func (o *GetAnimeExternalResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAnimeExternalResponse) GetExternalLinks() *components.ExternalLinks {
	if o == nil {
		return nil
	}
	return o.ExternalLinks
}
