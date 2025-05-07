package request

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jdplumst/pokeapigo/pkg/types"
)

func GetData[T any](url string, resource T) (*T, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, types.ErrFetch
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, types.ErrBody
	}

	if err = json.Unmarshal(body, &resource); err != nil {
		return nil, types.ErrUnmarshal
	}

	return &resource, nil
}
