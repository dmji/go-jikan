package components

type UserFavoritesAnime struct {
	Type      *string `json:"type,omitempty"`
	StartYear *int64  `json:"start_year,omitempty"`
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList URL
	URL    *string      `json:"url,omitempty"`
	Images *AnimeImages `json:"images,omitempty"`
	// Entry title
	Title *string `json:"title,omitempty"`
}

func (o *UserFavoritesAnime) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *UserFavoritesAnime) GetStartYear() *int64 {
	if o == nil {
		return nil
	}
	return o.StartYear
}

func (o *UserFavoritesAnime) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *UserFavoritesAnime) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *UserFavoritesAnime) GetImages() *AnimeImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *UserFavoritesAnime) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

type UserFavoritesManga struct {
	Type      *string `json:"type,omitempty"`
	StartYear *int64  `json:"start_year,omitempty"`
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList URL
	URL    *string      `json:"url,omitempty"`
	Images *MangaImages `json:"images,omitempty"`
	// Entry title
	Title *string `json:"title,omitempty"`
}

func (o *UserFavoritesManga) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *UserFavoritesManga) GetStartYear() *int64 {
	if o == nil {
		return nil
	}
	return o.StartYear
}

func (o *UserFavoritesManga) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *UserFavoritesManga) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *UserFavoritesManga) GetImages() *MangaImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *UserFavoritesManga) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

// Characters - Parsed URL Data
type Characters struct {
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList URL
	URL    *string          `json:"url,omitempty"`
	Images *CharacterImages `json:"images,omitempty"`
	// Entry name
	Name *string `json:"name,omitempty"`
	// Type of resource
	Type *string `json:"type,omitempty"`
	// Resource Name/Title
	Title *string `json:"title,omitempty"`
}

func (o *Characters) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *Characters) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *Characters) GetImages() *CharacterImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *Characters) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Characters) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Characters) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

type UserFavorites struct {
	// Favorite Anime
	Anime []UserFavoritesAnime `json:"anime,omitempty"`
	// Favorite Manga
	Manga []UserFavoritesManga `json:"manga,omitempty"`
	// Favorite Characters
	Characters []Characters `json:"characters,omitempty"`
	// Favorite People
	People []CharacterMeta `json:"people,omitempty"`
}

func (o *UserFavorites) GetAnime() []UserFavoritesAnime {
	if o == nil {
		return nil
	}
	return o.Anime
}

func (o *UserFavorites) GetManga() []UserFavoritesManga {
	if o == nil {
		return nil
	}
	return o.Manga
}

func (o *UserFavorites) GetCharacters() []Characters {
	if o == nil {
		return nil
	}
	return o.Characters
}

func (o *UserFavorites) GetPeople() []CharacterMeta {
	if o == nil {
		return nil
	}
	return o.People
}
