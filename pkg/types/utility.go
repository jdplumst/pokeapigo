package types

type Language struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Official bool   `json:"official"`
	Iso639   string `json:"iso639"`
	Iso3166  string `json:"iso3166"`
	Names    []Name `json:"names"`
}

type ApiResource struct {
	Url string `json:"url"`
}

type Description struct {
	Description string           `json:"description"`
	Language    NamedApiResource `json:"language"`
}

type Effect struct {
	Effect   string           `json:"effect"`
	Language NamedApiResource `json:"language"`
}

type Encounter struct {
	MinLevel        int                `json:"min_level"`
	MaxLevel        int                `json:"max_level"`
	ConditionValues []NamedApiResource `json:"condition_values"`
	Chance          int                `json:"chance"`
	Method          NamedApiResource   `json:"method"`
}

type FlavorText struct {
	FlavorText string           `json:"flavor_text"`
	Language   NamedApiResource `json:"language"`
	Version    NamedApiResource `json:"version"`
}

type GenerationGameIndex struct {
	GameIndex  int              `json:"game_index"`
	Generation NamedApiResource `json:"generation"`
}

type MachineVersionDetail struct {
	Machine      ApiResource      `json:"machine"`
	VersionGroup NamedApiResource `json:"version_group"`
}

type Name struct {
	Name     string           `json:"name"`
	Language NamedApiResource `json:"language"`
}

type NamedApiResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type VerboseEffect struct {
	Effect      string           `json:"effect"`
	ShortEffect string           `json:"short_effect"`
	Language    NamedApiResource `json:"language"`
}

type VersionEncounterDetail struct {
	Version          NamedApiResource `json:"version"`
	MaxChance        int              `json:"max_chance"`
	EncounterDetails []Encounter      `json:"encounter_details"`
}

type VersionGameIndex struct {
	GameIndex int              `json:"game_index"`
	Version   NamedApiResource `json:"version"`
}

type VersionGroupFlavorText struct {
	Text         string           `json:"text"`
	Language     NamedApiResource `json:"language"`
	VersionGroup NamedApiResource `json:"version_group"`
}
