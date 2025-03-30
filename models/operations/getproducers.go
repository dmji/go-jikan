package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetProducersRequest struct {
	Page  *int64  `queryParam:"style=form,explode=true,name=page"`
	Limit *int64  `queryParam:"style=form,explode=true,name=limit"`
	Q     *string `queryParam:"style=form,explode=true,name=q"`
	// Producers Search Query Order By
	OrderBy *components.ProducersQueryOrderby `queryParam:"style=form,explode=true,name=order_by"`
	// Search query sort direction
	Sort *components.SearchQuerySort `queryParam:"style=form,explode=true,name=sort"`
	// Return entries starting with the given letter
	Letter *string `queryParam:"style=form,explode=true,name=letter"`
}

func (o *GetProducersRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetProducersRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetProducersRequest) GetQ() *string {
	if o == nil {
		return nil
	}
	return o.Q
}

func (o *GetProducersRequest) GetOrderBy() *components.ProducersQueryOrderby {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *GetProducersRequest) GetSort() *components.SearchQuerySort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetProducersRequest) GetLetter() *string {
	if o == nil {
		return nil
	}
	return o.Letter
}

type GetProducersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns producers collection
	Producers *components.Producers
}

func (o *GetProducersResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetProducersResponse) GetProducers() *components.Producers {
	if o == nil {
		return nil
	}
	return o.Producers
}
