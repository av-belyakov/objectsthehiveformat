package caseobservables

import "github.com/av-belyakov/objectsthehiveformat/common"

// Observables список наблюдаемых сообщений
type Observables struct {
	Observables map[string][]Observable `json:"observables" bson:"observables"` //наблюдаемые сообщения
}

// Observable наблюдаемое сообщение
type Observable struct {
	Tags           map[string][]string `json:"tags" bson:"tags"`                               //список тегов
	SnortSid       []string            `json:"snortSid,omitempty" bson:"snortSid"`             //список идентификаторов сигнатур (строка)
	TagsAll        []string            `json:"tagsAll" bson:"tagsAll"`                         //список всех тегов
	SnortSidNumber []int               `json:"SnortSidNumber,omitempty" bson:"SnortSidNumber"` //список идентификаторов сигнатур (число)
	SensorId       string              `json:"sensorId,omitempty" bson:"sensorId"`             //идентификатор сенсора
	common.CommonObservableType
	Attachment common.AttachmentData `json:"attachment,omitzero" bson:"attachment"` //приложенные данные
}
