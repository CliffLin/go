package web

import (
	"context"
	"net/http"
	"time"
)

type adminHandler struct {
	store Store
}

func adminGet(store Store, w http.ResponseWriter, r *http.Request) {
	p := parseName("/admin/", r.URL.Path)

	if p == "" {
		writeJSONOk(w)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	if p == "dumps" {
		if golinks, err := store.GetLinks(ctx, ""); err != nil {
			writeJSONBackendError(w, err)
			return
		} else {
			writeJSON(w, golinks, http.StatusOK)
		}
	}

}

func (h *adminHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		adminGet(h.store, w, r)
	default:
		writeJSONError(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusOK) // fix
	}
}
