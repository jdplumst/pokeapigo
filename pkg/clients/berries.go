package clients

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/jdplumst/pokeapigo/pkg/types"
)

type IBerries interface {
	GetBerryById(id int) (types.Berry, error)
	GetBerryByName(name string) (types.Berry, error)
}

type Berries struct{}

func (b Berries) GetBerryById(id int) (types.Berry, error) {
	var berry types.Berry

	parsedId := strconv.Itoa(id)

	url := "https://pokeapi.co/api/v2/berry/" + parsedId

	response, err := http.Get(url)
	if err != nil {
		return berry, types.ErrFetch
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return berry, types.ErrBody
	}

	if err = json.Unmarshal(body, &berry); err != nil {
		return berry, types.ErrUnmarshal
	}

	return berry, nil
}
