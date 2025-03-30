package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetClubsByIDRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetClubsByIDRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

// GetClubsByIDResponseBody - Returns Club Resource
type GetClubsByIDResponseBody struct {
	// Club Resource
	Data *components.Club `json:"data,omitempty"`
}

func (o *GetClubsByIDResponseBody) GetData() *components.Club {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetClubsByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns Club Resource
	Object *GetClubsByIDResponseBody
}

func (o *GetClubsByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetClubsByIDResponse) GetObject() *GetClubsByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
