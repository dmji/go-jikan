package components

// Broadcast Details
type Broadcast struct {
	// Day of the week
	Day *string `json:"day,omitempty"`
	// Time in 24 hour format
	Time *string `json:"time,omitempty"`
	// Timezone (Tz Database format https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)
	Timezone *string `json:"timezone,omitempty"`
	// Raw parsed broadcast string
	String *string `json:"string,omitempty"`
}

func (o *Broadcast) GetDay() *string {
	if o == nil {
		return nil
	}
	return o.Day
}

func (o *Broadcast) GetTime() *string {
	if o == nil {
		return nil
	}
	return o.Time
}

func (o *Broadcast) GetTimezone() *string {
	if o == nil {
		return nil
	}
	return o.Timezone
}

func (o *Broadcast) GetString() *string {
	if o == nil {
		return nil
	}
	return o.String
}
