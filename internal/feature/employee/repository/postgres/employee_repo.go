package employee_repository_postgres

import (
	"github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/repository/postgres/model"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) Create(emp *model.Employee) error {
	return r.db.Create(emp).Error
}
