package handlers

import (
	"fmt"
	"net/http"

	"github.com/duongquoctue/insync/core/corehttp"
)

type SendMessageHandler struct {
	SomeToken string
}

func (h *SendMessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	corehttp.RespondJSON(w, http.StatusOK, fmt.Sprintf("this is good ehh %v", h.SomeToken))
}
