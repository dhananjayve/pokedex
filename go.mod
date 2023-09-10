module github.com/dhananjayve/pokedex

go 1.20

//replace github.com/dhananjayve/pokedex/internal/pockeapi v0.0.0 => ./internal/pockeapi

replace github.com/dhananjayve/pokedex/internal/pokeapi v0.0.0 => ./internal/pokeapi

replace github.com/dhananjayve/pokedex/internal/pokecache v0.0.0 => ./internal/pokecache

require (
	github.com/dhananjayve/pokedex/internal/pokeapi v0.0.0
	github.com/dhananjayve/pokedex/internal/pokecache v0.0.0
)
