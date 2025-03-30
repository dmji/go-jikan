package components

type PersonVoiceActingRolesData struct {
	// Person's Character's role in the anime
	Role      *string        `json:"role,omitempty"`
	Anime     *AnimeMeta     `json:"anime,omitempty"`
	Character *CharacterMeta `json:"character,omitempty"`
}

func (o *PersonVoiceActingRolesData) GetRole() *string {
	if o == nil {
		return nil
	}
	return o.Role
}

func (o *PersonVoiceActingRolesData) GetAnime() *AnimeMeta {
	if o == nil {
		return nil
	}
	return o.Anime
}

func (o *PersonVoiceActingRolesData) GetCharacter() *CharacterMeta {
	if o == nil {
		return nil
	}
	return o.Character
}

// PersonVoiceActingRoles - Person's voice acting roles
type PersonVoiceActingRoles struct {
	Data []PersonVoiceActingRolesData `json:"data,omitempty"`
}

func (o *PersonVoiceActingRoles) GetData() []PersonVoiceActingRolesData {
	if o == nil {
		return nil
	}
	return o.Data
}
