package components

// AnimeCharactersCharacter - Character details
type AnimeCharactersCharacter struct {
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList URL
	URL    *string          `json:"url,omitempty"`
	Images *CharacterImages `json:"images,omitempty"`
	// Character Name
	Name *string `json:"name,omitempty"`
}

func (o *AnimeCharactersCharacter) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *AnimeCharactersCharacter) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *AnimeCharactersCharacter) GetImages() *CharacterImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *AnimeCharactersCharacter) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type AnimeCharactersPerson struct {
	MalID  *int64        `json:"mal_id,omitempty"`
	URL    *string       `json:"url,omitempty"`
	Images *PeopleImages `json:"images,omitempty"`
	Name   *string       `json:"name,omitempty"`
}

func (o *AnimeCharactersPerson) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *AnimeCharactersPerson) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *AnimeCharactersPerson) GetImages() *PeopleImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *AnimeCharactersPerson) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type VoiceActors struct {
	Person   *AnimeCharactersPerson `json:"person,omitempty"`
	Language *string                `json:"language,omitempty"`
}

func (o *VoiceActors) GetPerson() *AnimeCharactersPerson {
	if o == nil {
		return nil
	}
	return o.Person
}

func (o *VoiceActors) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

type Data struct {
	// Character details
	Character *AnimeCharactersCharacter `json:"character,omitempty"`
	// Character's Role
	Role        *string       `json:"role,omitempty"`
	VoiceActors []VoiceActors `json:"voice_actors,omitempty"`
}

func (o *Data) GetCharacter() *AnimeCharactersCharacter {
	if o == nil {
		return nil
	}
	return o.Character
}

func (o *Data) GetRole() *string {
	if o == nil {
		return nil
	}
	return o.Role
}

func (o *Data) GetVoiceActors() []VoiceActors {
	if o == nil {
		return nil
	}
	return o.VoiceActors
}

// AnimeCharacters - Anime Characters Resource
type AnimeCharacters struct {
	Data []Data `json:"data,omitempty"`
}

func (o *AnimeCharacters) GetData() []Data {
	if o == nil {
		return nil
	}
	return o.Data
}
