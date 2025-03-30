package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetRecentAnimeReviewsRequest struct {
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
	// Any reviews left during an ongoing anime/manga, those reviews are tagged as preliminary. NOTE: Preliminary reviews are not returned by default so if the entry is airing/publishing you need to add this otherwise you will get an empty list. e.g usage: `?preliminary=true`
	Preliminary *bool `queryParam:"style=form,explode=true,name=preliminary"`
	// Any reviews that are tagged as a spoiler. Spoiler reviews are not returned by default. e.g usage: `?spoiler=true`
	Spoilers *bool `queryParam:"style=form,explode=true,name=spoilers"`
}

func (o *GetRecentAnimeReviewsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetRecentAnimeReviewsRequest) GetPreliminary() *bool {
	if o == nil {
		return nil
	}
	return o.Preliminary
}

func (o *GetRecentAnimeReviewsRequest) GetSpoilers() *bool {
	if o == nil {
		return nil
	}
	return o.Spoilers
}

type GetRecentAnimeReviewsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns recent anime reviews
	Any any
}

func (o *GetRecentAnimeReviewsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetRecentAnimeReviewsResponse) GetAny() any {
	if o == nil {
		return nil
	}
	return o.Any
}
