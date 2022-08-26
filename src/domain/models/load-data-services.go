package models

import (
	"time"
)

type DocumentField struct {
	ID                 int64      `gorm:"column:id"`
	Name               string     `gorm:"column:name"`
	X                  int32      `gorm:"column:x"`
	Y                  int32      `gorm:"column:y"`
	Width              int32      `gorm:"column:width"`
	Height             int32      `gorm:"column:height"`
	DataType           string     `gorm:"column:data_type"`
	PageNumber         int32      `gorm:"column:page_number"`
	DocumentTemplateId int64      `gorm:"column:document_templates_id"`
	CreatedAt          *time.Time `gorm:"column:created_at"`
	UpdatedAt          *time.Time `gorm:"column:updated_at"`
	DeletedAt          *time.Time `gorm:"column:deleted_at"`
}
