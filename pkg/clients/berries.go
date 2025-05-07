package clients

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/jdplumst/pokeapigo/pkg/types"
)

type IBerries interface {
	GetBerryById(id int) (*types.Berry, error)
	GetBerryByName(name string) (*types.Berry, error)
	GetBerryFirmnessById(id int) (*types.BerryFirmness, error)
}

type Berries struct{}

func (b Berries) GetBerryById(id int) (*types.Berry, error) {
	var berry types.Berry

	parsedId := strconv.Itoa(id)

	url := "https://pokeapi.co/api/v2/berry/" + parsedId

	response, err := http.Get(url)
	if err != nil {
		return nil, types.ErrFetch
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, types.ErrBody
	}

	if err = json.Unmarshal(body, &berry); err != nil {
		return nil, types.ErrUnmarshal
	}

	return &berry, nil
}

func (b Berries) GetBerryByName(name string) (*types.Berry, error) {
	var berry types.Berry

	url := "https://pokeapi.co/api/v2/berry/" + name

	response, err := http.Get(url)
	if err != nil {
		return nil, types.ErrFetch
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, types.ErrBody
	}

	if err = json.Unmarshal(body, &berry); err != nil {
		return nil, types.ErrUnmarshal
	}

	return &berry, nil
}

func (b Berries) GetBerryFirmnessById(id int) (*types.BerryFirmness, error) {
	var berryFirmness types.BerryFirmness

	parsedId := strconv.Itoa(id)

	url := "https://pokeapi.co/api/v2/berry-firmness/" + parsedId

	response, err := http.Get(url)
	if err != nil {
		return nil, types.ErrFetch
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, types.ErrBody
	}

	if err = json.Unmarshal(body, &berryFirmness); err != nil {
		return nil, types.ErrUnmarshal
	}

	return &berryFirmness, nil
}
