package components

type PersonFullAnime struct {
	// Person's position
	Position *string    `json:"position,omitempty"`
	Anime    *AnimeMeta `json:"anime,omitempty"`
}

func (o *PersonFullAnime) GetPosition() *string {
	if o == nil {
		return nil
	}
	return o.Position
}

func (o *PersonFullAnime) GetAnime() *AnimeMeta {
	if o == nil {
		return nil
	}
	return o.Anime
}

type PersonFullManga struct {
	// Person's position
	Position *string    `json:"position,omitempty"`
	Manga    *MangaMeta `json:"manga,omitempty"`
}

func (o *PersonFullManga) GetPosition() *string {
	if o == nil {
		return nil
	}
	return o.Position
}

func (o *PersonFullManga) GetManga() *MangaMeta {
	if o == nil {
		return nil
	}
	return o.Manga
}

type PersonFullVoices struct {
	// Person's Character's role in the anime
	Role      *string        `json:"role,omitempty"`
	Anime     *AnimeMeta     `json:"anime,omitempty"`
	Character *CharacterMeta `json:"character,omitempty"`
}

func (o *PersonFullVoices) GetRole() *string {
	if o == nil {
		return nil
	}
	return o.Role
}

func (o *PersonFullVoices) GetAnime() *AnimeMeta {
	if o == nil {
		return nil
	}
	return o.Anime
}

func (o *PersonFullVoices) GetCharacter() *CharacterMeta {
	if o == nil {
		return nil
	}
	return o.Character
}

// PersonFull - Person Resource
type PersonFull struct {
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList URL
	URL *string `json:"url,omitempty"`
	// Person's website URL
	WebsiteURL *string       `json:"website_url,omitempty"`
	Images     *PeopleImages `json:"images,omitempty"`
	// Name
	Name *string `json:"name,omitempty"`
	// Given Name
	GivenName *string `json:"given_name,omitempty"`
	// Family Name
	FamilyName *string `json:"family_name,omitempty"`
	// Other Names
	AlternateNames []string `json:"alternate_names,omitempty"`
	// Birthday Date ISO8601
	Birthday *string `json:"birthday,omitempty"`
	// Number of users who have favorited this entry
	Favorites *int64 `json:"favorites,omitempty"`
	// Biography
	About  *string            `json:"about,omitempty"`
	Anime  []PersonFullAnime  `json:"anime,omitempty"`
	Manga  []PersonFullManga  `json:"manga,omitempty"`
	Voices []PersonFullVoices `json:"voices,omitempty"`
}

func (o *PersonFull) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *PersonFull) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *PersonFull) GetWebsiteURL() *string {
	if o == nil {
		return nil
	}
	return o.WebsiteURL
}

func (o *PersonFull) GetImages() *PeopleImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *PersonFull) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PersonFull) GetGivenName() *string {
	if o == nil {
		return nil
	}
	return o.GivenName
}

func (o *PersonFull) GetFamilyName() *string {
	if o == nil {
		return nil
	}
	return o.FamilyName
}

func (o *PersonFull) GetAlternateNames() []string {
	if o == nil {
		return nil
	}
	return o.AlternateNames
}

func (o *PersonFull) GetBirthday() *string {
	if o == nil {
		return nil
	}
	return o.Birthday
}

func (o *PersonFull) GetFavorites() *int64 {
	if o == nil {
		return nil
	}
	return o.Favorites
}

func (o *PersonFull) GetAbout() *string {
	if o == nil {
		return nil
	}
	return o.About
}

func (o *PersonFull) GetAnime() []PersonFullAnime {
	if o == nil {
		return nil
	}
	return o.Anime
}

func (o *PersonFull) GetManga() []PersonFullManga {
	if o == nil {
		return nil
	}
	return o.Manga
}

func (o *PersonFull) GetVoices() []PersonFullVoices {
	if o == nil {
		return nil
	}
	return o.Voices
}
