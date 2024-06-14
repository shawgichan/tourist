package places

import (
	"context"
	"errors"
	"net/http"

	db "github.com/shawgichan/tourist/db/sqlc"

	"github.com/99designs/gqlgen/graphql"
	"github.com/jackc/pgx/v5"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

/**

   go get github.com/99designs/gqlgen/internal/imports
   go get github.com/99designs/gqlgen
   go run github.com/99designs/gqlgen generate

**/

type Resolver struct {
	Store db.Store
}

type queryResolver struct {
	*Resolver
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func buildNotFound(c context.Context, err error) error {
	return &gqlerror.Error{
		Message: err.Error(),
		Path:    graphql.GetPath(c),
		Extensions: map[string]interface{}{
			"code": http.StatusNotFound,
		},
	}
}

func buildError(c context.Context, err error) error {

	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return buildNotFound(c, err)

	default:
		return &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(c),
			Extensions: map[string]interface{}{
				"code": http.StatusInternalServerError,
			},
		}
	}
}
