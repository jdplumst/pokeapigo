package clients_test

import (
	"reflect"
	"testing"

	"github.com/jdplumst/pokeapigo/pkg/clients"
	"github.com/jdplumst/pokeapigo/pkg/types"
)

var (
	want = &types.Berry{
		Id:               1,
		Name:             "cheri",
		GrowthTime:       3,
		MaxHarvest:       5,
		NaturalGiftPower: 60,
		Size:             20,
		Smoothness:       25,
		SoilDryness:      15,
		Firmness: types.NamedApiResource{
			Name: "soft",
			Url:  "https://pokeapi.co/api/v2/berry-firmness/2/",
		},
		Flavors: []types.BerryFlavorMap{
			{Potency: 10,
				Flavor: types.NamedApiResource{
					Name: "spicy",
					Url:  "https://pokeapi.co/api/v2/berry-flavor/1/",
				},
			},
			{Potency: 0,
				Flavor: types.NamedApiResource{
					Name: "dry",
					Url:  "https://pokeapi.co/api/v2/berry-flavor/2/",
				},
			},
			{Potency: 0,
				Flavor: types.NamedApiResource{
					Name: "sweet",
					Url:  "https://pokeapi.co/api/v2/berry-flavor/3/",
				},
			},
			{Potency: 0,
				Flavor: types.NamedApiResource{
					Name: "bitter",
					Url:  "https://pokeapi.co/api/v2/berry-flavor/4/",
				},
			},
			{Potency: 0,
				Flavor: types.NamedApiResource{
					Name: "sour",
					Url:  "https://pokeapi.co/api/v2/berry-flavor/5/",
				},
			},
		},
		Item: types.NamedApiResource{
			Name: "cheri-berry",
			Url:  "https://pokeapi.co/api/v2/item/126/",
		},
		NaturalGiftType: types.NamedApiResource{
			Name: "fire",
			Url:  "https://pokeapi.co/api/v2/type/10/",
		},
	}
)

func TestGetBerryById(t *testing.T) {
	berryClient := clients.Berries{}

	t.Run("valid id", func(t *testing.T) {
		got, err := berryClient.GetBerryById(1)
		if err != nil {
			t.Errorf("did not return a berry, %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v\nwant %+v\n", got, want)
		}
	})

	t.Run("invalid id", func(t *testing.T) {
		_, err := berryClient.GetBerryById(100)

		switch err {
		case types.ErrFetch, types.ErrBody, types.ErrUnmarshal:
			return
		case nil:
			t.Error("got a berry when expected an error:", err)
		default:
			t.Error("got an unexpected error:", err)

		}
	})
}

func TestGerBerryByName(t *testing.T) {
	berryClient := clients.Berries{}

	t.Run("valid name", func(t *testing.T) {
		got, err := berryClient.GetBerryByName("cheri")
		if err != nil {
			t.Errorf("did not return a berry, %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v\nwant %+v\n", got, want)
		}
	})

	t.Run("invalid name", func(t *testing.T) {
		_, err := berryClient.GetBerryByName("invalid")

		switch err {
		case types.ErrFetch, types.ErrBody, types.ErrUnmarshal:
			return
		case nil:
			t.Error("got a berry when expected an error:", err)
		default:
			t.Error("got an unexpected error:", err)

		}
	})
}
