package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetAnimeRelationsRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetAnimeRelationsRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

// GetAnimeRelationsResponseBody - Returns anime relations
type GetAnimeRelationsResponseBody struct {
	Data []components.Relation `json:"data,omitempty"`
}

func (o *GetAnimeRelationsResponseBody) GetData() []components.Relation {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetAnimeRelationsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns anime relations
	Object *GetAnimeRelationsResponseBody
}

func (o *GetAnimeRelationsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAnimeRelationsResponse) GetObject() *GetAnimeRelationsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
