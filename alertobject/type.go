package alertobject

import "github.com/av-belyakov/objectsthehiveformat/common"

// EventAlertObject объект события
type EventAlertObject struct {
	common.CommonEventAlertObject
	Tags         map[string][]string `json:"tags" bson:"tags"`                           //теги после обработки
	TagsAll      []string            `json:"tagsAll" bson:"tagsAll"`                     //все теги
	CustomFields common.CustomFields `json:"customFields,omitempty" bson:"customFields"` //настраиваемые поля
}
