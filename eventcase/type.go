package objectsthehiveformat

import (
	"github.com/av-belyakov/objectsthehiveformat/casedetails"
	"github.com/av-belyakov/objectsthehiveformat/caseobject"
	"github.com/av-belyakov/objectsthehiveformat/common"
)

// TypeEvent информация о событии
type TypeEventForCase struct {
	common.CommonEventType
	Details casedetails.EventCaseDetails `json:"details,omitzero" bson:"details"` //детальная информация о событии
	Object  caseobject.EventCaseObject   `json:"object,omitzero" bson:"object"`   //объект события
}
