package main

import (
	"embed"
	"fmt"
	"net/http"
	"os"

	core_logger "github.com/cephalopagus/bkv-intalant-task/internal/core/logger"
	core_middleware "github.com/cephalopagus/bkv-intalant-task/internal/core/middleware"
	core_repository_postgres "github.com/cephalopagus/bkv-intalant-task/internal/core/repository/postgres"
	departments_repository_postgres "github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/repository/postgres"
	departments_service "github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/service"
	departments_transport_http "github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/transport/http"
	employee_repository_postgres "github.com/cephalopagus/bkv-intalant-task/internal/feature/employee/repository/postgres"
	employee_service "github.com/cephalopagus/bkv-intalant-task/internal/feature/employee/service"
	employee_transport_http "github.com/cephalopagus/bkv-intalant-task/internal/feature/employee/transport/http"
	migrate "github.com/cephalopagus/bkv-intalant-task/migrations"
	"go.uber.org/zap"
)

var migrationsFS embed.FS

func main() {

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init app logger")
		os.Exit(1)
	}
	defer logger.Close()

	cfg := core_repository_postgres.Load()

	db, err := core_repository_postgres.NewPostgresDB(cfg)
	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
		os.Exit(1)
	}
	logger.Info("connected to db",
		zap.String("host", cfg.DBHost),
		zap.String("db", cfg.DBName),
	)

	sqlDB, _ := db.DB()
	if err := migrate.Up(sqlDB); err != nil {
		logger.Fatal("migrations failed", zap.Error(err))
	}
	logger.Info("migrations applied")

	deptRepo := departments_repository_postgres.NewDepartmentRepository(db)
	empRepo := employee_repository_postgres.NewEmployeeRepository(db)

	deptSvc := departments_service.NewDepartmentService(deptRepo, logger)
	empSvc := employee_service.NewEmployeeService(empRepo, deptRepo, logger)

	mux := http.NewServeMux()

	deptH := departments_transport_http.NewDepartmentHandler(deptSvc, logger)
	empH := employee_transport_http.NewEmployeeHandler(empSvc, logger)

	deptH.Register(mux)
	empH.Register(mux)

	addr := ":8080"

	logger.Info("starting server", zap.String("addr", addr))
	if err := http.ListenAndServe(addr, core_middleware.Logging(mux, logger)); err != nil {
		logger.Fatal("ListenAndServe", zap.Error(err))
		os.Exit(1)
	}

}
