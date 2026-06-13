package main

type PokemonUrl struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonGet struct {
	Count    int          `json:"count"`
	Next     string       `json:"next"`
	Previous string       `json:"previous"`
	Results  []PokemonUrl `json:"results"`
}

type PokemonAbility struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	IsHidden bool   `json:"is_hidden"`
	Slot     int    `json:"slot"`
}

type Ability struct {
	Ability PokemonAbility `json:"ability"`
}

type Pokemon struct {
	Id        int       `json:"id"`
	Order     int       `json:"order"`
	Name      string    `json:"name"`
	Weight    int       `json:"weight"`
	Abilities []Ability `json:"abilities"`
}
