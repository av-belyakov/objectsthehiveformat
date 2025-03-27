package objectsthehiveformat

import (
	"github.com/av-belyakov/objectsthehiveformat/alertdetails"
	"github.com/av-belyakov/objectsthehiveformat/alertobject"
	"github.com/av-belyakov/objectsthehiveformat/common"
)

// TypeEventForAlert сообщение с информацией о событии
type TypeEventForAlert struct {
	common.CommonEventType
	Details alertdetails.EventAlertDetails `json:"details,omitzero" bson:"details"`
	Object  alertobject.EventAlertObject   `json:"object,omitzero" bson:"object"`
}
