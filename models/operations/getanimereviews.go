package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetAnimeReviewsRequest struct {
	ID   int64  `pathParam:"style=simple,explode=false,name=id"`
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
	// Any reviews left during an ongoing anime/manga, those reviews are tagged as preliminary. NOTE: Preliminary reviews are not returned by default so if the entry is airing/publishing you need to add this otherwise you will get an empty list. e.g usage: `?preliminary=true`
	Preliminary *bool `queryParam:"style=form,explode=true,name=preliminary"`
	// Any reviews that are tagged as a spoiler. Spoiler reviews are not returned by default. e.g usage: `?spoiler=true`
	Spoilers *bool `queryParam:"style=form,explode=true,name=spoilers"`
}

func (o *GetAnimeReviewsRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *GetAnimeReviewsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetAnimeReviewsRequest) GetPreliminary() *bool {
	if o == nil {
		return nil
	}
	return o.Preliminary
}

func (o *GetAnimeReviewsRequest) GetSpoilers() *bool {
	if o == nil {
		return nil
	}
	return o.Spoilers
}

type GetAnimeReviewsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns anime reviews
	AnimeReviews *components.AnimeReviews
}

func (o *GetAnimeReviewsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAnimeReviewsResponse) GetAnimeReviews() *components.AnimeReviews {
	if o == nil {
		return nil
	}
	return o.AnimeReviews
}
