package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetPeopleSearchRequest struct {
	Page  *int64  `queryParam:"style=form,explode=true,name=page"`
	Limit *int64  `queryParam:"style=form,explode=true,name=limit"`
	Q     *string `queryParam:"style=form,explode=true,name=q"`
	// Available People order_by properties
	OrderBy *components.PeopleSearchQueryOrderby `queryParam:"style=form,explode=true,name=order_by"`
	// Search query sort direction
	Sort *components.SearchQuerySort `queryParam:"style=form,explode=true,name=sort"`
	// Return entries starting with the given letter
	Letter *string `queryParam:"style=form,explode=true,name=letter"`
}

func (o *GetPeopleSearchRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetPeopleSearchRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetPeopleSearchRequest) GetQ() *string {
	if o == nil {
		return nil
	}
	return o.Q
}

func (o *GetPeopleSearchRequest) GetOrderBy() *components.PeopleSearchQueryOrderby {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *GetPeopleSearchRequest) GetSort() *components.SearchQuerySort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetPeopleSearchRequest) GetLetter() *string {
	if o == nil {
		return nil
	}
	return o.Letter
}

type GetPeopleSearchResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns search results for people
	PeopleSearch *components.PeopleSearch
}

func (o *GetPeopleSearchResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetPeopleSearchResponse) GetPeopleSearch() *components.PeopleSearch {
	if o == nil {
		return nil
	}
	return o.PeopleSearch
}
