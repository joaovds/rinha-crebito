package chi


import (
  c "github.com/go-chi/chi/v5"
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
  return NewChi()
}
