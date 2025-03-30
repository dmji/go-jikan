package components

type CharacterFullAnime struct {
	// Character's Role
	Role  *string    `json:"role,omitempty"`
	Anime *AnimeMeta `json:"anime,omitempty"`
}

func (o *CharacterFullAnime) GetRole() *string {
	if o == nil {
		return nil
	}
	return o.Role
}

func (o *CharacterFullAnime) GetAnime() *AnimeMeta {
	if o == nil {
		return nil
	}
	return o.Anime
}

type CharacterFullManga struct {
	// Character's Role
	Role  *string    `json:"role,omitempty"`
	Manga *MangaMeta `json:"manga,omitempty"`
}

func (o *CharacterFullManga) GetRole() *string {
	if o == nil {
		return nil
	}
	return o.Role
}

func (o *CharacterFullManga) GetManga() *MangaMeta {
	if o == nil {
		return nil
	}
	return o.Manga
}

type Voices struct {
	// Character's Role
	Language *string     `json:"language,omitempty"`
	Person   *PersonMeta `json:"person,omitempty"`
}

func (o *Voices) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *Voices) GetPerson() *PersonMeta {
	if o == nil {
		return nil
	}
	return o.Person
}

// CharacterFull - Character Resource
type CharacterFull struct {
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
	About  *string              `json:"about,omitempty"`
	Anime  []CharacterFullAnime `json:"anime,omitempty"`
	Manga  []CharacterFullManga `json:"manga,omitempty"`
	Voices []Voices             `json:"voices,omitempty"`
}

func (o *CharacterFull) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *CharacterFull) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *CharacterFull) GetImages() *CharacterImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *CharacterFull) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CharacterFull) GetNameKanji() *string {
	if o == nil {
		return nil
	}
	return o.NameKanji
}

func (o *CharacterFull) GetNicknames() []string {
	if o == nil {
		return nil
	}
	return o.Nicknames
}

func (o *CharacterFull) GetFavorites() *int64 {
	if o == nil {
		return nil
	}
	return o.Favorites
}

func (o *CharacterFull) GetAbout() *string {
	if o == nil {
		return nil
	}
	return o.About
}

func (o *CharacterFull) GetAnime() []CharacterFullAnime {
	if o == nil {
		return nil
	}
	return o.Anime
}

func (o *CharacterFull) GetManga() []CharacterFullManga {
	if o == nil {
		return nil
	}
	return o.Manga
}

func (o *CharacterFull) GetVoices() []Voices {
	if o == nil {
		return nil
	}
	return o.Voices
}
