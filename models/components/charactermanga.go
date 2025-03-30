package components

type CharacterMangaData struct {
	// Character's Role
	Role  *string    `json:"role,omitempty"`
	Manga *MangaMeta `json:"manga,omitempty"`
}

func (o *CharacterMangaData) GetRole() *string {
	if o == nil {
		return nil
	}
	return o.Role
}

func (o *CharacterMangaData) GetManga() *MangaMeta {
	if o == nil {
		return nil
	}
	return o.Manga
}

// CharacterManga - Character casted in manga
type CharacterManga struct {
	Data []CharacterMangaData `json:"data,omitempty"`
}

func (o *CharacterManga) GetData() []CharacterMangaData {
	if o == nil {
		return nil
	}
	return o.Data
}
