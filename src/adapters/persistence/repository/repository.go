package repository

import (
	"github.com/oswaldom-code/load-data-stock-donation-project/src/domain/models"
	"gorm.io/gorm"
)

var sessionConfig = &gorm.Session{SkipDefaultTransaction: true, FullSaveAssociations: false}

func (s *store) TestDb() error {
	db, err := s.db.Session(sessionConfig).DB()
	if err != nil {
		return err
	}
	return db.Ping()
}

func (s *store) InsertData(fields []models.DocumentField) (int64, error) {
	db := s.db.Session(sessionConfig).Omit("broker_id").Omit("deleted_at").Save(&fields)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
