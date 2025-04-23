package alert

import (
	"github.com/av-belyakov/objectsthehiveformat/alertartifacts"
	"github.com/av-belyakov/objectsthehiveformat/common"
)

// TypeAlert сообщение с информацией о Alert-те
type TypeAlert struct {
	Artifacts    map[string][]alertartifacts.Artifacts `json:"artifact" bson:"artifact"`                   //артефакты
	CustomFields common.CustomFields                   `json:"customFields,omitempty" bson:"customFields"` //настраиваемые поля
	Tags         map[string][]string                   `json:"tags" bson:"tags"`                           //теги после обработки
	TagsAll      []string                              `json:"tagsAll" bson:"tagsAll"`                     //все теги
	common.CommonEventAlertObject
	Severity uint64 `json:"severity,omitempty" bson:"severity"` //строгость
	Follow   bool   `json:"follow,omitempty" bson:"follow"`     //следовать
}
