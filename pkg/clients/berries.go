package clients

import (
	"strconv"

	"github.com/jdplumst/pokeapigo/internal/request"
	"github.com/jdplumst/pokeapigo/pkg/types"
)

type IBerries interface {
	GetBerryById(id int) (*types.Berry, error)
	GetBerryByName(name string) (*types.Berry, error)
	GetBerryFirmnessById(id int) (*types.BerryFirmness, error)
	GetBerryFirmnessByName(name string) (*types.BerryFirmness, error)
	GetBerryFlavorById(id int) (*types.BerryFlavor, error)
	GetBerryFlavorByName(name string) (*types.BerryFlavor, error)
}

type Berries struct{}

func (b Berries) GetBerryById(id int) (*types.Berry, error) {
	var berry types.Berry

	parsedId := strconv.Itoa(id)

	url := "https://pokeapi.co/api/v2/berry/" + parsedId

	return request.GetData(url, berry)
}

func (b Berries) GetBerryByName(name string) (*types.Berry, error) {
	var berry types.Berry

	url := "https://pokeapi.co/api/v2/berry/" + name

	return request.GetData(url, berry)
}

func (b Berries) GetBerryFirmnessById(id int) (*types.BerryFirmness, error) {
	var berryFirmness types.BerryFirmness

	parsedId := strconv.Itoa(id)

	url := "https://pokeapi.co/api/v2/berry-firmness/" + parsedId

	return request.GetData(url, berryFirmness)
}

func (b Berries) GetBerryFirmnessByName(name string) (*types.BerryFirmness, error) {
	var berryFirmness types.BerryFirmness

	url := "https://pokeapi.co/api/v2/berry-firmness/" + name

	return request.GetData(url, berryFirmness)
}

func (b Berries) GetBerryFlavorById(id int) (*types.BerryFlavor, error) {
	var berryFlavor types.BerryFlavor

	parsedId := strconv.Itoa(id)

	url := "https://pokeapi.co/api/v2/berry-flavor/" + parsedId

	return request.GetData(url, berryFlavor)
}

func (b Berries) GetBerryFlavorByName(name string) (*types.BerryFlavor, error) {
	var berryFlavor types.BerryFlavor

	url := "https://pokeapi.co/api/v2/berry-flavor/" + name

	return request.GetData(url, berryFlavor)
}
