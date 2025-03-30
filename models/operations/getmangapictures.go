package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetMangaPicturesRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetMangaPicturesRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetMangaPicturesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns a list of manga pictures
	MangaPictures *components.MangaPictures
}

func (o *GetMangaPicturesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetMangaPicturesResponse) GetMangaPictures() *components.MangaPictures {
	if o == nil {
		return nil
	}
	return o.MangaPictures
}
