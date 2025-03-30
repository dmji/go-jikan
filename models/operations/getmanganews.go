package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetMangaNewsRequest struct {
	ID   int64  `pathParam:"style=simple,explode=false,name=id"`
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
}

func (o *GetMangaNewsRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *GetMangaNewsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

type GetMangaNewsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns a list of manga news topics
	MangaNews *components.MangaNews
}

func (o *GetMangaNewsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetMangaNewsResponse) GetMangaNews() *components.MangaNews {
	if o == nil {
		return nil
	}
	return o.MangaNews
}
