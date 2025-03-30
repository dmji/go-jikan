package components

type CharacterAnimeData struct {
	// Character's Role
	Role  *string    `json:"role,omitempty"`
	Anime *AnimeMeta `json:"anime,omitempty"`
}

func (o *CharacterAnimeData) GetRole() *string {
	if o == nil {
		return nil
	}
	return o.Role
}

func (o *CharacterAnimeData) GetAnime() *AnimeMeta {
	if o == nil {
		return nil
	}
	return o.Anime
}

// CharacterAnime - Character casted in anime
type CharacterAnime struct {
	Data []CharacterAnimeData `json:"data,omitempty"`
}

func (o *CharacterAnime) GetData() []CharacterAnimeData {
	if o == nil {
		return nil
	}
	return o.Data
}
