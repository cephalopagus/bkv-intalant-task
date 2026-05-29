package departments_transport_http

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	departments_service "github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/service"
)

type Handler struct {
	svc *departments_service.DepartmentService
}

func New(svc *departments_service.DepartmentService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /departments/", h.Create)
	mux.HandleFunc("GET /departments/{id}", h.GetByID)
	mux.HandleFunc("PATCH /departments/{id}", h.Update)
	mux.HandleFunc("DELETE /departments/{id}", h.Delete)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name     string `json:"name"`
		ParentID *uint  `json:"parent_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	dept, err := h.svc.Create(body.Name, body.ParentID)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, dept)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, ok := pathUint(r, "id")
	if !ok {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	depth := queryInt(r, "depth", 1)
	includeEmployees := queryBool(r, "include_employees", true)

	dept, err := h.svc.GetWithTree(id, depth, includeEmployees)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, dept)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, ok := pathUint(r, "id")
	if !ok {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var body struct {
		Name     *string `json:"name"`
		ParentID *uint   `json:"parent_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	dept, err := h.svc.Update(id, body.Name, body.ParentID)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, dept)
}

// DELETE /departments/{id}
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, ok := pathUint(r, "id")
	if !ok {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.svc.Delete(id); err != nil {
		writeServiceError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		slog.Error("writeJSON", "err", err)
	}
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func writeServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, departments_service.ErrNotFound):
		writeError(w, http.StatusNotFound, err.Error())
	case errors.Is(err, departments_service.ErrConflict):
		writeError(w, http.StatusConflict, err.Error())
	case errors.Is(err, departments_service.ErrBadRequest):
		writeError(w, http.StatusBadRequest, err.Error())
	default:
		slog.Error("internal error", "err", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
	}
}

func pathUint(r *http.Request, key string) (uint, bool) {
	s := r.PathValue(key)
	v, err := strconv.ParseUint(s, 10, 64)
	return uint(v), err == nil
}

func queryInt(r *http.Request, key string, def int) int {
	s := r.URL.Query().Get(key)
	if s == "" {
		return def
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return v
}

func queryBool(r *http.Request, key string, def bool) bool {
	s := strings.ToLower(r.URL.Query().Get(key))
	switch s {
	case "true", "1":
		return true
	case "false", "0":
		return false
	default:
		return def
	}
}
