package components

import (
	"encoding/json"
	"fmt"
)

// Club Category
type Category string

const (
	CategoryActorsAndArtists       Category = "actors & artists"
	CategoryAnime                  Category = "anime"
	CategoryCharacters             Category = "characters"
	CategoryCitiesAndNeighborhoods Category = "cities & neighborhoods"
	CategoryCompanies              Category = "companies"
	CategoryConventions            Category = "conventions"
	CategoryGames                  Category = "games"
	CategoryJapan                  Category = "japan"
	CategoryManga                  Category = "manga"
	CategoryMusic                  Category = "music"
	CategoryOthers                 Category = "others"
	CategorySchools                Category = "schools"
)

func (e Category) ToPointer() *Category {
	return &e
}

func (e *Category) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "actors & artists":
		fallthrough
	case "anime":
		fallthrough
	case "characters":
		fallthrough
	case "cities & neighborhoods":
		fallthrough
	case "companies":
		fallthrough
	case "conventions":
		fallthrough
	case "games":
		fallthrough
	case "japan":
		fallthrough
	case "manga":
		fallthrough
	case "music":
		fallthrough
	case "others":
		fallthrough
	case "schools":
		*e = Category(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Category: %v", v)
	}
}

// Access - Club access
type Access string

const (
	AccessPublic  Access = "public"
	AccessPrivate Access = "private"
	AccessSecret  Access = "secret"
)

func (e Access) ToPointer() *Access {
	return &e
}

func (e *Access) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "public":
		fallthrough
	case "private":
		fallthrough
	case "secret":
		*e = Access(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Access: %v", v)
	}
}

// Club Resource
type Club struct {
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// Club name
	Name *string `json:"name,omitempty"`
	// Club URL
	URL    *string       `json:"url,omitempty"`
	Images *CommonImages `json:"images,omitempty"`
	// Number of club members
	Members *int64 `json:"members,omitempty"`
	// Club Category
	Category *Category `json:"category,omitempty"`
	// Date Created ISO8601
	Created *string `json:"created,omitempty"`
	// Club access
	Access *Access `json:"access,omitempty"`
}

func (o *Club) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *Club) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Club) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *Club) GetImages() *CommonImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *Club) GetMembers() *int64 {
	if o == nil {
		return nil
	}
	return o.Members
}

func (o *Club) GetCategory() *Category {
	if o == nil {
		return nil
	}
	return o.Category
}

func (o *Club) GetCreated() *string {
	if o == nil {
		return nil
	}
	return o.Created
}

func (o *Club) GetAccess() *Access {
	if o == nil {
		return nil
	}
	return o.Access
}
