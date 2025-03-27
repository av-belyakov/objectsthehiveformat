package caseobject

import "github.com/av-belyakov/objectsthehiveformat/common"

// EventCaseObject объект события
type EventCaseObject struct {
	Tags         map[string][]string `json:"tags" bson:"tags"`                           //список тегов, где ключ это тип тега
	TagsAll      []string            `json:"tagsAll" bson:"tagsAll"`                     //список всех тегов
	CustomFields common.CustomFields `json:"customFields,omitempty" bson:"customFields"` //настраиваемые поля
	common.CommonEventCaseObject
}
