package departments_repository_postgres

import (
	"errors"

	"github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/repository/postgres/model"
	"gorm.io/gorm"
)

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

func (r *DepartmentRepository) IsDescendant(ancestorID, candidateID uint) (bool, error) {
	visited := map[uint]bool{}
	return r.isDescendantRecursive(ancestorID, candidateID, visited)
}

func (r *DepartmentRepository) isDescendantRecursive(ancestorID, candidateID uint, visited map[uint]bool) (bool, error) {
	if visited[ancestorID] {
		return false, nil
	}
	visited[ancestorID] = true

	var children []model.Department
	if err := r.db.Select("id").Where("parent_id = ?", ancestorID).Find(&children).Error; err != nil {
		return false, err
	}
	for _, c := range children {
		if c.ID == candidateID {
			return true, nil
		}
		if found, err := r.isDescendantRecursive(c.ID, candidateID, visited); err != nil || found {
			return found, err
		}
	}
	return false, nil
}
