package handlers

import (
	"fmt"
	"net/http"

	"github.com/duongquoctue/insync/core/corehttp"
	websocket "github.com/duongquoctue/insync/core/corews"
	"github.com/duongquoctue/insync/core/logger"
)

type SubcribeHandler struct {
	SomeToken string
}

func (h *SubcribeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := logger.NewLoggerCtx(r.Context())

	wsConn, err := websocket.NewConnection(w, r)
	if err != nil {
		corehttp.RespondErr(w, http.StatusForbidden, err.Error())
		return
	}

	defer func() {
		defer wsConn.Close(websocket.InternalError, "")

	}()

	wsConn.Close(websocket.NormalClosure, "")
	corehttp.RespondJSON(w, http.StatusOK, fmt.Sprintf("this is good ehh %v", h.SomeToken))
}
