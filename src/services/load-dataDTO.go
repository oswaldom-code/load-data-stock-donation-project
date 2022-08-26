package services

import (
	"encoding/json"
	"strconv"

	"github.com/oswaldom-code/load-data-stock-donation-project/src/domain/models"
)

type JsonObject struct {
	Field      string `json:"field"`
	Width      int32  `json:"width"`
	Height     int32  `json:"height"`
	X          int32  `json:"x"`
	Y          int32  `json:"y"`
	PageNumber string `json:"page_number"`
	BrokerID   string `json:"broker_id"`
}

type JsonObjects []JsonObject

func UnmarshalField(data []byte) ([]models.DocumentField, error) {
	var jsonObjects JsonObjects
	err := json.Unmarshal(data, &jsonObjects)
	if err != nil {
		return []models.DocumentField{}, err
	}
	return JsonObjects2Fields(jsonObjects)
}

func (r *JsonObject) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func JsonObject2Field(jsonObject JsonObject) (models.DocumentField, error) {
	// string to int32
	pageNumber, err := strconv.ParseInt(jsonObject.PageNumber, 10, 32)
	if err != nil {
		return models.DocumentField{}, err
	}
	return models.DocumentField{
		Name:               jsonObject.Field,
		Width:              jsonObject.Width,
		Height:             jsonObject.Height,
		X:                  jsonObject.X,
		Y:                  jsonObject.Y,
		PageNumber:         int32(pageNumber),
		DocumentTemplateId: 1, // TODO: set this value to the template id
	}, nil
}

func JsonObjects2Fields(jsonObjects JsonObjects) ([]models.DocumentField, error) {
	var fields []models.DocumentField
	for _, jsonObject := range jsonObjects {
		field, err := JsonObject2Field(jsonObject)
		if err != nil {
			return fields, err
		}
		fields = append(fields, field)
	}
	return fields, nil
}
