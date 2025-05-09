package types

type ContestType struct {
	Id          int              `json:"id"`
	Name        string           `json:"name"`
	BerryFlavor NamedApiResource `json:"berry_flavor"`
	Names       []ContestName    `json:"names"`
}

type ContestName struct {
	Name     string           `json:"name"`
	Color    string           `json:"color"`
	Language NamedApiResource `json:"language"`
}

type ContestEffect struct {
	Id                int          `json:"id"`
	Appeal            int          `json:"appeal"`
	Jam               int          `json:"jam"`
	EffectEntries     []Effect     `json:"effect_entries"`
	FlavorTextEntries []FlavorText `json:"flavor_text_entries"`
}

type SuperContestEffect struct {
	Id                int                `json:"id"`
	Appeal            int                `json:"appeal"`
	FlavorTextEntries []FlavorText       `json:"flavor_text_entries"`
	Moves             []NamedApiResource `json:"moves"`
}
