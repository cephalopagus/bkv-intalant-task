package departments_transport_http

import (
	"net/http"

	core_response "github.com/cephalopagus/bkv-intalant-task/internal/core/response"
)

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, ok := core_response.PathUint(r, "id")
	if !ok {
		core_response.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}

	depth := core_response.QueryInt(r, "depth", 1)
	includeEmployees := core_response.QueryBool(r, "include_employees", true)

	dept, err := h.svc.GetWithTree(id, depth, includeEmployees)
	if err != nil {
		core_response.WriteServiceError(w, err)
		return
	}
	core_response.WriteJSON(w, http.StatusOK, dept)
}
