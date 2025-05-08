package clients_test

import (
	"reflect"
	"testing"

	"github.com/jdplumst/pokeapigo/pkg/clients"
	"github.com/jdplumst/pokeapigo/pkg/types"
)

var (
	cheriBerry = &types.Berry{
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
			{
				Potency: 10,
				Flavor: types.NamedApiResource{
					Name: "spicy",
					Url:  "https://pokeapi.co/api/v2/berry-flavor/1/",
				},
			},
			{
				Potency: 0,
				Flavor: types.NamedApiResource{
					Name: "dry",
					Url:  "https://pokeapi.co/api/v2/berry-flavor/2/",
				},
			},
			{
				Potency: 0,
				Flavor: types.NamedApiResource{
					Name: "sweet",
					Url:  "https://pokeapi.co/api/v2/berry-flavor/3/",
				},
			},
			{
				Potency: 0,
				Flavor: types.NamedApiResource{
					Name: "bitter",
					Url:  "https://pokeapi.co/api/v2/berry-flavor/4/",
				},
			},
			{
				Potency: 0,
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

	verySoftFirmness = &types.BerryFirmness{
		Id:   1,
		Name: "very-soft",
		Berries: []types.NamedApiResource{
			{
				Name: "pecha",
				Url:  "https://pokeapi.co/api/v2/berry/3/",
			},
			{
				Name: "pamtre",
				Url:  "https://pokeapi.co/api/v2/berry/32/",
			},
			{
				Name: "belue",
				Url:  "https://pokeapi.co/api/v2/berry/35/",
			},
			{
				Name: "wacan",
				Url:  "https://pokeapi.co/api/v2/berry/38/",
			},
			{
				Name: "tanga",
				Url:  "https://pokeapi.co/api/v2/berry/46/",
			},
			{
				Name: "charti",
				Url:  "https://pokeapi.co/api/v2/berry/47/",
			},
			{
				Name: "chilan",
				Url:  "https://pokeapi.co/api/v2/berry/52/",
			},
			{
				Name: "rowap",
				Url:  "https://pokeapi.co/api/v2/berry/64/",
			},
		},
		Names: []types.Name{
			{
				Name: "とてもやわらかい",
				Language: types.NamedApiResource{
					Name: "ja-Hrkt",
					Url:  "https://pokeapi.co/api/v2/language/1/",
				},
			},
			{
				Name: "很柔軟",
				Language: types.NamedApiResource{
					Name: "zh-Hant",
					Url:  "https://pokeapi.co/api/v2/language/4/",
				},
			},
			{
				Name: "Très tendre",
				Language: types.NamedApiResource{
					Name: "fr",
					Url:  "https://pokeapi.co/api/v2/language/5/",
				},
			},
			{
				Name: "Muy blanda",
				Language: types.NamedApiResource{
					Name: "es",
					Url:  "https://pokeapi.co/api/v2/language/7/",
				},
			},
			{
				Name: "Very Soft",
				Language: types.NamedApiResource{
					Name: "en",
					Url:  "https://pokeapi.co/api/v2/language/9/",
				},
			},
			{
				Name: "很柔软",
				Language: types.NamedApiResource{
					Name: "zh-Hans",
					Url:  "https://pokeapi.co/api/v2/language/12/",
				},
			},
		},
	}
)

func TestGetBerry(t *testing.T) {
	berryClient := clients.Berries{}

	t.Run("valid id", func(t *testing.T) {
		got, err := berryClient.GetBerryById(1)
		if err != nil {
			t.Errorf("did not return a berry, %v", err)
		}

		if !reflect.DeepEqual(got, cheriBerry) {
			t.Errorf("got %+v\nwant %+v\n", got, cheriBerry)
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

	t.Run("valid name", func(t *testing.T) {
		got, err := berryClient.GetBerryByName("cheri")
		if err != nil {
			t.Errorf("did not return a berry, %v", err)
		}

		if !reflect.DeepEqual(got, cheriBerry) {
			t.Errorf("got %+v\nwant %+v\n", got, cheriBerry)
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

func TestGetBerryFirmness(t *testing.T) {
	berryClient := clients.Berries{}

	t.Run("valid id", func(t *testing.T) {
		got, err := berryClient.GetBerryFirmnessById(1)
		if err != nil {
			t.Errorf("did not return a berry firmness, %v", err)
		}

		if !reflect.DeepEqual(got, verySoftFirmness) {
			t.Errorf("got %+v\nwant %+v\n", got, verySoftFirmness)
		}
	})

	t.Run("invalid id", func(t *testing.T) {
		_, err := berryClient.GetBerryFirmnessById(100)

		switch err {
		case types.ErrFetch, types.ErrBody, types.ErrUnmarshal:
			return
		case nil:
			t.Error("got a berry firmness when expected an error:", err)
		default:
			t.Error("got an unexpected error:", err)

		}
	})

	t.Run("valid name", func(t *testing.T) {
		got, err := berryClient.GetBerryFirmnessByName("very-soft")
		if err != nil {
			t.Errorf("did not return a berry firmness, %v", err)
		}

		if !reflect.DeepEqual(got, verySoftFirmness) {
			t.Errorf("got %+v\nwant %+v\n", got, verySoftFirmness)
		}
	})

	t.Run("invalid name", func(t *testing.T) {
		_, err := berryClient.GetBerryFirmnessByName("invalid")

		switch err {
		case types.ErrFetch, types.ErrBody, types.ErrUnmarshal:
			return
		case nil:
			t.Error("got a berry firmness when expected an error:", err)
		default:
			t.Error("got an unexpected error:", err)

		}
	})
}

func TestGetBerryFlavor(t *testing.T) {
	berryClient := clients.Berries{}

	want := struct {
		Id              int
		Name            string
		NumberOfBerries int
		ContestType     types.NamedApiResource
		NumberOfNames   int
	}{
		Id:              1,
		Name:            "spicy",
		NumberOfBerries: 29,
		ContestType: types.NamedApiResource{
			Name: "cool",
			Url:  "https://pokeapi.co/api/v2/contest-type/1/",
		},
		NumberOfNames: 5,
	}

	t.Run("valid id", func(t *testing.T) {
		got, err := berryClient.GetBerryFlavorById(1)
		if err != nil {
			t.Errorf("did not return a berry flavor, %v", err)
		}

		if got.Id != want.Id {
			t.Errorf("got an id of %v, expected %v", got.Name, want.Id)
		}

		if got.Name != want.Name {
			t.Errorf("got a name of %v, expected %v", got.Name, want.Name)
		}

		if len(got.Berries) != want.NumberOfBerries {
			t.Errorf("got %v berries, expected %v", len(got.Berries), want.NumberOfBerries)
		}

		if got.ContestType != want.ContestType {
			t.Errorf("got contest type %v, want %v", got.ContestType, want.ContestType)
		}

		if len(got.Names) != want.NumberOfNames {
			t.Errorf("got %v names, expected %v", len(got.Names), want.NumberOfNames)
		}

	})

	t.Run("invalid id", func(t *testing.T) {
		_, err := berryClient.GetBerryFlavorById(100)

		switch err {
		case types.ErrFetch, types.ErrBody, types.ErrUnmarshal:
			return
		case nil:
			t.Error("got a berry flavor when expected an error:", err)
		default:
			t.Error("got an unexpected error:", err)

		}
	})

	t.Run("valid name", func(t *testing.T) {
		got, err := berryClient.GetBerryFlavorByName("spicy")
		if err != nil {
			t.Errorf("did not return a berry flavor, %v", err)
		}

		if got.Id != want.Id {
			t.Errorf("got an id of %v, expected %v", got.Name, want.Id)
		}

		if got.Name != want.Name {
			t.Errorf("got a name of %v, expected %v", got.Name, want.Name)
		}

		if len(got.Berries) != want.NumberOfBerries {
			t.Errorf("got %v berries, expected %v", len(got.Berries), want.NumberOfBerries)
		}

		if got.ContestType != want.ContestType {
			t.Errorf("got contest type %v, want %v", got.ContestType, want.ContestType)
		}

		if len(got.Names) != want.NumberOfNames {
			t.Errorf("got %v names, expected %v", len(got.Names), want.NumberOfNames)
		}

	})

	t.Run("invalid name", func(t *testing.T) {
		_, err := berryClient.GetBerryFlavorByName("invalid")

		switch err {
		case types.ErrFetch, types.ErrBody, types.ErrUnmarshal:
			return
		case nil:
			t.Error("got a berry flavor when expected an error:", err)
		default:
			t.Error("got an unexpected error:", err)

		}
	})
}
