package middlewares

import (
	"context"
	"log"
	"net/http"
	"strconv"

	c "github.com/go-chi/chi/v5"
	"github.com/joaovds/rinha-crebito/internal/di"
	"github.com/joaovds/rinha-crebito/internal/domain"
	cc "github.com/joaovds/rinha-crebito/internal/infra/chi/contracts"
)

func AccountFromContext(ctx context.Context) (*domain.Account, bool) {
	account, ok := ctx.Value("account").(*domain.Account)
	return account, ok
}

func CheckIDParam(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIdParam := c.URLParam(r, "id")
		clientId, err := strconv.Atoi(clientIdParam)
		if err != nil {
			http.Error(w, cc.NewErrorResponse(http.StatusNotFound, "account not found").String(), http.StatusNotFound)
			log.Println(err.Error())
			return
		}

		accountUC := di.NewAccountUsecases()
		account, err := accountUC.GetAccountByID(clientId)
		if err != nil {
			if err == domain.ErrAccountNotFound {
				http.Error(w, cc.NewErrorResponse(http.StatusNotFound, err.Error()).String(), http.StatusNotFound)
				log.Println(err.Error())
				return
			}

			http.Error(w, cc.NewErrorResponse(http.StatusInternalServerError, "internal server error").String(), http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), "account", account)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
