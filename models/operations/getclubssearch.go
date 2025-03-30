package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetClubsSearchRequest struct {
	Page  *int64  `queryParam:"style=form,explode=true,name=page"`
	Limit *int64  `queryParam:"style=form,explode=true,name=limit"`
	Q     *string `queryParam:"style=form,explode=true,name=q"`
	// Club Search Query Type
	Type *components.ClubSearchQueryType `queryParam:"style=form,explode=true,name=type"`
	// Club Search Query Category
	Category *components.ClubSearchQueryCategory `queryParam:"style=form,explode=true,name=category"`
	// Club Search Query OrderBy
	OrderBy *components.ClubSearchQueryOrderby `queryParam:"style=form,explode=true,name=order_by"`
	// Search query sort direction
	Sort *components.SearchQuerySort `queryParam:"style=form,explode=true,name=sort"`
	// Return entries starting with the given letter
	Letter *string `queryParam:"style=form,explode=true,name=letter"`
}

func (o *GetClubsSearchRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetClubsSearchRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetClubsSearchRequest) GetQ() *string {
	if o == nil {
		return nil
	}
	return o.Q
}

func (o *GetClubsSearchRequest) GetType() *components.ClubSearchQueryType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *GetClubsSearchRequest) GetCategory() *components.ClubSearchQueryCategory {
	if o == nil {
		return nil
	}
	return o.Category
}

func (o *GetClubsSearchRequest) GetOrderBy() *components.ClubSearchQueryOrderby {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *GetClubsSearchRequest) GetSort() *components.SearchQuerySort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetClubsSearchRequest) GetLetter() *string {
	if o == nil {
		return nil
	}
	return o.Letter
}

type GetClubsSearchResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns search results for clubs
	ClubsSearch *components.ClubsSearch
}

func (o *GetClubsSearchResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetClubsSearchResponse) GetClubsSearch() *components.ClubsSearch {
	if o == nil {
		return nil
	}
	return o.ClubsSearch
}
