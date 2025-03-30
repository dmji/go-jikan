package operations

import (
	"encoding/json"
	"fmt"

	"github.com/dmji/go-jikan/models/components"
)

// Filter topics
type Filter string

const (
	FilterAll     Filter = "all"
	FilterEpisode Filter = "episode"
	FilterOther   Filter = "other"
)

func (e Filter) ToPointer() *Filter {
	return &e
}

func (e *Filter) UnmarshalJSON(data []byte) error {
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
		*e = Filter(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Filter: %v", v)
	}
}

type GetAnimeForumRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
	// Filter topics
	Filter *Filter `queryParam:"style=form,explode=true,name=filter"`
}

func (o *GetAnimeForumRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *GetAnimeForumRequest) GetFilter() *Filter {
	if o == nil {
		return nil
	}
	return o.Filter
}

type GetAnimeForumResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns a list of forum topics related to the entry
	Forum *components.Forum
}

func (o *GetAnimeForumResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAnimeForumResponse) GetForum() *components.Forum {
	if o == nil {
		return nil
	}
	return o.Forum
}
