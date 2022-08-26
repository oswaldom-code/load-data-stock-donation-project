package ports

import "github.com/oswaldom-code/load-data-stock-donation-project/src/domain/models"

type Repository interface {
	TestDb() error
	InsertData(field []models.DocumentField) (int64, error)
}
