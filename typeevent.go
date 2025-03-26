package objectsthehiveformat

import "github.com/av-belyakov/objectsthehiveformat/common"

// TypeEvent информация о событии
type TypeEvent struct {
	common.CommonEventType
	Details EventCaseDetails `json:"details,omitzero" bson:"details"` //детальная информация о событии
	Object  EventCaseObject  `json:"object,omitzero" bson:"object"`   //объект события
}
