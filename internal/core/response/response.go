package core_response

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	core_errors "github.com/cephalopagus/bkv-intalant-task/internal/core/errors"
)

func PathUint(r *http.Request, key string) (uint, bool) {
	s := r.PathValue(key)
	v, err := strconv.ParseUint(s, 10, 64)
	return uint(v), err == nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		slog.Error("writeJSON", "errror", err)
	}
}

func WriteError(w http.ResponseWriter, status int, msg string) {
	WriteJSON(w, status, map[string]string{"error": msg})
}

func WriteServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, core_errors.ErrNotFound):
		WriteError(w, http.StatusNotFound, err.Error())
	case errors.Is(err, core_errors.ErrConflict):
		WriteError(w, http.StatusConflict, err.Error())
	case errors.Is(err, core_errors.ErrBadRequest):
		WriteError(w, http.StatusBadRequest, err.Error())
	default:
		slog.Error("internal error", "err", err)
		WriteError(w, http.StatusInternalServerError, "internal server error")
	}
}
func QueryInt(r *http.Request, key string, def int) int {
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

func QueryBool(r *http.Request, key string, def bool) bool {
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
