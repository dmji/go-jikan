package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetAnimeStreamingRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetAnimeStreamingRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetAnimeStreamingResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns anime streaming links
	ExternalLinks *components.ExternalLinks
}

func (o *GetAnimeStreamingResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAnimeStreamingResponse) GetExternalLinks() *components.ExternalLinks {
	if o == nil {
		return nil
	}
	return o.ExternalLinks
}
