package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetMangaReviewsRequest struct {
	ID   int64  `pathParam:"style=simple,explode=false,name=id"`
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
	// Any reviews left during an ongoing anime/manga, those reviews are tagged as preliminary. NOTE: Preliminary reviews are not returned by default so if the entry is airing/publishing you need to add this otherwise you will get an empty list. e.g usage: `?preliminary=true`
	Preliminary *bool `queryParam:"style=form,explode=true,name=preliminary"`
	// Any reviews that are tagged as a spoiler. Spoiler reviews are not returned by default. e.g usage: `?spoiler=true`
	Spoilers *bool `queryParam:"style=form,explode=true,name=spoilers"`
}

func (o *GetMangaReviewsRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *GetMangaReviewsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetMangaReviewsRequest) GetPreliminary() *bool {
	if o == nil {
		return nil
	}
	return o.Preliminary
}

func (o *GetMangaReviewsRequest) GetSpoilers() *bool {
	if o == nil {
		return nil
	}
	return o.Spoilers
}

type GetMangaReviewsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns manga reviews
	MangaReviews *components.MangaReviews
}

func (o *GetMangaReviewsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetMangaReviewsResponse) GetMangaReviews() *components.MangaReviews {
	if o == nil {
		return nil
	}
	return o.MangaReviews
}
