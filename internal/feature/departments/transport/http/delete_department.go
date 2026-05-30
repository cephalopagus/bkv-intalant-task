package departments_transport_http

import (
	"net/http"
	"strconv"

	core_response "github.com/cephalopagus/bkv-intalant-task/internal/core/response"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, ok := core_response.PathUint(r, "id")
	if !ok {
		core_response.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}

	mode := r.URL.Query().Get("mode")
	if mode == "" {
		core_response.WriteError(w, http.StatusBadRequest, "mode is required (cascade or reassign)")
		return
	}

	var reassignTo *uint
	if s := r.URL.Query().Get("reassign_to_department_id"); s != "" {
		v, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			core_response.WriteError(w, http.StatusBadRequest, "invalid reassign_to_department_id")
			return
		}
		u := uint(v)
		reassignTo = &u
	}

	if err := h.svc.Delete(id, mode, reassignTo); err != nil {
		core_response.WriteServiceError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
