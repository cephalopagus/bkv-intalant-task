package main

import (
	"log/slog"
	"net/http"
	"os"

	core_repository_postgres "github.com/cephalopagus/bkv-intalant-task/internal/core/repository/postgres"
	departments_repository_postgres "github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/repository/postgres"
	departments_service "github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/service"
	departments_transport_http "github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/transport/http"
	employee_repository_postgres "github.com/cephalopagus/bkv-intalant-task/internal/feature/employee/repository/postgres"
	employee_service "github.com/cephalopagus/bkv-intalant-task/internal/feature/employee/service"
	employee_transport_http "github.com/cephalopagus/bkv-intalant-task/internal/feature/employee/transport/http"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})))

	cfg := core_repository_postgres.Load()

	db, err := core_repository_postgres.NewPostgresDB(cfg)
	if err != nil {
		slog.Error("failed to connect to db", "err", err)
		os.Exit(1)
	}
	slog.Info("connected to db", "host", cfg.DBHost, "db", cfg.DBName)

	deptRepo := departments_repository_postgres.NewDepartmentRepository(db)
	empRepo := employee_repository_postgres.NewEmployeeRepository(db)

	deptSvc := departments_service.NewDepartmentService(deptRepo)
	empSvc := employee_service.NewEmployeeService(empRepo, deptRepo)

	mux := http.NewServeMux()

	deptH := departments_transport_http.NewDepartmentHandler(deptSvc)
	empH := employee_transport_http.NewEmployeeHandler(empSvc)

	deptH.Register(mux)
	empH.Register(mux)

	addr := ":8080"
	slog.Info("starting server", "addr", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		slog.Error("ListenAndServe", "err", err)
		os.Exit(1)
	}

}
