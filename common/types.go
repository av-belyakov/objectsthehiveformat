package common

// CommonEventType общие поля для описания события
type CommonEventType struct {
	Base           bool   `json:"base,omitempty" bson:"base"`                     //основа
	StartDate      string `json:"startDate,omitempty" bson:"startDate"`           //начальная дата в формате RFC3339
	RootId         string `json:"rootId,omitempty" bson:"rootId"`                 //основной уникальный идентификатор
	Organisation   string `json:"organisation,omitempty" bson:"organisation"`     //наименование организации
	OrganisationId string `json:"organisationId,omitempty" bson:"organisationId"` //уникальный идентификатор организации
	ObjectId       string `json:"objectId,omitempty" bson:"objectId"`             //уникальный идентификатор объекта
	ObjectType     string `json:"objectType,omitempty" bson:"objectType"`         //тип объекта
	Operation      string `json:"operation,omitempty" bson:"operation"`           //операция
	RequestId      string `json:"requestId,omitempty" bson:"requestId"`           //уникальный идентификатор запроса
}
