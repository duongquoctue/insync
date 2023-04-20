package handlers

import (
	"fmt"
	"net/http"

	"github.com/duongquoctue/insync/core/corehttp"
	insync_model "github.com/duongquoctue/insync/pkg/proto/user/model"
	insync_service "github.com/duongquoctue/insync/pkg/proto/user/service"
)

type GetUserHandler struct {
	UserClient insync_service.UserClient
}

func (h *GetUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	model, err := h.UserClient.Create(r.Context(), &insync_service.CreateRequest{
		User: &insync_model.User{
			Username:  "duongquoctue",
			FirstName: "Tue",
			LastName:  "Duong",
		},
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	corehttp.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"nameeee": model.FirstName,
		"another": model.Username,
	})
}
