package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetProducerFullByIDRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetProducerFullByIDRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

// GetProducerFullByIDResponseBody - Returns producer resource
type GetProducerFullByIDResponseBody struct {
	// Producers Resource
	Data *components.ProducerFull `json:"data,omitempty"`
}

func (o *GetProducerFullByIDResponseBody) GetData() *components.ProducerFull {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetProducerFullByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns producer resource
	Object *GetProducerFullByIDResponseBody
}

func (o *GetProducerFullByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetProducerFullByIDResponse) GetObject() *GetProducerFullByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
