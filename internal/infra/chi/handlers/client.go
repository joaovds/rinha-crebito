package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	c "github.com/go-chi/chi/v5"
)

type clientHandler struct {}

func NewClientHandler() *clientHandler {
  return &clientHandler{}
}

func (h *clientHandler) GetExtract(w http.ResponseWriter, r *http.Request) {
  clientIdParam := c.URLParam(r, "id")
  clientId, err:= strconv.Atoi(clientIdParam)
  if err != nil {
    http.Error(w, "Client not found", http.StatusNotFound)
    return
  }

  w.WriteHeader(http.StatusOK)
  w.Write([]byte(fmt.Sprintf("Extract for client %d", clientId)))
}
