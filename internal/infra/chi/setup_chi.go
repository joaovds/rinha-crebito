package chi

import (
	c "github.com/go-chi/chi/v5"
	"github.com/joaovds/rinha-crebito/internal/infra/chi/routes"
)

type Chi struct {
	Mux *c.Mux
}

func NewChi() *Chi {
	return &Chi{
		Mux: c.NewMux(),
	}
}

func SetupChi() *Chi {
	chiInstance := NewChi()

	routes.SetupRoutes(chiInstance.Mux)

	return chiInstance
}
