package departments_service

import (
	departments_repository_postgres "github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/repository/postgres"
)

type DepartmentService struct {
	repo *departments_repository_postgres.DepartmentRepository
}

func NewDepartmentService(repo *departments_repository_postgres.DepartmentRepository) *DepartmentService {
	return &DepartmentService{repo: repo}
}
