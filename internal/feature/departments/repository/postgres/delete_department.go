package departments_repository_postgres

import (
	"github.com/cephalopagus/bkv-intalant-task/internal/feature/departments/repository/postgres/model"
	"gorm.io/gorm"
)

func (r *DepartmentRepository) Delete(id uint) error {
	return r.db.Delete(&model.Department{}, id).Error
}

func (r *DepartmentRepository) DeleteAndReassign(id, targetID uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Employee{}).
			Where("department_id = ?", id).
			Update("department_id", targetID).Error; err != nil {
			return err
		}

		if err := reassignDescendants(tx, id, targetID); err != nil {
			return err
		}

		return tx.Delete(&model.Department{}, id).Error
	})
}

func reassignDescendants(tx *gorm.DB, parentID, targetID uint) error {
	var children []model.Department
	if err := tx.Where("parent_id = ?", parentID).Find(&children).Error; err != nil {
		return err
	}
	for _, child := range children {
		if err := tx.Model(&model.Employee{}).
			Where("department_id = ?", child.ID).
			Update("department_id", targetID).Error; err != nil {
			return err
		}
		if err := reassignDescendants(tx, child.ID, targetID); err != nil {
			return err
		}
	}
	return nil
}
