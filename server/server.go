package server

import (
	"fmt"
	"net/http"
	"sync"
)

// Handler ...
type Handler struct {
	sync.Mutex
	count int
}

// ServerHTTP ...
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()
	h.count++

	fmt.Fprintf(w, "hits %d", h.count)
}
