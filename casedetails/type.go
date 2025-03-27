package casedetails

import "github.com/av-belyakov/objectsthehiveformat/common"

// EventDetails детальная информация о событии
type EventCaseDetails struct {
	CustomFields     common.CustomFields `json:"customFields,omitempty" bson:"customFields"`         //настраиваемые поля
	EndDate          string              `json:"endDate,omitempty" bson:"endDate"`                   //конечное дата и время в формате RFC3339
	ResolutionStatus string              `json:"resolutionStatus,omitempty" bson:"resolutionStatus"` //статус постановления
	Summary          string              `json:"summary,omitempty" bson:"summary"`                   //резюме
	Status           string              `json:"status,omitempty" bson:"status"`                     //статус
	ImpactStatus     string              `json:"impactStatus,omitempty" bson:"impactStatus"`         //краткое описание воздействия
}
