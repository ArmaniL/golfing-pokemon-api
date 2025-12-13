package api

import (
	"math/rand/v2"
	"net/http"

	"github.com/labstack/echo/v4"
)

// optional code omitted

type Server struct {
	ListOfPokemon []PokemonInput
	Count         int
}

func NewServer() Server {
	listofPokemon := loadFile()
	return Server{
		ListOfPokemon: listofPokemon,
		Count:         len(listofPokemon),
	}
}

// (GET /ping)
func (s Server) GetPokemon(ctx echo.Context) error {

	randomNumber := rand.IntN(s.Count)
	randomPokemon := s.ListOfPokemon[randomNumber]

	resp := Pokemon{
		Name: randomPokemon.English,
	}

	return ctx.JSON(http.StatusOK, resp)
}
