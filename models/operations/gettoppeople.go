package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetTopPeopleRequest struct {
	Page  *int64 `queryParam:"style=form,explode=true,name=page"`
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
}

func (o *GetTopPeopleRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetTopPeopleRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

type GetTopPeopleResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns top people
	PeopleSearch *components.PeopleSearch
}

func (o *GetTopPeopleResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetTopPeopleResponse) GetPeopleSearch() *components.PeopleSearch {
	if o == nil {
		return nil
	}
	return o.PeopleSearch
}
