package types

type Berry struct {
	Id               int              `json:"id"`
	Name             string           `json:"name"`
	GrowthTime       int              `json:"growth_time"`
	MaxHarvest       int              `json:"max_harvest"`
	NaturalGiftPower int              `json:"natural_gift_power"`
	Size             int              `json:"size"`
	Smoothness       int              `json:"smoothness"`
	SoilDryness      int              `json:"soil_dryness"`
	Firmness         NamedApiResource `json:"firmness"`
	Flavors          []BerryFlavorMap `json:"flavor"`
	Item             NamedApiResource `json:"item"`
	NaturalGiftType  NamedApiResource `json:"natural_gift_type"`
}

type BerryFlavorMap struct {
	Potency int              `json:"potency"`
	Flavor  NamedApiResource `json:"flavor"`
}

type BerryFirmness struct {
	Id      int                `json:"id"`
	Name    string             `json:"name"`
	Berries []NamedApiResource `json:"berries"`
	Names   []Name             `json:"names"`
}

type BerryFlavor struct {
	Id          int              `json:"id"`
	Name        string           `json:"name"`
	Berries     []FlavorBerryMap `json:"berries"`
	ContestType NamedApiResource `json:"contest_type"`
	Names       []Name           `json:"names"`
}

type FlavorBerryMap struct {
	Potency int              `json:"potency"`
	Berry   NamedApiResource `json:"berry"`
}
