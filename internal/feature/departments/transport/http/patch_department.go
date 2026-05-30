package departments_transport_http

import (
	"encoding/json"
	"net/http"

	core_response "github.com/cephalopagus/bkv-intalant-task/internal/core/response"
)

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, ok := core_response.PathUint(r, "id")
	if !ok {
		core_response.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var body struct {
		Name     *string `json:"name"`
		ParentID *uint   `json:"parent_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		core_response.WriteError(w, http.StatusBadRequest, "invalid json")
		return
	}

	dept, err := h.svc.Update(id, body.Name, body.ParentID)
	if err != nil {
		core_response.WriteServiceError(w, err)
		return
	}
	core_response.WriteJSON(w, http.StatusOK, dept)
}
