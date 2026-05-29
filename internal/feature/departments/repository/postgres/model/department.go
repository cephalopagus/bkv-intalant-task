package model

import "time"

type Department struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:200;not null"        json:"name"`
	ParentID  *uint     `gorm:"index"                    json:"parent_id"`
	CreatedAt time.Time `                                json:"created_at"`

	Children  []*Department `gorm:"foreignKey:ParentID"      json:"children,omitempty"`
	Employees []Employee    `gorm:"foreignKey:DepartmentID"  json:"employees,omitempty"`
}

type Employee struct {
	ID           uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	DepartmentID uint       `gorm:"not null;index"           json:"department_id"`
	FullName     string     `gorm:"size:200;not null"        json:"full_name"`
	Position     string     `gorm:"size:200;not null"        json:"position"`
	HiredAt      *time.Time `                                json:"hired_at"`
	CreatedAt    time.Time  `                                json:"created_at"`
}
