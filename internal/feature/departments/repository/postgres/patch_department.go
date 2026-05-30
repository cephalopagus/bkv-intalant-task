package departments_repository_postgres

import "github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/repository/postgres/model"

func (r *DepartmentRepository) Update(dept *model.Department) error {
	return r.db.Save(dept).Error
}
