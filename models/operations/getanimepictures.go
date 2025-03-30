package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetAnimePicturesRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetAnimePicturesRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetAnimePicturesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns pictures related to the entry
	PicturesVariants *components.PicturesVariants
}

func (o *GetAnimePicturesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAnimePicturesResponse) GetPicturesVariants() *components.PicturesVariants {
	if o == nil {
		return nil
	}
	return o.PicturesVariants
}
