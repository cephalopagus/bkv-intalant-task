package departments_repository_postgres

import (
	"errors"

	"github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/repository/postgres/model"
	"gorm.io/gorm"
)

type DepartmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) *DepartmentRepository {
	return &DepartmentRepository{db: db}
}

func (r *DepartmentRepository) Create(dept *model.Department) error {
	return r.db.Create(dept).Error
}

func (r *DepartmentRepository) GetByID(id uint) (*model.Department, error) {
	var dept model.Department
	err := r.db.First(&dept, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &dept, err
}
func (r *DepartmentRepository) GetWithTree(id uint, depth int, includeEmployees bool) (*model.Department, error) {
	var dept model.Department
	if err := r.db.First(&dept, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	if includeEmployees {
		r.db.Where("department_id = ?", id).
			Order("created_at").
			Find(&dept.Employees)
	}

	if depth > 0 {
		r.loadChildren(&dept, depth-1, includeEmployees)
	}

	return &dept, nil
}

func (r *DepartmentRepository) loadChildren(dept *model.Department, remaining int, includeEmployees bool) error {
	var children []*model.Department
	if err := r.db.Where("parent_id = ?", dept.ID).Find(&children).Error; err != nil {
		return err
	}

	for _, child := range children {
		if includeEmployees {
			r.db.Where("department_id = ?", child.ID).
				Order("created_at").
				Find(&child.Employees)
		}
		if remaining > 0 {
			r.loadChildren(child, remaining-1, includeEmployees)
		}
	}

	dept.Children = children
	return nil
}

func (r *DepartmentRepository) Update(dept *model.Department) error {
	return r.db.Save(dept).Error
}

func (r *DepartmentRepository) Delete(id uint) error {
	return r.db.Delete(&model.Department{}, id).Error
}
