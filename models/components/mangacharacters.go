package components

type MangaCharactersData struct {
	Character *CharacterMeta `json:"character,omitempty"`
	// Character's Role
	Role *string `json:"role,omitempty"`
}

func (o *MangaCharactersData) GetCharacter() *CharacterMeta {
	if o == nil {
		return nil
	}
	return o.Character
}

func (o *MangaCharactersData) GetRole() *string {
	if o == nil {
		return nil
	}
	return o.Role
}

// MangaCharacters - Manga Characters Resource
type MangaCharacters struct {
	Data []MangaCharactersData `json:"data,omitempty"`
}

func (o *MangaCharacters) GetData() []MangaCharactersData {
	if o == nil {
		return nil
	}
	return o.Data
}
