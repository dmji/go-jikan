package components

type CharacterVoiceActorsData struct {
	// Character's Role
	Language *string     `json:"language,omitempty"`
	Person   *PersonMeta `json:"person,omitempty"`
}

func (o *CharacterVoiceActorsData) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *CharacterVoiceActorsData) GetPerson() *PersonMeta {
	if o == nil {
		return nil
	}
	return o.Person
}

// CharacterVoiceActors - Character voice actors
type CharacterVoiceActors struct {
	Data []CharacterVoiceActorsData `json:"data,omitempty"`
}

func (o *CharacterVoiceActors) GetData() []CharacterVoiceActorsData {
	if o == nil {
		return nil
	}
	return o.Data
}
