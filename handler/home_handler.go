package handler

import (
	"net/http"

	"github.com/rysmaadit/go-template/common/responder"
)

func Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := "Welcome"
		responder.NewHttpResponse(r, w, http.StatusOK, res, nil)
	}
}
