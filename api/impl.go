package api

import (
	"context"
	"gen/db"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	db *db.Queries
}

func NewServer() Server {
	ctx := context.Background()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}

	return Server{
		db: db.New(pool),
	}
}

func (s Server) GetPokemon(ctx echo.Context, params GetPokemonParams) error {

	offset := int32(0)
	limit := int32(20)

	if params.Offset != nil {
		offset = *params.Offset
	}
	if params.Limit != nil {
		limit = *params.Limit
	}
	pokemonList, dbError := s.db.ListPokemon(
		ctx.Request().Context(),
		db.ListPokemonParams{
			Offset: offset,
			Limit:  limit,
		},
	)

	if dbError != nil {
		return ctx.JSON(http.StatusInternalServerError, Error{
			Message: "internal server error",
		})
	}

	response := make([]Pokemon, 0, len(pokemonList))

	for _, p := range pokemonList {
		response = append(response, Pokemon{
			Id:           int64Ptr(p.ID),
			NameEnglish:  textPtr(p.NameEnglish),
			NameJapanese: textPtr(p.NameJapanese),
			NameChinese:  textPtr(p.NameChinese),
			NameFrench:   textPtr(p.NameFrench),

			Types: slicePtr(p.Types),

			BaseHp:        int64Ptr(p.BaseHp),
			BaseAttack:    int64Ptr(p.BaseAttack),
			BaseDefense:   int64Ptr(p.BaseDefense),
			BaseSpAttack:  int64Ptr(p.BaseSpAttack),
			BaseSpDefense: int64Ptr(p.BaseSpDefense),
			BaseSpeed:     int64Ptr(p.BaseSpeed),

			Species:     textPtr(p.Species),
			Description: textPtr(p.Description),

			EvolutionNext: slicePtr(p.EvolutionNext),
			EvolutionPrev: slicePtr(p.EvolutionPrev),

			Height: textPtr(p.Height),
			Weight: textPtr(p.Weight),

			EggGroups: slicePtr(p.EggGroups),
			Abilities: slicePtr(p.Abilities),

			Gender: textPtr(p.Gender),

			SpriteUrl:    textPtr(p.SpriteUrl),
			ThumbnailUrl: textPtr(p.ThumbnailUrl),
			HiresUrl:     textPtr(p.HiresUrl),
		})
	}

	return ctx.JSON(http.StatusOK, response)
}

func textPtr(v pgtype.Text) *string {
	if !v.Valid {
		return nil
	}
	return &v.String
}

func int64Ptr(v pgtype.Int8) *int64 {
	if !v.Valid {
		return nil
	}
	return &v.Int64
}

func slicePtr[T any](s []T) *[]T {
	if len(s) == 0 {
		return nil
	}
	return &s
}
