package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetPersonPicturesRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetPersonPicturesRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetPersonPicturesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns a list of pictures of the person
	PersonPictures *components.PersonPictures
}

func (o *GetPersonPicturesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetPersonPicturesResponse) GetPersonPictures() *components.PersonPictures {
	if o == nil {
		return nil
	}
	return o.PersonPictures
}
