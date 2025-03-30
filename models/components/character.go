package components

// Character Resource
type Character struct {
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList URL
	URL    *string          `json:"url,omitempty"`
	Images *CharacterImages `json:"images,omitempty"`
	// Name
	Name *string `json:"name,omitempty"`
	// Name
	NameKanji *string `json:"name_kanji,omitempty"`
	// Other Names
	Nicknames []string `json:"nicknames,omitempty"`
	// Number of users who have favorited this entry
	Favorites *int64 `json:"favorites,omitempty"`
	// Biography
	About *string `json:"about,omitempty"`
}

func (o *Character) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *Character) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *Character) GetImages() *CharacterImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *Character) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Character) GetNameKanji() *string {
	if o == nil {
		return nil
	}
	return o.NameKanji
}

func (o *Character) GetNicknames() []string {
	if o == nil {
		return nil
	}
	return o.Nicknames
}

func (o *Character) GetFavorites() *int64 {
	if o == nil {
		return nil
	}
	return o.Favorites
}

func (o *Character) GetAbout() *string {
	if o == nil {
		return nil
	}
	return o.About
}
