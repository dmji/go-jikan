package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetMagazinesRequest struct {
	Page  *int64  `queryParam:"style=form,explode=true,name=page"`
	Limit *int64  `queryParam:"style=form,explode=true,name=limit"`
	Q     *string `queryParam:"style=form,explode=true,name=q"`
	// Order by magazine data
	OrderBy *components.MagazinesQueryOrderby `queryParam:"style=form,explode=true,name=order_by"`
	// Search query sort direction
	Sort *components.SearchQuerySort `queryParam:"style=form,explode=true,name=sort"`
	// Return entries starting with the given letter
	Letter *string `queryParam:"style=form,explode=true,name=letter"`
}

func (o *GetMagazinesRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetMagazinesRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetMagazinesRequest) GetQ() *string {
	if o == nil {
		return nil
	}
	return o.Q
}

func (o *GetMagazinesRequest) GetOrderBy() *components.MagazinesQueryOrderby {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *GetMagazinesRequest) GetSort() *components.SearchQuerySort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetMagazinesRequest) GetLetter() *string {
	if o == nil {
		return nil
	}
	return o.Letter
}

type GetMagazinesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns magazines collection
	Magazines *components.Magazines
}

func (o *GetMagazinesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetMagazinesResponse) GetMagazines() *components.Magazines {
	if o == nil {
		return nil
	}
	return o.Magazines
}
