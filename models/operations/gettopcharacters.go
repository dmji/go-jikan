package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetTopCharactersRequest struct {
	Page  *int64 `queryParam:"style=form,explode=true,name=page"`
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
}

func (o *GetTopCharactersRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetTopCharactersRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

type GetTopCharactersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns top characters
	CharactersSearch *components.CharactersSearch
}

func (o *GetTopCharactersResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetTopCharactersResponse) GetCharactersSearch() *components.CharactersSearch {
	if o == nil {
		return nil
	}
	return o.CharactersSearch
}
