-- name: ListPokemon :many
select * from pokemon offset $1 limit $2;