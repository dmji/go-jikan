package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetCharactersSearchRequest struct {
	Page  *int64  `queryParam:"style=form,explode=true,name=page"`
	Limit *int64  `queryParam:"style=form,explode=true,name=limit"`
	Q     *string `queryParam:"style=form,explode=true,name=q"`
	// Available Character order_by properties
	OrderBy *components.CharactersSearchQueryOrderby `queryParam:"style=form,explode=true,name=order_by"`
	// Search query sort direction
	Sort *components.SearchQuerySort `queryParam:"style=form,explode=true,name=sort"`
	// Return entries starting with the given letter
	Letter *string `queryParam:"style=form,explode=true,name=letter"`
}

func (o *GetCharactersSearchRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetCharactersSearchRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetCharactersSearchRequest) GetQ() *string {
	if o == nil {
		return nil
	}
	return o.Q
}

func (o *GetCharactersSearchRequest) GetOrderBy() *components.CharactersSearchQueryOrderby {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *GetCharactersSearchRequest) GetSort() *components.SearchQuerySort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetCharactersSearchRequest) GetLetter() *string {
	if o == nil {
		return nil
	}
	return o.Letter
}

type GetCharactersSearchResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns search results for characters
	CharactersSearch *components.CharactersSearch
}

func (o *GetCharactersSearchResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetCharactersSearchResponse) GetCharactersSearch() *components.CharactersSearch {
	if o == nil {
		return nil
	}
	return o.CharactersSearch
}
