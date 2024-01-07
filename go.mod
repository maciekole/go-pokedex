module github.com/maciekole/pokedex

go 1.21.1

require (
	github.com/maciekole/pokedex/pokeapi v0.0.0
	github.com/maciekole/pokedex/pokecache v0.0.0
)
replace github.com/maciekole/pokedex/pokeapi v0.0.0 => ./pokeapi
replace github.com/maciekole/pokedex/pokecache v0.0.0 => ./pokecache