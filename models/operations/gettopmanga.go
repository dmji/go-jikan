package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetTopMangaRequest struct {
	// Available Manga types
	Type *components.MangaSearchQueryType `queryParam:"style=form,explode=true,name=type"`
	// Top items filter types
	Filter *components.TopMangaFilter `queryParam:"style=form,explode=true,name=filter"`
	Page   *int64                     `queryParam:"style=form,explode=true,name=page"`
	Limit  *int64                     `queryParam:"style=form,explode=true,name=limit"`
}

func (o *GetTopMangaRequest) GetType() *components.MangaSearchQueryType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *GetTopMangaRequest) GetFilter() *components.TopMangaFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *GetTopMangaRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetTopMangaRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

type GetTopMangaResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns top manga
	MangaSearch *components.MangaSearch
}

func (o *GetTopMangaResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetTopMangaResponse) GetMangaSearch() *components.MangaSearch {
	if o == nil {
		return nil
	}
	return o.MangaSearch
}
