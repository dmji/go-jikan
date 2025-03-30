package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetProducerExternalRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetProducerExternalRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetProducerExternalResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns producer's external links
	ExternalLinks *components.ExternalLinks
}

func (o *GetProducerExternalResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetProducerExternalResponse) GetExternalLinks() *components.ExternalLinks {
	if o == nil {
		return nil
	}
	return o.ExternalLinks
}
