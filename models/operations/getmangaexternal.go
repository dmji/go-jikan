package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetMangaExternalRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetMangaExternalRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetMangaExternalResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns manga external links
	ExternalLinks *components.ExternalLinks
}

func (o *GetMangaExternalResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetMangaExternalResponse) GetExternalLinks() *components.ExternalLinks {
	if o == nil {
		return nil
	}
	return o.ExternalLinks
}
