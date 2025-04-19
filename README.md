# pokedex-cli

A basic Pokedex following Boot.dev's "Build a Pokedex in Go" Course which covers the following commands:
  - pokedex: Show caught Pokemons
  - map: Next available page of locations
  - mapb: Previous available page of locations
  - explore: Explore a location
  - catch: Attempt to catch a Pokemon
  - inspect: Inspect a caught Pokemon
  - help: Display available commands and usage
  - exit: Close the program

Data provided by PokeApi(https://pokeapi.co/).

Possible improvements:
  * Update the CLI to support the "up" arrow to cycle through previous commands
  * Simulate battles between pokemon
  * Add more unit tests
  * Refactor your code to organize it better and make it more testable
  * Keep pokemon in a "party" and allow them to level up
  * Allow for pokemon that are caught to evolve after a set amount of time
  * Persist a user's Pokedex to disk so they can save progress between sessions
  * Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
  * Random encounters with wild pokemon
  * Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon
