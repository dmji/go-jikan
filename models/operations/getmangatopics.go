package operations

import (
	"encoding/json"
	"fmt"

	"github.com/dmji/go-jikan/models/components"
)

// QueryParamFilter - Filter topics
type QueryParamFilter string

const (
	QueryParamFilterAll     QueryParamFilter = "all"
	QueryParamFilterEpisode QueryParamFilter = "episode"
	QueryParamFilterOther   QueryParamFilter = "other"
)

func (e QueryParamFilter) ToPointer() *QueryParamFilter {
	return &e
}

func (e *QueryParamFilter) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "all":
		fallthrough
	case "episode":
		fallthrough
	case "other":
		*e = QueryParamFilter(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamFilter: %v", v)
	}
}

type GetMangaTopicsRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
	// Filter topics
	Filter *QueryParamFilter `queryParam:"style=form,explode=true,name=filter"`
}

func (o *GetMangaTopicsRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *GetMangaTopicsRequest) GetFilter() *QueryParamFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

type GetMangaTopicsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns a list of manga forum topics
	Forum *components.Forum
}

func (o *GetMangaTopicsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetMangaTopicsResponse) GetForum() *components.Forum {
	if o == nil {
		return nil
	}
	return o.Forum
}
