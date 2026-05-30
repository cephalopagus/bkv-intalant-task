package departments_service

import (
	core_logger "github.com/cephalopagus/bkv-intalant-task/internal/core/logger"
	departments_repository_postgres "github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/repository/postgres"
)

type DepartmentService struct {
	repo   *departments_repository_postgres.DepartmentRepository
	logger *core_logger.Logger
}

func NewDepartmentService(
	repo *departments_repository_postgres.DepartmentRepository,
	logger *core_logger.Logger) *DepartmentService {
	return &DepartmentService{repo: repo, logger: logger}
}
