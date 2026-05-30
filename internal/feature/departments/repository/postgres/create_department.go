package departments_repository_postgres

import "github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/repository/postgres/model"

func (r *DepartmentRepository) Create(dept *model.Department) error {
	return r.db.Create(dept).Error
}
