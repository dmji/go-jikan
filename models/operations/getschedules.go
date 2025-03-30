package operations

import (
	"encoding/json"
	"fmt"

	"github.com/dmji/jikan/models/components"
)

// GetSchedulesQueryParamFilter - Filter by day
type GetSchedulesQueryParamFilter string

const (
	GetSchedulesQueryParamFilterMonday    GetSchedulesQueryParamFilter = "monday"
	GetSchedulesQueryParamFilterTuesday   GetSchedulesQueryParamFilter = "tuesday"
	GetSchedulesQueryParamFilterWednesday GetSchedulesQueryParamFilter = "wednesday"
	GetSchedulesQueryParamFilterThursday  GetSchedulesQueryParamFilter = "thursday"
	GetSchedulesQueryParamFilterFriday    GetSchedulesQueryParamFilter = "friday"
	GetSchedulesQueryParamFilterSaturday  GetSchedulesQueryParamFilter = "saturday"
	GetSchedulesQueryParamFilterSunday    GetSchedulesQueryParamFilter = "sunday"
	GetSchedulesQueryParamFilterUnknown   GetSchedulesQueryParamFilter = "unknown"
	GetSchedulesQueryParamFilterOther     GetSchedulesQueryParamFilter = "other"
)

func (e GetSchedulesQueryParamFilter) ToPointer() *GetSchedulesQueryParamFilter {
	return &e
}

func (e *GetSchedulesQueryParamFilter) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "monday":
		fallthrough
	case "tuesday":
		fallthrough
	case "wednesday":
		fallthrough
	case "thursday":
		fallthrough
	case "friday":
		fallthrough
	case "saturday":
		fallthrough
	case "sunday":
		fallthrough
	case "unknown":
		fallthrough
	case "other":
		*e = GetSchedulesQueryParamFilter(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSchedulesQueryParamFilter: %v", v)
	}
}

// Kids - When supplied, it will filter entries with the `Kids` Genre Demographic. When supplied as `kids=true`, it will return only Kid entries and when supplied as `kids=false`, it will filter out any Kid entries. Defaults to `false`.
type Kids string

const (
	KidsTrue  Kids = "true"
	KidsFalse Kids = "false"
)

func (e Kids) ToPointer() *Kids {
	return &e
}

func (e *Kids) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "true":
		fallthrough
	case "false":
		*e = Kids(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Kids: %v", v)
	}
}

// Sfw - 'Safe For Work'. When supplied, it will filter entries with the `Hentai` Genre. When supplied as `sfw=true`, it will return only SFW entries and when supplied as `sfw=false`, it will filter out any Hentai entries. Defaults to `false`.
type Sfw string

const (
	SfwTrue  Sfw = "true"
	SfwFalse Sfw = "false"
)

func (e Sfw) ToPointer() *Sfw {
	return &e
}

func (e *Sfw) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "true":
		fallthrough
	case "false":
		*e = Sfw(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Sfw: %v", v)
	}
}

type GetSchedulesRequest struct {
	// Filter by day
	Filter *GetSchedulesQueryParamFilter `queryParam:"style=form,explode=true,name=filter"`
	// When supplied, it will filter entries with the `Kids` Genre Demographic. When supplied as `kids=true`, it will return only Kid entries and when supplied as `kids=false`, it will filter out any Kid entries. Defaults to `false`.
	Kids *Kids `queryParam:"style=form,explode=true,name=kids"`
	// 'Safe For Work'. When supplied, it will filter entries with the `Hentai` Genre. When supplied as `sfw=true`, it will return only SFW entries and when supplied as `sfw=false`, it will filter out any Hentai entries. Defaults to `false`.
	Sfw *Sfw `queryParam:"style=form,explode=true,name=sfw"`
	// This is a flag. When supplied it will include entries which are unapproved. Unapproved entries on MyAnimeList are those that are user submitted and have not yet been approved by MAL to show up on other pages. They will have their own specifc pages and are often removed resulting in a 404 error. You do not need to pass a value to it. e.g usage: `?unapproved`
	Unapproved *bool  `queryParam:"style=form,explode=true,name=unapproved"`
	Page       *int64 `queryParam:"style=form,explode=true,name=page"`
	Limit      *int64 `queryParam:"style=form,explode=true,name=limit"`
}

func (o *GetSchedulesRequest) GetFilter() *GetSchedulesQueryParamFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *GetSchedulesRequest) GetKids() *Kids {
	if o == nil {
		return nil
	}
	return o.Kids
}

func (o *GetSchedulesRequest) GetSfw() *Sfw {
	if o == nil {
		return nil
	}
	return o.Sfw
}

func (o *GetSchedulesRequest) GetUnapproved() *bool {
	if o == nil {
		return nil
	}
	return o.Unapproved
}

func (o *GetSchedulesRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetSchedulesRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

type GetSchedulesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns weekly schedule
	Schedules *components.Schedules
}

func (o *GetSchedulesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetSchedulesResponse) GetSchedules() *components.Schedules {
	if o == nil {
		return nil
	}
	return o.Schedules
}
